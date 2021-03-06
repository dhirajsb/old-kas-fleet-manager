// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"context"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/constants"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"sync"
)

// Ensure, that KafkaServiceMock does implement KafkaService.
// If this is not the case, regenerate this file with moq.
var _ KafkaService = &KafkaServiceMock{}

// KafkaServiceMock is a mock implementation of KafkaService.
//
//     func TestSomethingThatUsesKafkaService(t *testing.T) {
//
//         // make and configure a mocked KafkaService
//         mockedKafkaService := &KafkaServiceMock{
//             ChangeKafkaCNAMErecordsFunc: func(kafkaRequest *api.KafkaRequest, clusterDNS string, action string) (*route53.ChangeResourceRecordSetsOutput, *errors.ServiceError) {
// 	               panic("mock out the ChangeKafkaCNAMErecords method")
//             },
//             CreateFunc: func(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the Create method")
//             },
//             DeleteFunc: func(in1 *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the Delete method")
//             },
//             GetFunc: func(ctx context.Context, id string) (*api.KafkaRequest, *errors.ServiceError) {
// 	               panic("mock out the Get method")
//             },
//             GetByIdFunc: func(id string) (*api.KafkaRequest, *errors.ServiceError) {
// 	               panic("mock out the GetById method")
//             },
//             ListFunc: func(ctx context.Context, listArgs *ListArguments) (api.KafkaList, *api.PagingMeta, *errors.ServiceError) {
// 	               panic("mock out the List method")
//             },
//             ListByStatusFunc: func(status constants.KafkaStatus) ([]*api.KafkaRequest, *errors.ServiceError) {
// 	               panic("mock out the ListByStatus method")
//             },
//             RegisterKafkaDeprovisionJobFunc: func(ctx context.Context, id string) *errors.ServiceError {
// 	               panic("mock out the RegisterKafkaDeprovisionJob method")
//             },
//             RegisterKafkaJobFunc: func(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the RegisterKafkaJob method")
//             },
//             UpdateFunc: func(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the Update method")
//             },
//             UpdateStatusFunc: func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
// 	               panic("mock out the UpdateStatus method")
//             },
//         }
//
//         // use mockedKafkaService in code that requires KafkaService
//         // and then make assertions.
//
//     }
type KafkaServiceMock struct {
	// ChangeKafkaCNAMErecordsFunc mocks the ChangeKafkaCNAMErecords method.
	ChangeKafkaCNAMErecordsFunc func(kafkaRequest *api.KafkaRequest, clusterDNS string, action string) (*route53.ChangeResourceRecordSetsOutput, *errors.ServiceError)

	// CreateFunc mocks the Create method.
	CreateFunc func(kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(in1 *api.KafkaRequest) *errors.ServiceError

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id string) (*api.KafkaRequest, *errors.ServiceError)

	// GetByIdFunc mocks the GetById method.
	GetByIdFunc func(id string) (*api.KafkaRequest, *errors.ServiceError)

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, listArgs *ListArguments) (api.KafkaList, *api.PagingMeta, *errors.ServiceError)

	// ListByStatusFunc mocks the ListByStatus method.
	ListByStatusFunc func(status constants.KafkaStatus) ([]*api.KafkaRequest, *errors.ServiceError)

	// RegisterKafkaDeprovisionJobFunc mocks the RegisterKafkaDeprovisionJob method.
	RegisterKafkaDeprovisionJobFunc func(ctx context.Context, id string) *errors.ServiceError

	// RegisterKafkaJobFunc mocks the RegisterKafkaJob method.
	RegisterKafkaJobFunc func(kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// UpdateFunc mocks the Update method.
	UpdateFunc func(kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// UpdateStatusFunc mocks the UpdateStatus method.
	UpdateStatusFunc func(id string, status constants.KafkaStatus) (bool, *errors.ServiceError)

	// calls tracks calls to the methods.
	calls struct {
		// ChangeKafkaCNAMErecords holds details about calls to the ChangeKafkaCNAMErecords method.
		ChangeKafkaCNAMErecords []struct {
			// KafkaRequest is the kafkaRequest argument value.
			KafkaRequest *api.KafkaRequest
			// ClusterDNS is the clusterDNS argument value.
			ClusterDNS string
			// Action is the action argument value.
			Action string
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// KafkaRequest is the kafkaRequest argument value.
			KafkaRequest *api.KafkaRequest
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// In1 is the in1 argument value.
			In1 *api.KafkaRequest
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetById holds details about calls to the GetById method.
		GetById []struct {
			// ID is the id argument value.
			ID string
		}
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ListArgs is the listArgs argument value.
			ListArgs *ListArguments
		}
		// ListByStatus holds details about calls to the ListByStatus method.
		ListByStatus []struct {
			// Status is the status argument value.
			Status constants.KafkaStatus
		}
		// RegisterKafkaDeprovisionJob holds details about calls to the RegisterKafkaDeprovisionJob method.
		RegisterKafkaDeprovisionJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// RegisterKafkaJob holds details about calls to the RegisterKafkaJob method.
		RegisterKafkaJob []struct {
			// KafkaRequest is the kafkaRequest argument value.
			KafkaRequest *api.KafkaRequest
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// KafkaRequest is the kafkaRequest argument value.
			KafkaRequest *api.KafkaRequest
		}
		// UpdateStatus holds details about calls to the UpdateStatus method.
		UpdateStatus []struct {
			// ID is the id argument value.
			ID string
			// Status is the status argument value.
			Status constants.KafkaStatus
		}
	}
	lockChangeKafkaCNAMErecords     sync.RWMutex
	lockCreate                      sync.RWMutex
	lockDelete                      sync.RWMutex
	lockGet                         sync.RWMutex
	lockGetById                     sync.RWMutex
	lockList                        sync.RWMutex
	lockListByStatus                sync.RWMutex
	lockRegisterKafkaDeprovisionJob sync.RWMutex
	lockRegisterKafkaJob            sync.RWMutex
	lockUpdate                      sync.RWMutex
	lockUpdateStatus                sync.RWMutex
}

// ChangeKafkaCNAMErecords calls ChangeKafkaCNAMErecordsFunc.
func (mock *KafkaServiceMock) ChangeKafkaCNAMErecords(kafkaRequest *api.KafkaRequest, clusterDNS string, action string) (*route53.ChangeResourceRecordSetsOutput, *errors.ServiceError) {
	if mock.ChangeKafkaCNAMErecordsFunc == nil {
		panic("KafkaServiceMock.ChangeKafkaCNAMErecordsFunc: method is nil but KafkaService.ChangeKafkaCNAMErecords was just called")
	}
	callInfo := struct {
		KafkaRequest *api.KafkaRequest
		ClusterDNS   string
		Action       string
	}{
		KafkaRequest: kafkaRequest,
		ClusterDNS:   clusterDNS,
		Action:       action,
	}
	mock.lockChangeKafkaCNAMErecords.Lock()
	mock.calls.ChangeKafkaCNAMErecords = append(mock.calls.ChangeKafkaCNAMErecords, callInfo)
	mock.lockChangeKafkaCNAMErecords.Unlock()
	return mock.ChangeKafkaCNAMErecordsFunc(kafkaRequest, clusterDNS, action)
}

// ChangeKafkaCNAMErecordsCalls gets all the calls that were made to ChangeKafkaCNAMErecords.
// Check the length with:
//     len(mockedKafkaService.ChangeKafkaCNAMErecordsCalls())
func (mock *KafkaServiceMock) ChangeKafkaCNAMErecordsCalls() []struct {
	KafkaRequest *api.KafkaRequest
	ClusterDNS   string
	Action       string
} {
	var calls []struct {
		KafkaRequest *api.KafkaRequest
		ClusterDNS   string
		Action       string
	}
	mock.lockChangeKafkaCNAMErecords.RLock()
	calls = mock.calls.ChangeKafkaCNAMErecords
	mock.lockChangeKafkaCNAMErecords.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *KafkaServiceMock) Create(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
	if mock.CreateFunc == nil {
		panic("KafkaServiceMock.CreateFunc: method is nil but KafkaService.Create was just called")
	}
	callInfo := struct {
		KafkaRequest *api.KafkaRequest
	}{
		KafkaRequest: kafkaRequest,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(kafkaRequest)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedKafkaService.CreateCalls())
func (mock *KafkaServiceMock) CreateCalls() []struct {
	KafkaRequest *api.KafkaRequest
} {
	var calls []struct {
		KafkaRequest *api.KafkaRequest
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *KafkaServiceMock) Delete(in1 *api.KafkaRequest) *errors.ServiceError {
	if mock.DeleteFunc == nil {
		panic("KafkaServiceMock.DeleteFunc: method is nil but KafkaService.Delete was just called")
	}
	callInfo := struct {
		In1 *api.KafkaRequest
	}{
		In1: in1,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(in1)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedKafkaService.DeleteCalls())
func (mock *KafkaServiceMock) DeleteCalls() []struct {
	In1 *api.KafkaRequest
} {
	var calls []struct {
		In1 *api.KafkaRequest
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *KafkaServiceMock) Get(ctx context.Context, id string) (*api.KafkaRequest, *errors.ServiceError) {
	if mock.GetFunc == nil {
		panic("KafkaServiceMock.GetFunc: method is nil but KafkaService.Get was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedKafkaService.GetCalls())
func (mock *KafkaServiceMock) GetCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetById calls GetByIdFunc.
func (mock *KafkaServiceMock) GetById(id string) (*api.KafkaRequest, *errors.ServiceError) {
	if mock.GetByIdFunc == nil {
		panic("KafkaServiceMock.GetByIdFunc: method is nil but KafkaService.GetById was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetById.Lock()
	mock.calls.GetById = append(mock.calls.GetById, callInfo)
	mock.lockGetById.Unlock()
	return mock.GetByIdFunc(id)
}

// GetByIdCalls gets all the calls that were made to GetById.
// Check the length with:
//     len(mockedKafkaService.GetByIdCalls())
func (mock *KafkaServiceMock) GetByIdCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetById.RLock()
	calls = mock.calls.GetById
	mock.lockGetById.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *KafkaServiceMock) List(ctx context.Context, listArgs *ListArguments) (api.KafkaList, *api.PagingMeta, *errors.ServiceError) {
	if mock.ListFunc == nil {
		panic("KafkaServiceMock.ListFunc: method is nil but KafkaService.List was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ListArgs *ListArguments
	}{
		Ctx:      ctx,
		ListArgs: listArgs,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx, listArgs)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedKafkaService.ListCalls())
func (mock *KafkaServiceMock) ListCalls() []struct {
	Ctx      context.Context
	ListArgs *ListArguments
} {
	var calls []struct {
		Ctx      context.Context
		ListArgs *ListArguments
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// ListByStatus calls ListByStatusFunc.
func (mock *KafkaServiceMock) ListByStatus(status constants.KafkaStatus) ([]*api.KafkaRequest, *errors.ServiceError) {
	if mock.ListByStatusFunc == nil {
		panic("KafkaServiceMock.ListByStatusFunc: method is nil but KafkaService.ListByStatus was just called")
	}
	callInfo := struct {
		Status constants.KafkaStatus
	}{
		Status: status,
	}
	mock.lockListByStatus.Lock()
	mock.calls.ListByStatus = append(mock.calls.ListByStatus, callInfo)
	mock.lockListByStatus.Unlock()
	return mock.ListByStatusFunc(status)
}

// ListByStatusCalls gets all the calls that were made to ListByStatus.
// Check the length with:
//     len(mockedKafkaService.ListByStatusCalls())
func (mock *KafkaServiceMock) ListByStatusCalls() []struct {
	Status constants.KafkaStatus
} {
	var calls []struct {
		Status constants.KafkaStatus
	}
	mock.lockListByStatus.RLock()
	calls = mock.calls.ListByStatus
	mock.lockListByStatus.RUnlock()
	return calls
}

// RegisterKafkaDeprovisionJob calls RegisterKafkaDeprovisionJobFunc.
func (mock *KafkaServiceMock) RegisterKafkaDeprovisionJob(ctx context.Context, id string) *errors.ServiceError {
	if mock.RegisterKafkaDeprovisionJobFunc == nil {
		panic("KafkaServiceMock.RegisterKafkaDeprovisionJobFunc: method is nil but KafkaService.RegisterKafkaDeprovisionJob was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockRegisterKafkaDeprovisionJob.Lock()
	mock.calls.RegisterKafkaDeprovisionJob = append(mock.calls.RegisterKafkaDeprovisionJob, callInfo)
	mock.lockRegisterKafkaDeprovisionJob.Unlock()
	return mock.RegisterKafkaDeprovisionJobFunc(ctx, id)
}

// RegisterKafkaDeprovisionJobCalls gets all the calls that were made to RegisterKafkaDeprovisionJob.
// Check the length with:
//     len(mockedKafkaService.RegisterKafkaDeprovisionJobCalls())
func (mock *KafkaServiceMock) RegisterKafkaDeprovisionJobCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockRegisterKafkaDeprovisionJob.RLock()
	calls = mock.calls.RegisterKafkaDeprovisionJob
	mock.lockRegisterKafkaDeprovisionJob.RUnlock()
	return calls
}

// RegisterKafkaJob calls RegisterKafkaJobFunc.
func (mock *KafkaServiceMock) RegisterKafkaJob(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
	if mock.RegisterKafkaJobFunc == nil {
		panic("KafkaServiceMock.RegisterKafkaJobFunc: method is nil but KafkaService.RegisterKafkaJob was just called")
	}
	callInfo := struct {
		KafkaRequest *api.KafkaRequest
	}{
		KafkaRequest: kafkaRequest,
	}
	mock.lockRegisterKafkaJob.Lock()
	mock.calls.RegisterKafkaJob = append(mock.calls.RegisterKafkaJob, callInfo)
	mock.lockRegisterKafkaJob.Unlock()
	return mock.RegisterKafkaJobFunc(kafkaRequest)
}

// RegisterKafkaJobCalls gets all the calls that were made to RegisterKafkaJob.
// Check the length with:
//     len(mockedKafkaService.RegisterKafkaJobCalls())
func (mock *KafkaServiceMock) RegisterKafkaJobCalls() []struct {
	KafkaRequest *api.KafkaRequest
} {
	var calls []struct {
		KafkaRequest *api.KafkaRequest
	}
	mock.lockRegisterKafkaJob.RLock()
	calls = mock.calls.RegisterKafkaJob
	mock.lockRegisterKafkaJob.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *KafkaServiceMock) Update(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
	if mock.UpdateFunc == nil {
		panic("KafkaServiceMock.UpdateFunc: method is nil but KafkaService.Update was just called")
	}
	callInfo := struct {
		KafkaRequest *api.KafkaRequest
	}{
		KafkaRequest: kafkaRequest,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(kafkaRequest)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedKafkaService.UpdateCalls())
func (mock *KafkaServiceMock) UpdateCalls() []struct {
	KafkaRequest *api.KafkaRequest
} {
	var calls []struct {
		KafkaRequest *api.KafkaRequest
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// UpdateStatus calls UpdateStatusFunc.
func (mock *KafkaServiceMock) UpdateStatus(id string, status constants.KafkaStatus) (bool, *errors.ServiceError) {
	if mock.UpdateStatusFunc == nil {
		panic("KafkaServiceMock.UpdateStatusFunc: method is nil but KafkaService.UpdateStatus was just called")
	}
	callInfo := struct {
		ID     string
		Status constants.KafkaStatus
	}{
		ID:     id,
		Status: status,
	}
	mock.lockUpdateStatus.Lock()
	mock.calls.UpdateStatus = append(mock.calls.UpdateStatus, callInfo)
	mock.lockUpdateStatus.Unlock()
	return mock.UpdateStatusFunc(id, status)
}

// UpdateStatusCalls gets all the calls that were made to UpdateStatus.
// Check the length with:
//     len(mockedKafkaService.UpdateStatusCalls())
func (mock *KafkaServiceMock) UpdateStatusCalls() []struct {
	ID     string
	Status constants.KafkaStatus
} {
	var calls []struct {
		ID     string
		Status constants.KafkaStatus
	}
	mock.lockUpdateStatus.RLock()
	calls = mock.calls.UpdateStatus
	mock.lockUpdateStatus.RUnlock()
	return calls
}
