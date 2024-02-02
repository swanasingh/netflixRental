// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	movie "netflixRental/internal/models/movie"

	mock "github.com/stretchr/testify/mock"
)

// MovieRepository is an autogenerated mock type for the MovieRepository type
type MovieRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: criteria
func (_m *MovieRepository) Get(criteria movie.Criteria) []movie.Movie {
	ret := _m.Called(criteria)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []movie.Movie
	if rf, ok := ret.Get(0).(func(movie.Criteria) []movie.Movie); ok {
		r0 = rf(criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movie.Movie)
		}
	}

	return r0
}

// GetMovieDetails provides a mock function with given fields: id
func (_m *MovieRepository) GetMovieDetails(id int) (movie.Movie, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetMovieDetails")
	}

	var r0 movie.Movie
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (movie.Movie, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) movie.Movie); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(movie.Movie)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMovieRepository creates a new instance of MovieRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMovieRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MovieRepository {
	mock := &MovieRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
