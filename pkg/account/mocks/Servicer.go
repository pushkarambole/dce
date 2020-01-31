// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import account "github.com/Optum/dce/pkg/account"

import mock "github.com/stretchr/testify/mock"

// Servicer is an autogenerated mock type for the Servicer type
type Servicer struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *Servicer) Create(data *account.Account) (*account.Account, error) {
	ret := _m.Called(data)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(*account.Account) *account.Account); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*account.Account) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: data
func (_m *Servicer) Delete(data *account.Account) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ID
func (_m *Servicer) Get(ID string) (*account.Account, error) {
	ret := _m.Called(ID)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(string) *account.Account); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: query
func (_m *Servicer) List(query *account.Account) (*account.Accounts, error) {
	ret := _m.Called(query)

	var r0 *account.Accounts
	if rf, ok := ret.Get(0).(func(*account.Account) *account.Accounts); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Accounts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*account.Account) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: data
func (_m *Servicer) Save(data *account.Account) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ID, data
func (_m *Servicer) Update(ID string, data *account.Account) (*account.Account, error) {
	ret := _m.Called(ID, data)

	var r0 *account.Account
	if rf, ok := ret.Get(0).(func(string, *account.Account) *account.Account); ok {
		r0 = rf(ID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*account.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *account.Account) error); ok {
		r1 = rf(ID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
