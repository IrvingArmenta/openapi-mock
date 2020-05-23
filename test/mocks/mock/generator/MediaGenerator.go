// Code generated by mockery v1.0.0. DO NOT EDIT.

package generatormock

import (
	context "context"
	mock "github.com/stretchr/testify/mock"
	"swagger-mock/internal/openapi/generator/data"

	openapi3 "github.com/getkin/kin-openapi/openapi3"
)

// MediaGenerator is an autogenerated mock type for the MediaGenerator type
type MediaGenerator struct {
	mock.Mock
}

// GenerateData provides a mock function with given fields: ctx, mediaType
func (_m *MediaGenerator) GenerateData(ctx context.Context, mediaType *openapi3.MediaType) (data.Data, error) {
	ret := _m.Called(ctx, mediaType)

	var r0 data.Data
	if rf, ok := ret.Get(0).(func(context.Context, *openapi3.MediaType) data.Data); ok {
		r0 = rf(ctx, mediaType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(data.Data)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *openapi3.MediaType) error); ok {
		r1 = rf(ctx, mediaType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
