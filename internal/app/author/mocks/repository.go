// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	author "gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
)

// MockAuthorRepo is a mock of AuthorRepo interface.
type MockAuthorRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorRepoMockRecorder
}

// MockAuthorRepoMockRecorder is the mock recorder for MockAuthorRepo.
type MockAuthorRepoMockRecorder struct {
	mock *MockAuthorRepo
}

// NewMockAuthorRepo creates a new mock instance.
func NewMockAuthorRepo(ctrl *gomock.Controller) *MockAuthorRepo {
	mock := &MockAuthorRepo{ctrl: ctrl}
	mock.recorder = &MockAuthorRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorRepo) EXPECT() *MockAuthorRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthorRepo) Create(ctx context.Context, authorModel *author.AuthorRow) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, authorModel)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthorRepoMockRecorder) Create(ctx, authorModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorRepo)(nil).Create), ctx, authorModel)
}

// Delete mocks base method.
func (m *MockAuthorRepo) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorRepoMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorRepo)(nil).Delete), ctx, id)
}

// GetById mocks base method.
func (m *MockAuthorRepo) GetById(ctx context.Context, id int64) (*author.AuthorRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(*author.AuthorRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAuthorRepoMockRecorder) GetById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAuthorRepo)(nil).GetById), ctx, id)
}

// Update mocks base method.
func (m *MockAuthorRepo) Update(ctx context.Context, id int64, authorModel *author.AuthorRow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, authorModel)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAuthorRepoMockRecorder) Update(ctx, id, authorModel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthorRepo)(nil).Update), ctx, id, authorModel)
}
