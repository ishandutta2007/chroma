// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	dbmodel "github.com/chroma-core/chroma/go/pkg/sysdb/metastore/db/dbmodel"
	mock "github.com/stretchr/testify/mock"
)

// ICollectionDb is an autogenerated mock type for the ICollectionDb type
type ICollectionDb struct {
	mock.Mock
}

// DeleteAll provides a mock function with given fields:
func (_m *ICollectionDb) DeleteAll() error {
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

// DeleteCollectionByID provides a mock function with given fields: collectionID
func (_m *ICollectionDb) DeleteCollectionByID(collectionID string) (int, error) {
	ret := _m.Called(collectionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCollectionByID")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(collectionID)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(collectionID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionEntry provides a mock function with given fields: collectionID, databaseName
func (_m *ICollectionDb) GetCollectionEntry(collectionID *string, databaseName *string) (*dbmodel.Collection, error) {
	ret := _m.Called(collectionID, databaseName)

	if len(ret) == 0 {
		panic("no return value specified for GetCollectionEntry")
	}

	var r0 *dbmodel.Collection
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string) (*dbmodel.Collection, error)); ok {
		return rf(collectionID, databaseName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string) *dbmodel.Collection); ok {
		r0 = rf(collectionID, databaseName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dbmodel.Collection)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(collectionID, databaseName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollectionSize provides a mock function with given fields: collectionID
func (_m *ICollectionDb) GetCollectionSize(collectionID string) (uint64, error) {
	ret := _m.Called(collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetCollectionSize")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uint64, error)); ok {
		return rf(collectionID)
	}
	if rf, ok := ret.Get(0).(func(string) uint64); ok {
		r0 = rf(collectionID)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(collectionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollections provides a mock function with given fields: collectionID, collectionName, tenantID, databaseName, limit, offset
func (_m *ICollectionDb) GetCollections(collectionID *string, collectionName *string, tenantID string, databaseName string, limit *int32, offset *int32) ([]*dbmodel.CollectionAndMetadata, error) {
	ret := _m.Called(collectionID, collectionName, tenantID, databaseName, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetCollections")
	}

	var r0 []*dbmodel.CollectionAndMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, string, string, *int32, *int32) ([]*dbmodel.CollectionAndMetadata, error)); ok {
		return rf(collectionID, collectionName, tenantID, databaseName, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, string, string, *int32, *int32) []*dbmodel.CollectionAndMetadata); ok {
		r0 = rf(collectionID, collectionName, tenantID, databaseName, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.CollectionAndMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, string, string, *int32, *int32) error); ok {
		r1 = rf(collectionID, collectionName, tenantID, databaseName, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSoftDeletedCollections provides a mock function with given fields: collectionID, tenantID, databaseName, limit
func (_m *ICollectionDb) GetSoftDeletedCollections(collectionID *string, tenantID string, databaseName string, limit int32) ([]*dbmodel.CollectionAndMetadata, error) {
	ret := _m.Called(collectionID, tenantID, databaseName, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetSoftDeletedCollections")
	}

	var r0 []*dbmodel.CollectionAndMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, string, string, int32) ([]*dbmodel.CollectionAndMetadata, error)); ok {
		return rf(collectionID, tenantID, databaseName, limit)
	}
	if rf, ok := ret.Get(0).(func(*string, string, string, int32) []*dbmodel.CollectionAndMetadata); ok {
		r0 = rf(collectionID, tenantID, databaseName, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.CollectionAndMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, string, string, int32) error); ok {
		r1 = rf(collectionID, tenantID, databaseName, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: in
func (_m *ICollectionDb) Insert(in *dbmodel.Collection) error {
	ret := _m.Called(in)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.Collection) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCollectionsToGc provides a mock function with given fields:
func (_m *ICollectionDb) ListCollectionsToGc() ([]*dbmodel.CollectionToGc, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListCollectionsToGc")
	}

	var r0 []*dbmodel.CollectionToGc
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*dbmodel.CollectionToGc, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*dbmodel.CollectionToGc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.CollectionToGc)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: in
func (_m *ICollectionDb) Update(in *dbmodel.Collection) error {
	ret := _m.Called(in)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*dbmodel.Collection) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLogPositionVersionAndTotalRecords provides a mock function with given fields: collectionID, logPosition, currentCollectionVersion, totalRecordsPostCompaction
func (_m *ICollectionDb) UpdateLogPositionVersionAndTotalRecords(collectionID string, logPosition int64, currentCollectionVersion int32, totalRecordsPostCompaction uint64) (int32, error) {
	ret := _m.Called(collectionID, logPosition, currentCollectionVersion, totalRecordsPostCompaction)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLogPositionVersionAndTotalRecords")
	}

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, int32, uint64) (int32, error)); ok {
		return rf(collectionID, logPosition, currentCollectionVersion, totalRecordsPostCompaction)
	}
	if rf, ok := ret.Get(0).(func(string, int64, int32, uint64) int32); ok {
		r0 = rf(collectionID, logPosition, currentCollectionVersion, totalRecordsPostCompaction)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(string, int64, int32, uint64) error); ok {
		r1 = rf(collectionID, logPosition, currentCollectionVersion, totalRecordsPostCompaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewICollectionDb creates a new instance of ICollectionDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICollectionDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICollectionDb {
	mock := &ICollectionDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
