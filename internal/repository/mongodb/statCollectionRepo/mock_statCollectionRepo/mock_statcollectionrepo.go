// Code generated by MockGen. DO NOT EDIT.
// Source: stat_collection_repo.go

// Package mock_statcollectionrepo is a generated GoMock package.
package mock_statcollectionrepo

import (
	context "context"
	reflect "reflect"

	models "github.com/ZecretBone/ips-rssi-service/apps/rssi/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// InsertOne mocks base method.
func (m *MockRepository) InsertOne(ctx context.Context, document models.RSSIStatModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", ctx, document)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockRepositoryMockRecorder) InsertOne(ctx, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockRepository)(nil).InsertOne), ctx, document)
}
