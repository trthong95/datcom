// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package item

import (
	"sync"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

var (
	lockServiceMockAdd            sync.RWMutex
	lockServiceMockCheckItemExist sync.RWMutex
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             AddFunc: func(in1 *domain.Item) (*models.Item, error) {
// 	               panic("mock out the Add method")
//             },
//             CheckItemExistFunc: func(in1 *domain.Item) (bool, error) {
// 	               panic("mock out the CheckItemExist method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(in1 *domain.Item) (*models.Item, error)

	// CheckItemExistFunc mocks the CheckItemExist method.
	CheckItemExistFunc func(in1 *domain.Item) (bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// In1 is the in1 argument value.
			In1 *domain.Item
		}
		// CheckItemExist holds details about calls to the CheckItemExist method.
		CheckItemExist []struct {
			// In1 is the in1 argument value.
			In1 *domain.Item
		}
	}
}

// Add calls AddFunc.
func (mock *ServiceMock) Add(in1 *domain.Item) (*models.Item, error) {
	if mock.AddFunc == nil {
		panic("ServiceMock.AddFunc: method is nil but Service.Add was just called")
	}
	callInfo := struct {
		In1 *domain.Item
	}{
		In1: in1,
	}
	lockServiceMockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	lockServiceMockAdd.Unlock()
	return mock.AddFunc(in1)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedService.AddCalls())
func (mock *ServiceMock) AddCalls() []struct {
	In1 *domain.Item
} {
	var calls []struct {
		In1 *domain.Item
	}
	lockServiceMockAdd.RLock()
	calls = mock.calls.Add
	lockServiceMockAdd.RUnlock()
	return calls
}

// CheckItemExist calls CheckItemExistFunc.
func (mock *ServiceMock) CheckItemExist(in1 *domain.Item) (bool, error) {
	if mock.CheckItemExistFunc == nil {
		panic("ServiceMock.CheckItemExistFunc: method is nil but Service.CheckItemExist was just called")
	}
	callInfo := struct {
		In1 *domain.Item
	}{
		In1: in1,
	}
	lockServiceMockCheckItemExist.Lock()
	mock.calls.CheckItemExist = append(mock.calls.CheckItemExist, callInfo)
	lockServiceMockCheckItemExist.Unlock()
	return mock.CheckItemExistFunc(in1)
}

// CheckItemExistCalls gets all the calls that were made to CheckItemExist.
// Check the length with:
//     len(mockedService.CheckItemExistCalls())
func (mock *ServiceMock) CheckItemExistCalls() []struct {
	In1 *domain.Item
} {
	var calls []struct {
		In1 *domain.Item
	}
	lockServiceMockCheckItemExist.RLock()
	calls = mock.calls.CheckItemExist
	lockServiceMockCheckItemExist.RUnlock()
	return calls
}