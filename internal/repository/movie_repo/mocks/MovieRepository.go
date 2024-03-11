// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	movie "netflixRental/internal/models/movie"

	mock "github.com/stretchr/testify/mock"
)

// MovieRepository is an autogenerated mock type for the MovieRepository type
type MovieRepository struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: order
func (_m *MovieRepository) CreateOrder(order movie.OrderPayload) error {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(movie.OrderPayload) error); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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

// GetInvoice provides a mock function with given fields: orderId
func (_m *MovieRepository) GetInvoice(orderId int) ([]movie.Invoice, movie.User, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetInvoice")
	}

	var r0 []movie.Invoice
	var r1 movie.User
	var r2 error
	if rf, ok := ret.Get(0).(func(int) ([]movie.Invoice, movie.User, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(int) []movie.Invoice); ok {
		r0 = rf(orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]movie.Invoice)
		}
	}

	if rf, ok := ret.Get(1).(func(int) movie.User); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Get(1).(movie.User)
	}

	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(orderId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
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

// GetUserDetails provides a mock function with given fields: userId
func (_m *MovieRepository) GetUserDetails(userId int) (movie.User, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserDetails")
	}

	var r0 movie.User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (movie.User, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int) movie.User); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(movie.User)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveCartData provides a mock function with given fields: cartItem
func (_m *MovieRepository) SaveCartData(cartItem movie.CartItem) error {
	ret := _m.Called(cartItem)

	if len(ret) == 0 {
		panic("no return value specified for SaveCartData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(movie.CartItem) error); ok {
		r0 = rf(cartItem)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ViewCart provides a mock function with given fields: user_id
func (_m *MovieRepository) ViewCart(user_id int) []movie.Movie {
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
