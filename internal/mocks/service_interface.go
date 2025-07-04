// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	model "blog-api/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: post
func (_m *ServiceInterface) CreatePost(post *model.BlogPost) (string, error) {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for CreatePost")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.BlogPost) (string, error)); ok {
		return rf(post)
	}
	if rf, ok := ret.Get(0).(func(*model.BlogPost) string); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*model.BlogPost) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePostByID provides a mock function with given fields: id
func (_m *ServiceInterface) DeletePostByID(id uint) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePostByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPosts provides a mock function with no fields
func (_m *ServiceInterface) GetAllPosts() ([]*model.BlogPost, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllPosts")
	}

	var r0 []*model.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.BlogPost, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.BlogPost); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostByID provides a mock function with given fields: id
func (_m *ServiceInterface) GetPostByID(id uint) (*model.BlogPost, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPostByID")
	}

	var r0 *model.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*model.BlogPost, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *model.BlogPost); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePostByID provides a mock function with given fields: post
func (_m *ServiceInterface) UpdatePostByID(post *model.BlogPost) error {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePostByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.BlogPost) error); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServiceInterface creates a new instance of ServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceInterface {
	mock := &ServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
