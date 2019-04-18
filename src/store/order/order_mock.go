// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package order

import (
	"datcom/backend/src/store/item"
	"sync"
)

var (
	lockServiceMockAdd    sync.RWMutex
	lockServiceMockDelete sync.RWMutex
	lockServiceMockExist  sync.RWMutex
	lockServiceMockGet    sync.RWMutex
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
//             AddFunc: func(o *Order) error {
// 	               panic("mock out the Add method")
//             },
//             DeleteFunc: func(o *Order) error {
// 	               panic("mock out the Delete method")
//             },
//             ExistFunc: func(o *Order) bool {
// 	               panic("mock out the Exist method")
//             },
//             GetFunc: func(userID int) ([]*item.Item, error) {
// 	               panic("mock out the Get method")
//             },
//         }
//
//         // use mockedService in code that requires Service
//         // and then make assertions.
//
//     }
type ServiceMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(o *Order) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(o *Order) error

	// ExistFunc mocks the Exist method.
	ExistFunc func(o *Order) bool

	// GetFunc mocks the Get method.
	GetFunc func(userID int) ([]*item.Item, error)

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// O is the o argument value.
			O *Order
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// O is the o argument value.
			O *Order
		}
		// Exist holds details about calls to the Exist method.
		Exist []struct {
			// O is the o argument value.
			O *Order
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// UserID is the userID argument value.
			UserID int
		}
	}
}

// Add calls AddFunc.
func (mock *ServiceMock) Add(o *Order) error {
	if mock.AddFunc == nil {
		panic("ServiceMock.AddFunc: method is nil but Service.Add was just called")
	}
	callInfo := struct {
		O *Order
	}{
		O: o,
	}
	lockServiceMockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	lockServiceMockAdd.Unlock()
	return mock.AddFunc(o)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedService.AddCalls())
func (mock *ServiceMock) AddCalls() []struct {
	O *Order
} {
	var calls []struct {
		O *Order
	}
	lockServiceMockAdd.RLock()
	calls = mock.calls.Add
	lockServiceMockAdd.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ServiceMock) Delete(o *Order) error {
	if mock.DeleteFunc == nil {
		panic("ServiceMock.DeleteFunc: method is nil but Service.Delete was just called")
	}
	callInfo := struct {
		O *Order
	}{
		O: o,
	}
	lockServiceMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockServiceMockDelete.Unlock()
	return mock.DeleteFunc(o)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedService.DeleteCalls())
func (mock *ServiceMock) DeleteCalls() []struct {
	O *Order
} {
	var calls []struct {
		O *Order
	}
	lockServiceMockDelete.RLock()
	calls = mock.calls.Delete
	lockServiceMockDelete.RUnlock()
	return calls
}

// Exist calls ExistFunc.
func (mock *ServiceMock) Exist(o *Order) bool {
	if mock.ExistFunc == nil {
		panic("ServiceMock.ExistFunc: method is nil but Service.Exist was just called")
	}
	callInfo := struct {
		O *Order
	}{
		O: o,
	}
	lockServiceMockExist.Lock()
	mock.calls.Exist = append(mock.calls.Exist, callInfo)
	lockServiceMockExist.Unlock()
	return mock.ExistFunc(o)
}

// ExistCalls gets all the calls that were made to Exist.
// Check the length with:
//     len(mockedService.ExistCalls())
func (mock *ServiceMock) ExistCalls() []struct {
	O *Order
} {
	var calls []struct {
		O *Order
	}
	lockServiceMockExist.RLock()
	calls = mock.calls.Exist
	lockServiceMockExist.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *ServiceMock) Get(userID int) ([]*item.Item, error) {
	if mock.GetFunc == nil {
		panic("ServiceMock.GetFunc: method is nil but Service.Get was just called")
	}
	callInfo := struct {
		UserID int
	}{
		UserID: userID,
	}
	lockServiceMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockServiceMockGet.Unlock()
	return mock.GetFunc(userID)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedService.GetCalls())
func (mock *ServiceMock) GetCalls() []struct {
	UserID int
} {
	var calls []struct {
		UserID int
	}
	lockServiceMockGet.RLock()
	calls = mock.calls.Get
	lockServiceMockGet.RUnlock()
	return calls
}
