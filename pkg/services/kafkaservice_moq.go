// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package services

import (
	"context"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/constants"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
	"sync"
)

var (
	lockKafkaServiceMockCreate             sync.RWMutex
	lockKafkaServiceMockDelete             sync.RWMutex
	lockKafkaServiceMockGet                sync.RWMutex
	lockKafkaServiceMockList               sync.RWMutex
	lockKafkaServiceMockListByStatus       sync.RWMutex
	lockKafkaServiceMockRegisterKafkaInSSO sync.RWMutex
	lockKafkaServiceMockRegisterKafkaJob   sync.RWMutex
	lockKafkaServiceMockUpdate             sync.RWMutex
	lockKafkaServiceMockUpdateStatus       sync.RWMutex
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
//             CreateFunc: func(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the Create method")
//             },
//             DeleteFunc: func(ctx context.Context, id string) *errors.ServiceError {
// 	               panic("mock out the Delete method")
//             },
//             GetFunc: func(id string) (*api.KafkaRequest, *errors.ServiceError) {
// 	               panic("mock out the Get method")
//             },
//             ListFunc: func(ctx context.Context, listArgs *ListArguments) (api.KafkaList, *api.PagingMeta, *errors.ServiceError) {
// 	               panic("mock out the List method")
//             },
//             ListByStatusFunc: func(status constants.KafkaStatus) ([]*api.KafkaRequest, *errors.ServiceError) {
// 	               panic("mock out the ListByStatus method")
//             },
//             RegisterKafkaInSSOFunc: func(ctx context.Context, kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the RegisterKafkaInSSO method")
//             },
//             RegisterKafkaJobFunc: func(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the RegisterKafkaJob method")
//             },
//             UpdateFunc: func(kafkaRequest *api.KafkaRequest) *errors.ServiceError {
// 	               panic("mock out the Update method")
//             },
//             UpdateStatusFunc: func(id string, status constants.KafkaStatus) *errors.ServiceError {
// 	               panic("mock out the UpdateStatus method")
//             },
//         }
//
//         // use mockedKafkaService in code that requires KafkaService
//         // and then make assertions.
//
//     }
type KafkaServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id string) *errors.ServiceError

	// GetFunc mocks the Get method.
	GetFunc func(id string) (*api.KafkaRequest, *errors.ServiceError)

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, listArgs *ListArguments) (api.KafkaList, *api.PagingMeta, *errors.ServiceError)

	// ListByStatusFunc mocks the ListByStatus method.
	ListByStatusFunc func(status constants.KafkaStatus) ([]*api.KafkaRequest, *errors.ServiceError)

	// RegisterKafkaInSSOFunc mocks the RegisterKafkaInSSO method.
	RegisterKafkaInSSOFunc func(ctx context.Context, kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// RegisterKafkaJobFunc mocks the RegisterKafkaJob method.
	RegisterKafkaJobFunc func(kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// UpdateFunc mocks the Update method.
	UpdateFunc func(kafkaRequest *api.KafkaRequest) *errors.ServiceError

	// UpdateStatusFunc mocks the UpdateStatus method.
	UpdateStatusFunc func(id string, status constants.KafkaStatus) *errors.ServiceError

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// KafkaRequest is the kafkaRequest argument value.
			KafkaRequest *api.KafkaRequest
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// Get holds details about calls to the Get method.
		Get []struct {
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
		// RegisterKafkaInSSO holds details about calls to the RegisterKafkaInSSO method.
		RegisterKafkaInSSO []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// KafkaRequest is the kafkaRequest argument value.
			KafkaRequest *api.KafkaRequest
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
	lockKafkaServiceMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockKafkaServiceMockCreate.Unlock()
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
	lockKafkaServiceMockCreate.RLock()
	calls = mock.calls.Create
	lockKafkaServiceMockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *KafkaServiceMock) Delete(ctx context.Context, id string) *errors.ServiceError {
	if mock.DeleteFunc == nil {
		panic("KafkaServiceMock.DeleteFunc: method is nil but KafkaService.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	lockKafkaServiceMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockKafkaServiceMockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedKafkaService.DeleteCalls())
func (mock *KafkaServiceMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	lockKafkaServiceMockDelete.RLock()
	calls = mock.calls.Delete
	lockKafkaServiceMockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *KafkaServiceMock) Get(id string) (*api.KafkaRequest, *errors.ServiceError) {
	if mock.GetFunc == nil {
		panic("KafkaServiceMock.GetFunc: method is nil but KafkaService.Get was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockKafkaServiceMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockKafkaServiceMockGet.Unlock()
	return mock.GetFunc(id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedKafkaService.GetCalls())
func (mock *KafkaServiceMock) GetCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockKafkaServiceMockGet.RLock()
	calls = mock.calls.Get
	lockKafkaServiceMockGet.RUnlock()
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
	lockKafkaServiceMockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	lockKafkaServiceMockList.Unlock()
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
	lockKafkaServiceMockList.RLock()
	calls = mock.calls.List
	lockKafkaServiceMockList.RUnlock()
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
	lockKafkaServiceMockListByStatus.Lock()
	mock.calls.ListByStatus = append(mock.calls.ListByStatus, callInfo)
	lockKafkaServiceMockListByStatus.Unlock()
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
	lockKafkaServiceMockListByStatus.RLock()
	calls = mock.calls.ListByStatus
	lockKafkaServiceMockListByStatus.RUnlock()
	return calls
}

// RegisterKafkaInSSO calls RegisterKafkaInSSOFunc.
func (mock *KafkaServiceMock) RegisterKafkaInSSO(ctx context.Context, kafkaRequest *api.KafkaRequest) *errors.ServiceError {
	if mock.RegisterKafkaInSSOFunc == nil {
		panic("KafkaServiceMock.RegisterKafkaInSSOFunc: method is nil but KafkaService.RegisterKafkaInSSO was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		KafkaRequest *api.KafkaRequest
	}{
		Ctx:          ctx,
		KafkaRequest: kafkaRequest,
	}
	lockKafkaServiceMockRegisterKafkaInSSO.Lock()
	mock.calls.RegisterKafkaInSSO = append(mock.calls.RegisterKafkaInSSO, callInfo)
	lockKafkaServiceMockRegisterKafkaInSSO.Unlock()
	return mock.RegisterKafkaInSSOFunc(ctx, kafkaRequest)
}

// RegisterKafkaInSSOCalls gets all the calls that were made to RegisterKafkaInSSO.
// Check the length with:
//     len(mockedKafkaService.RegisterKafkaInSSOCalls())
func (mock *KafkaServiceMock) RegisterKafkaInSSOCalls() []struct {
	Ctx          context.Context
	KafkaRequest *api.KafkaRequest
} {
	var calls []struct {
		Ctx          context.Context
		KafkaRequest *api.KafkaRequest
	}
	lockKafkaServiceMockRegisterKafkaInSSO.RLock()
	calls = mock.calls.RegisterKafkaInSSO
	lockKafkaServiceMockRegisterKafkaInSSO.RUnlock()
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
	lockKafkaServiceMockRegisterKafkaJob.Lock()
	mock.calls.RegisterKafkaJob = append(mock.calls.RegisterKafkaJob, callInfo)
	lockKafkaServiceMockRegisterKafkaJob.Unlock()
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
	lockKafkaServiceMockRegisterKafkaJob.RLock()
	calls = mock.calls.RegisterKafkaJob
	lockKafkaServiceMockRegisterKafkaJob.RUnlock()
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
	lockKafkaServiceMockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	lockKafkaServiceMockUpdate.Unlock()
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
	lockKafkaServiceMockUpdate.RLock()
	calls = mock.calls.Update
	lockKafkaServiceMockUpdate.RUnlock()
	return calls
}

// UpdateStatus calls UpdateStatusFunc.
func (mock *KafkaServiceMock) UpdateStatus(id string, status constants.KafkaStatus) *errors.ServiceError {
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
	lockKafkaServiceMockUpdateStatus.Lock()
	mock.calls.UpdateStatus = append(mock.calls.UpdateStatus, callInfo)
	lockKafkaServiceMockUpdateStatus.Unlock()
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
	lockKafkaServiceMockUpdateStatus.RLock()
	calls = mock.calls.UpdateStatus
	lockKafkaServiceMockUpdateStatus.RUnlock()
	return calls
}
