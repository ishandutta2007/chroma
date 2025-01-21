// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/chroma-core/chroma/go/pkg/sysdb/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// IDatabaseDb is an autogenerated mock type for the IDatabaseDb type
type IDatabaseDb struct {
	mock.Mock
}

// Delete provides a mock function with given fields: databaseID
func (_m *IDatabaseDb) Delete(databaseID string) error {
	ret := _m.Called(databaseID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(databaseID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAll provides a mock function with given fields:
func (_m *IDatabaseDb) DeleteAll() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeleteAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllDatabases provides a mock function with given fields:
func (_m *IDatabaseDb) GetAllDatabases() ([]*dbmodel.Database, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllDatabases")
	}

	var r0 []*dbmodel.Database
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*dbmodel.Database, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*dbmodel.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Database)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDatabases provides a mock function with given fields: tenantID, databaseName
func (_m *IDatabaseDb) GetDatabases(tenantID string, databaseName string) ([]*dbmodel.Database, error) {
	ret := _m.Called(tenantID, databaseName)

	if len(ret) == 0 {
		panic("no return value specified for GetDatabases")
	}

	var r0 []*dbmodel.Database
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*dbmodel.Database, error)); ok {
		return rf(tenantID, databaseName)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*dbmodel.Database); ok {
		r0 = rf(tenantID, databaseName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Database)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(tenantID, databaseName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: in
func (_m *IDatabaseDb) Insert(in *dbmodel.Database) error {
	ret := _m.Called(in)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.Database) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListDatabases provides a mock function with given fields: limit, offset, tenantID
func (_m *IDatabaseDb) ListDatabases(limit *int32, offset *int32, tenantID string) ([]*dbmodel.Database, error) {
	ret := _m.Called(limit, offset, tenantID)

	if len(ret) == 0 {
		panic("no return value specified for ListDatabases")
	}

	var r0 []*dbmodel.Database
	var r1 error
	if rf, ok := ret.Get(0).(func(*int32, *int32, string) ([]*dbmodel.Database, error)); ok {
		return rf(limit, offset, tenantID)
	}
	if rf, ok := ret.Get(0).(func(*int32, *int32, string) []*dbmodel.Database); ok {
		r0 = rf(limit, offset, tenantID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.Database)
		}
	}

	if rf, ok := ret.Get(1).(func(*int32, *int32, string) error); ok {
		r1 = rf(limit, offset, tenantID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIDatabaseDb creates a new instance of IDatabaseDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDatabaseDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDatabaseDb {
	mock := &IDatabaseDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
