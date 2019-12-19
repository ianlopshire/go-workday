// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	raas "github.com/ianlopshire/go-workday/pkg/raas"

	session "github.com/ianlopshire/go-workday/pkg/wd/session"

	url "net/url"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// Call provides a mock function with given fields: ctx, sess, def, query
func (_m *Backend) Call(ctx context.Context, sess *session.Session, def raas.Definition, query url.Values) (io.ReadCloser, error) {
	ret := _m.Called(ctx, sess, def, query)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, *session.Session, raas.Definition, url.Values) io.ReadCloser); ok {
		r0 = rf(ctx, sess, def, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *session.Session, raas.Definition, url.Values) error); ok {
		r1 = rf(ctx, sess, def, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
