package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api/openapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api/presenters"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/auth"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/client/observatorium"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/constants"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/services"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/test"
	utils "github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/test/common"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/test/mocks"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	mockResourceNamespace = "my-kafka-namespace"
	mockKafkaClusterName  = "my-cluster"
)

func TestObservatorium_ResourceStateMetric(t *testing.T) {
	ocmServerBuilder := mocks.NewMockConfigurableServerBuilder()
	ocmServer := ocmServerBuilder.Build()
	defer ocmServer.Close()

	// start servers
	h, _, teardown := test.RegisterIntegration(t, ocmServer)
	defer teardown()
	h.Env().Config.ObservabilityConfiguration.EnableMock = true
	err := h.Env().LoadClients()
	Expect(err).NotTo(HaveOccurred(), "Error occurred when loading clients: %v", err)

	defer h.Reset()
	service := services.NewObservatoriumService(h.Env().Clients.Observatorium, h.Env().Services.Kafka)
	kafkaState, err := service.GetKafkaState(mockKafkaClusterName, mockResourceNamespace)
	Expect(err).NotTo(HaveOccurred(), "Error getting kafka state:  %v", err)
	Expect(kafkaState.State).NotTo(BeEmpty(), "Should return state")
}

func TestObservatorium_GetMetrics(t *testing.T) {

	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, client, teardown := test.RegisterIntegration(t, ocmServer)
	defer teardown()

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)
	k := openapi.KafkaRequestPayload{
		Region:        mocks.MockCluster.Region().ID(),
		CloudProvider: mocks.MockCluster.CloudProvider().ID(),
		Name:          mockKafkaName,
		MultiAz:       testMultiAZ,
	}

	seedKafka, _, err := client.DefaultApi.CreateKafka(ctx, true, k)
	if err != nil {
		t.Fatalf("failed to create seeded kafka request: %s", err.Error())
	}

	h.Env().Config.ObservabilityConfiguration.EnableMock = true
	ctx = auth.SetUsernameContext(context.TODO(), account.Username())
	err = h.Env().LoadClients()
	Expect(err).NotTo(HaveOccurred(), "Error occurred when loading clients: %v", err)
	service := services.NewObservatoriumService(h.Env().Clients.Observatorium, h.Env().Services.Kafka)
	metricsList := &observatorium.KafkaMetrics{}
	q := observatorium.MetricsReqParams{}
	q.ResultType = observatorium.RangeQuery
	q.FillDefaults()
	_, err = service.GetMetricsByKafkaId(ctx, metricsList, seedKafka.Id, q)
	Expect(err).NotTo(HaveOccurred(), "Error getting kafka metrics:  %v", err)
	Expect(len(*metricsList)).NotTo(Equal(0), "Should return length greater then zero")
}

