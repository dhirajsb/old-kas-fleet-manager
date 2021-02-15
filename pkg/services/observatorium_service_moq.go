// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"context"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/client/observatorium"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
	"sync"
)

// Ensure, that ObservatoriumServiceMock does implement ObservatoriumService.
// If this is not the case, regenerate this file with moq.
var _ ObservatoriumService = &ObservatoriumServiceMock{}

// ObservatoriumServiceMock is a mock implementation of ObservatoriumService.
//
//     func TestSomethingThatUsesObservatoriumService(t *testing.T) {
//
//         // make and configure a mocked ObservatoriumService
//         mockedObservatoriumService := &ObservatoriumServiceMock{
//             GetKafkaStateFunc: func(name string, namespaceName string) (observatorium.KafkaState, error) {
// 	               panic("mock out the GetKafkaState method")
//             },
//             GetMetricsByKafkaIdFunc: func(ctx context.Context, csMetrics *observatorium.KafkaMetrics, id string, query observatorium.RangeQuery) (string, *errors.ServiceError) {
// 	               panic("mock out the GetMetricsByKafkaId method")
//             },
//         }
//
//         // use mockedObservatoriumService in code that requires ObservatoriumService
//         // and then make assertions.
//
//     }
type ObservatoriumServiceMock struct {
	// GetKafkaStateFunc mocks the GetKafkaState method.
	GetKafkaStateFunc func(name string, namespaceName string) (observatorium.KafkaState, error)

	// GetMetricsByKafkaIdFunc mocks the GetMetricsByKafkaId method.
	GetMetricsByKafkaIdFunc func(ctx context.Context, csMetrics *observatorium.KafkaMetrics, id string, query observatorium.MetricsReqParams) (string, *errors.ServiceError)

	// calls tracks calls to the methods.
	calls struct {
		// GetKafkaState holds details about calls to the GetKafkaState method.
		GetKafkaState []struct {
			// Name is the name argument value.
			Name string
			// NamespaceName is the namespaceName argument value.
			NamespaceName string
		}
		// GetMetricsByKafkaId holds details about calls to the GetMetricsByKafkaId method.
		GetMetricsByKafkaId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CsMetrics is the csMetrics argument value.
			CsMetrics *observatorium.KafkaMetrics
			// ID is the id argument value.
			ID string
			// Query is the query argument value.
			Query observatorium.MetricsReqParams
		}
	}
	lockGetKafkaState       sync.RWMutex
	lockGetMetricsByKafkaId sync.RWMutex
}

// GetKafkaState calls GetKafkaStateFunc.
func (mock *ObservatoriumServiceMock) GetKafkaState(name string, namespaceName string) (observatorium.KafkaState, error) {
	if mock.GetKafkaStateFunc == nil {
		panic("ObservatoriumServiceMock.GetKafkaStateFunc: method is nil but ObservatoriumService.GetKafkaState was just called")
	}
	callInfo := struct {
		Name          string
		NamespaceName string
	}{
		Name:          name,
		NamespaceName: namespaceName,
	}
	mock.lockGetKafkaState.Lock()
	mock.calls.GetKafkaState = append(mock.calls.GetKafkaState, callInfo)
	mock.lockGetKafkaState.Unlock()
	return mock.GetKafkaStateFunc(name, namespaceName)
}

// GetKafkaStateCalls gets all the calls that were made to GetKafkaState.
// Check the length with:
//     len(mockedObservatoriumService.GetKafkaStateCalls())
func (mock *ObservatoriumServiceMock) GetKafkaStateCalls() []struct {
	Name          string
	NamespaceName string
} {
	var calls []struct {
		Name          string
		NamespaceName string
	}
	mock.lockGetKafkaState.RLock()
	calls = mock.calls.GetKafkaState
	mock.lockGetKafkaState.RUnlock()
	return calls
}

// GetMetricsByKafkaId calls GetMetricsByKafkaIdFunc.
func (mock *ObservatoriumServiceMock) GetMetricsByKafkaId(ctx context.Context, csMetrics *observatorium.KafkaMetrics, id string, query observatorium.MetricsReqParams) (string, *errors.ServiceError) {
	if mock.GetMetricsByKafkaIdFunc == nil {
		panic("ObservatoriumServiceMock.GetMetricsByKafkaIdFunc: method is nil but ObservatoriumService.GetMetricsByKafkaId was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		CsMetrics *observatorium.KafkaMetrics
		ID        string
		Query     observatorium.MetricsReqParams
	}{
		Ctx:       ctx,
		CsMetrics: csMetrics,
		ID:        id,
		Query:     query,
	}
	mock.lockGetMetricsByKafkaId.Lock()
	mock.calls.GetMetricsByKafkaId = append(mock.calls.GetMetricsByKafkaId, callInfo)
	mock.lockGetMetricsByKafkaId.Unlock()
	return mock.GetMetricsByKafkaIdFunc(ctx, csMetrics, id, query)
}

// GetMetricsByKafkaIdCalls gets all the calls that were made to GetMetricsByKafkaId.
// Check the length with:
//     len(mockedObservatoriumService.GetMetricsByKafkaIdCalls())
func (mock *ObservatoriumServiceMock) GetMetricsByKafkaIdCalls() []struct {
	Ctx       context.Context
	CsMetrics *observatorium.KafkaMetrics
	ID        string
	Query     observatorium.MetricsReqParams
} {
	var calls []struct {
		Ctx       context.Context
		CsMetrics *observatorium.KafkaMetrics
		ID        string
		Query     observatorium.MetricsReqParams
	}
	mock.lockGetMetricsByKafkaId.RLock()
	calls = mock.calls.GetMetricsByKafkaId
	mock.lockGetMetricsByKafkaId.RUnlock()
	return calls
}
