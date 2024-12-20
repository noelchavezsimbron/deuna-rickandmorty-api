// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	episode "deuna-rickandmorty-api/internal/episode"

	mock "github.com/stretchr/testify/mock"
)

// EpisodeGetterUseCaseMock is an autogenerated mock type for the episodeGetterUseCase type
type EpisodeGetterUseCaseMock struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *EpisodeGetterUseCaseMock) GetAll(ctx context.Context) ([]episode.Episode, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []episode.Episode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]episode.Episode, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []episode.Episode); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]episode.Episode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, ID
func (_m *EpisodeGetterUseCaseMock) GetByID(ctx context.Context, ID int64) (episode.Episode, error) {
	ret := _m.Called(ctx, ID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 episode.Episode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (episode.Episode, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) episode.Episode); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(episode.Episode)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEpisodeGetterUseCaseMock creates a new instance of EpisodeGetterUseCaseMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEpisodeGetterUseCaseMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *EpisodeGetterUseCaseMock {
	mock := &EpisodeGetterUseCaseMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}