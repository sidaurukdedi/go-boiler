// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongodb "github.com/sidaurukdedi/go-boiler/pkg/mongodb"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Connect provides a mock function with given fields: ctx
func (_m *Client) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database provides a mock function with given fields: name, opts
func (_m *Client) Database(name string, opts ...*options.DatabaseOptions) mongodb.Database {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongodb.Database
	if rf, ok := ret.Get(0).(func(string, ...*options.DatabaseOptions) mongodb.Database); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongodb.Database)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields: ctx
func (_m *Client) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
