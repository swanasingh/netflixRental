// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	movie "netflixRental/internal/models/movie"

	mock "github.com/stretchr/testify/mock"
)

// MovieService is an autogenerated mock type for the MovieService type
type MovieService struct {
	mock.Mock
}

// AddToCart provides a mock function with given fields: cartItem
func (_m *MovieService) AddToCart(cartItem movie.CartItem) error {
	ret := _m.Called(cartItem)

	if len(ret) == 0 {
		panic("no return value specified for AddToCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(movie.CartItem) error); ok {
		r0 = rf(cartItem)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: criteria
func (_m *MovieService) Get(criteria movie.Criteria) []movie.Movie {
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
func (_m *MovieService) GetMovieDetails(id int) (movie.Movie, error) {
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

// ViewCart provides a mock function with given fields: user_id
func (_m *MovieService) ViewCart(user_id int) []movie.Movie {
	ret := _m.Called(user_id)

	if len(ret) == 0 {
		panic("no return value specified for ViewCart")
	}

	var r0 []movie.Movie
	if rf, ok := ret.Get(0).(func(int) []movie.Movie); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movie.Movie)
		}
	}

	return r0
}

// NewMovieService creates a new instance of MovieService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMovieService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MovieService {
	mock := &MovieService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