func TestObservatorium_GetMetricsByQueryRange(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, client, teardown := test.RegisterIntegration(t, ocmServer)
	defer teardown()
	h.Env().Config.ObservabilityConfiguration.EnableMock = true

	clusterID, getClusterErr := utils.GetRunningOsdClusterID(h, t)
	if getClusterErr != nil {
		t.Fatalf("Failed to retrieve cluster details from persisted .json file: %v", getClusterErr)
	}
	if clusterID == "" {
		panic("No cluster found")
	}

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)
	k := openapi.KafkaRequestPayload{
		Region:        mocks.MockCluster.Region().ID(),
		CloudProvider: mocks.MockCluster.CloudProvider().ID(),
		Name:          mockKafkaName,
		MultiAz:       testMultiAZ,
	}

	seedKafka, _, err := client.DefaultApi.CreateKafka(ctx, true, k)
	if err != nil {
		t.Fatalf("failed to create seeded kafka request: %s", err.Error())
	}

	var foundKafka openapi.KafkaRequest
	_ = wait.PollImmediate(kafkaCheckInterval, kafkaReadyTimeout, func() (done bool, err error) {
		foundKafka, _, err = client.DefaultApi.GetKafkaById(ctx, seedKafka.Id)
		if err != nil {
			return true, err
		}
		return foundKafka.Status == constants.KafkaRequestStatusReady.String(), nil
	})

	// 200 OK
	kafka, resp, err := client.DefaultApi.GetKafkaById(ctx, seedKafka.Id)
	Expect(err).NotTo(HaveOccurred(), "Error occurred when attempting to get kafka request:  %v", err)
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
	Expect(kafka.Id).NotTo(BeEmpty(), "Expected ID assigned on creation")
	Expect(kafka.Kind).To(Equal(presenters.KindKafka))
	Expect(kafka.Href).To(Equal(fmt.Sprintf("/api/managed-services-api/v1/kafkas/%s", kafka.Id)))
	Expect(kafka.Region).To(Equal(mocks.MockCluster.Region().ID()))
	Expect(kafka.CloudProvider).To(Equal(mocks.MockCluster.CloudProvider().ID()))
	Expect(kafka.Name).To(Equal(mockKafkaName))
	Expect(kafka.Status).To(Equal(constants.KafkaRequestStatusReady.String()))

	// 404 Not Found
	kafka, resp, _ = client.DefaultApi.GetKafkaById(ctx, fmt.Sprintf("not-%s", seedKafka.Id))
	Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

	// different account but same org, should be able to read the Kafka cluster
	acc := h.NewRandAccount()
	context := h.NewAuthenticatedContext(acc)
	kafka, _, _ = client.DefaultApi.GetKafkaById(context, seedKafka.Id)
	Expect(kafka.Id).NotTo(BeEmpty())
	h.Env().Config.ObservabilityConfiguration.EnableMock = true
	err = h.Env().LoadClients()
	Expect(err).NotTo(HaveOccurred(), "Error occurred when loading clients: %v", err)
	filters := openapi.GetMetricsByRangeQueryOpts{}
	metrics, resp, err := client.DefaultApi.GetMetricsByRangeQuery(context, kafka.Id, 5, 30, &filters)
	Expect(err).NotTo(HaveOccurred(), "Error occurred when attempting to get metrics data:  %v", err)
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
	Expect(len(metrics.Items)).NotTo(Equal(0))
	deleteTestKafka(ctx, client, foundKafka.Id)
}
func TestObservatorium_GetMetricsByQueryInstant(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, client, teardown := test.RegisterIntegration(t, ocmServer)
	defer teardown()
	h.Env().Config.ObservabilityConfiguration.EnableMock = true

	clusterID, getClusterErr := utils.GetRunningOsdClusterID(h, t)
	if getClusterErr != nil {
		t.Fatalf("Failed to retrieve cluster details from persisted .json file: %v", getClusterErr)
	}
	if clusterID == "" {
		panic("No cluster found")
	}

	account := h.NewRandAccount()
	ctx := h.NewAuthenticatedContext(account)
	k := openapi.KafkaRequestPayload{
		Region:        mocks.MockCluster.Region().ID(),
		CloudProvider: mocks.MockCluster.CloudProvider().ID(),
		Name:          mockKafkaName,
		MultiAz:       testMultiAZ,
	}

	seedKafka, _, err := client.DefaultApi.CreateKafka(ctx, true, k)
	if err != nil {
		t.Fatalf("failed to create seeded kafka request: %s", err.Error())
	}

	var foundKafka openapi.KafkaRequest
	_ = wait.PollImmediate(kafkaCheckInterval, kafkaReadyTimeout, func() (done bool, err error) {
		foundKafka, _, err = client.DefaultApi.GetKafkaById(ctx, seedKafka.Id)
		if err != nil {
			return true, err
		}
		return foundKafka.Status == constants.KafkaRequestStatusReady.String(), nil
	})

	// 200 OK
	kafka, resp, err := client.DefaultApi.GetKafkaById(ctx, seedKafka.Id)
	Expect(err).NotTo(HaveOccurred(), "Error occurred when attempting to get kafka request:  %v", err)
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
	Expect(kafka.Id).NotTo(BeEmpty(), "Expected ID assigned on creation")
	Expect(kafka.Kind).To(Equal(presenters.KindKafka))
	Expect(kafka.Href).To(Equal(fmt.Sprintf("/api/managed-services-api/v1/kafkas/%s", kafka.Id)))
	Expect(kafka.Region).To(Equal(mocks.MockCluster.Region().ID()))
	Expect(kafka.CloudProvider).To(Equal(mocks.MockCluster.CloudProvider().ID()))
	Expect(kafka.Name).To(Equal(mockKafkaName))
	Expect(kafka.Status).To(Equal(constants.KafkaRequestStatusReady.String()))

	// 404 Not Found
	kafka, resp, _ = client.DefaultApi.GetKafkaById(ctx, fmt.Sprintf("not-%s", seedKafka.Id))
	Expect(resp.StatusCode).To(Equal(http.StatusNotFound))

	// different account but same org, should be able to read the Kafka cluster
	acc := h.NewRandAccount()
	context := h.NewAuthenticatedContext(acc)
	kafka, _, _ = client.DefaultApi.GetKafkaById(context, seedKafka.Id)
	Expect(kafka.Id).NotTo(BeEmpty())
	h.Env().Config.ObservabilityConfiguration.EnableMock = true
	err = h.Env().LoadClients()
	Expect(err).NotTo(HaveOccurred(), "Error occurred when loading clients: %v", err)
	filters := openapi.GetMetricsByInstantQueryOpts{}
	metrics, resp, err := client.DefaultApi.GetMetricsByInstantQuery(context, kafka.Id, &filters)
	Expect(err).NotTo(HaveOccurred(), "Error occurred when attempting to get metrics data:  %v", err)
	Expect(resp.StatusCode).To(Equal(http.StatusOK))
	Expect(len(metrics.Items)).NotTo(Equal(0))
	deleteTestKafka(ctx, client, foundKafka.Id)
}
