package services

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
)

const (
	resourceType           = "sampleResource"
	mockKafkaRequestID     = "1ilzo99dVkVAoQNJeovhP8pFIzS" // sample kafka request ID generated by ksuid
	mockIDWithInvalidChars = "vp&xG^nl9MStC@SI*#c$6V^TKq0"
)

func Test_handleGetError(t *testing.T) {
	type args struct {
		resourceType string
		field        string
		value        interface{}
		err          error
	}
	tests := []struct {
		name string
		args args
		want *errors.ServiceError
	}{
		{
			name: "Handler should return a general error for any errors other than record not found",
			args: args{
				resourceType: resourceType,
				field:        "id",
				value:        "sample-id",
				err:          gorm.ErrInvalidSQL,
			},
			want: errors.GeneralError("Unable to find %s with id='sample-id': %s", resourceType, gorm.ErrInvalidSQL.Error()),
		},
		{
			name: "Handler should return a not found error if record was not found in the database",
			args: args{
				resourceType: resourceType,
				field:        "id",
				value:        "sample-id",
				err:          gorm.ErrRecordNotFound,
			},
			want: errors.NotFound("%s with id='sample-id' not found", resourceType),
		},
		{
			name: "Handler should redact sensitive fields from the error message",
			args: args{
				resourceType: resourceType,
				field:        "email",
				value:        "sample@example.com",
				err:          gorm.ErrRecordNotFound,
			},
			want: errors.NotFound("%s with email='<redacted>' not found", resourceType),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleGetError(tt.args.resourceType, tt.args.field, tt.args.value, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleGetError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleCreateError(t *testing.T) {
	type args struct {
		resourceType string
		err          error
	}
	tests := []struct {
		name string
		args args
		want *errors.ServiceError
	}{
		{
			name: "Handler should return a general error for any other errors than violating unique constraints",
			args: args{
				resourceType: resourceType,
				err:          gorm.ErrInvalidSQL,
			},
			want: errors.GeneralError("Unable to create %s: %s", resourceType, gorm.ErrInvalidSQL.Error()),
		},
		{
			name: "Handler should return a conflict error if creation error is due to violating unique constraints",
			args: args{
				resourceType: resourceType,
				err:          fmt.Errorf("transaction violates unique constraints"),
			},
			want: errors.Conflict("This %s already exists", resourceType),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleCreateError(tt.args.resourceType, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleCreateError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handleUpdateError(t *testing.T) {
	type args struct {
		resourceType string
		err          error
	}
	tests := []struct {
		name string
		args args
		want *errors.ServiceError
	}{
		{
			name: "Handler should return a general error for any other errors than violating unique constraints",
			args: args{
				resourceType: resourceType,
				err:          gorm.ErrInvalidSQL,
			},
			want: errors.GeneralError("Unable to update %s: %s", resourceType, gorm.ErrInvalidSQL.Error()),
		},
		{
			name: "Handler should return a conflict error if update error is due to violating unique constraints",
			args: args{
				resourceType: resourceType,
				err:          fmt.Errorf("transaction violates unique constraints"),
			},
			want: errors.Conflict("Changes to %s conflict with existing records", resourceType),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleUpdateError(tt.args.resourceType, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleUpdateError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_truncateString(t *testing.T) {
	exampleString := "example-string"
	type args struct {
		str string
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should truncate string successfully",
			args: args{
				str: exampleString,
				num: 10,
			},
			want: exampleString[0:10],
		},
		{
			name: "should not truncate string if wanted length is less than given string length",
			args: args{
				str: exampleString,
				num: 15,
			},
			want: exampleString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := truncateString(tt.args.str, tt.args.num); got != tt.want {
				t.Errorf("truncateString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildKafkaNamespaceIdentifier(t *testing.T) {
	mockShortOwnerUsername := "sample_owner_username_short"
	mockLongOwnerUsername := fmt.Sprintf("sample_owner_username_long_%s", mockKafkaRequestID)
	namespaceLimit := 63 // Maximum namespace name length as validated by OpenShift

	type args struct {
		kafkaRequest *api.KafkaRequest
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "build kafka namespace id successfully with a short owner username",
			args: args{
				kafkaRequest: &api.KafkaRequest{
					Owner: mockShortOwnerUsername,
				},
			},
			want: fmt.Sprintf("%s-%s", mockShortOwnerUsername, strings.ToLower(mockKafkaRequestID)),
		},
		{
			name: "build kafka namespace id successfully with a long owner username",
			args: args{
				kafkaRequest: &api.KafkaRequest{
					Owner: mockLongOwnerUsername,
				},
			},
			want: fmt.Sprintf("%s-%s", mockLongOwnerUsername[0:truncatedNamespaceLen], strings.ToLower(mockKafkaRequestID)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.kafkaRequest.ID = mockKafkaRequestID
			got := buildKafkaNamespaceIdentifier(tt.args.kafkaRequest)
			if got != tt.want {
				t.Errorf("buildKafkaNamespaceIdentifier() = %v, want %v", got, tt.want)
			}
			if len(got) > namespaceLimit {
				t.Errorf("buildKafkaNamespaceIdentifier() namespace identifier length is %v, this is over the %v maximum namespace name limit", len(got), namespaceLimit)
			}
		})
	}
}

func Test_buildTruncateKafkaIdentifier(t *testing.T) {
	mockShortKafkaName := "kafka"
	mockLongKafkaName := "sample-kafka-name-long"

	type args struct {
		kafkaRequest *api.KafkaRequest
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "build kafka identifier with a short name successfully",
			args: args{
				kafkaRequest: &api.KafkaRequest{
					Name: mockShortKafkaName,
				},
			},
			want: fmt.Sprintf("%s-%s", mockShortKafkaName, strings.ToLower(mockKafkaRequestID)),
		},
		{
			name: "build kafka identifier with a long name successfully",
			args: args{
				kafkaRequest: &api.KafkaRequest{
					Name: mockLongKafkaName,
				},
			},
			want: fmt.Sprintf("%s-%s", mockLongKafkaName[0:truncatedNameLen], strings.ToLower(mockKafkaRequestID)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.kafkaRequest.ID = mockKafkaRequestID
			if got := buildTruncateKafkaIdentifier(tt.args.kafkaRequest); got != tt.want {
				t.Errorf("buildTruncateKafkaIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildSyncsetIdentifier(t *testing.T) {
	mockKafkaName := "example-kafka"

	type args struct {
		kafkaRequest *api.KafkaRequest
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "build syncset identifier successfully",
			args: args{
				kafkaRequest: &api.KafkaRequest{
					Name: mockKafkaName,
				},
			},
			want: fmt.Sprintf("ext-%s-%s", mockKafkaName, strings.ToLower(mockKafkaRequestID)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.kafkaRequest.ID = mockKafkaRequestID
			if got := buildSyncsetIdentifier(tt.args.kafkaRequest); got != tt.want {
				t.Errorf("buildSyncsetIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskProceedingandTrailingDash(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should replace '-' prefix and suffix with a subdomain safe value",
			args: args{
				name: "-example-name-",
			},
			want: fmt.Sprintf("%[1]sexample-name%[1]s", appendChar),
		},
		{
			name: "should not replace '-' if its not a prefix or suffix of the given string",
			args: args{
				name: "example-name",
			},
			want: "example-name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskProceedingandTrailingDash(tt.args.name); got != tt.want {
				t.Errorf("maskProceedingandTrailingDash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceNamespaceSpecialChar(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "replace all invalid characters in an invalid namespace name",
			args: args{
				name: fmt.Sprintf("project-%s-", mockIDWithInvalidChars),
			},
			want: "project-vp-xg-nl9mstc-si-c-6v-tkq0a",
		},
		{
			name: "valid namespace should not be modified",
			args: args{
				name: fmt.Sprintf("project-%s", mockKafkaRequestID),
			},
			want: fmt.Sprintf("project-%s", strings.ToLower(mockKafkaRequestID)),
		},
		{
			name: "should return an error if given namespace name is an empty string",
			args: args{
				name: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := replaceNamespaceSpecialChar(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("replaceNamespaceSpecialChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("replaceNamespaceSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceHostSpecialChar(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "replace all invalid characters in an invalid host name",
			args: args{
				name: fmt.Sprintf("-host-%s", mockIDWithInvalidChars),
			},
			want: "ahost-vp-xg-nl-mstc-si-c--v-tkqa",
		},
		{
			name: "valid hostname should not be modified",
			args: args{
				name: "sample-host-name",
			},
			want: "sample-host-name",
		},
		{
			name: "should return an error if given host name is an empty string",
			args: args{
				name: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := replaceHostSpecialChar(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("replaceHostSpecialChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("replaceHostSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	searchedString := "findMe"
	someSlice := []string{"some", "string", "values"}
	sliceWithFindMe := []string{"some", "string", "values", "findMe"}
	var emptySlice []string
	type args struct {
		slice []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check for a string in an empty slice",
			args: args{
				s:     searchedString,
				slice: emptySlice,
			},
			want: false,
		},
		{
			name: "Check for a string in a non-empty slice that doesn't contain the string",
			args: args{
				s:     searchedString,
				slice: someSlice,
			},
			want: false,
		},
		{
			name: "Check for a string in a non-empty slice that contains that string",
			args: args{
				s:     searchedString,
				slice: sliceWithFindMe,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.args.slice, tt.args.s)
			if got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
