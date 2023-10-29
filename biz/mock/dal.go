// Code generated by MockGen. DO NOT EDIT.
// Source: ./biz/dal/dal.go
//
// Generated by this command:
//
//	mockgen -source=./biz/dal/dal.go -destination=./biz/mock/dal.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	redis_rate "github.com/go-redis/redis_rate/v10"
	dal "github.com/stonebirdjx/topx/biz/dal"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	gomock "go.uber.org/mock/gomock"
)

// MockDaler is a mock of Daler interface.
type MockDaler struct {
	ctrl     *gomock.Controller
	recorder *MockDalerMockRecorder
}

// MockDalerMockRecorder is the mock recorder for MockDaler.
type MockDalerMockRecorder struct {
	mock *MockDaler
}

// NewMockDaler creates a new mock instance.
func NewMockDaler(ctrl *gomock.Controller) *MockDaler {
	mock := &MockDaler{ctrl: ctrl}
	mock.recorder = &MockDalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaler) EXPECT() *MockDalerMockRecorder {
	return m.recorder
}

// Allow mocks base method.
func (m *MockDaler) Allow(ctx context.Context, key string, limit redis_rate.Limit) (*redis_rate.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Allow", ctx, key, limit)
	ret0, _ := ret[0].(*redis_rate.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Allow indicates an expected call of Allow.
func (mr *MockDalerMockRecorder) Allow(ctx, key, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allow", reflect.TypeOf((*MockDaler)(nil).Allow), ctx, key, limit)
}

// CreateAction mocks base method.
func (m *MockDaler) CreateAction(ctx context.Context, action *dal.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAction", ctx, action)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAction indicates an expected call of CreateAction.
func (mr *MockDalerMockRecorder) CreateAction(ctx, action any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAction", reflect.TypeOf((*MockDaler)(nil).CreateAction), ctx, action)
}

// DeleteActionByID mocks base method.
func (m *MockDaler) DeleteActionByID(ctx context.Context, id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActionByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActionByID indicates an expected call of DeleteActionByID.
func (mr *MockDalerMockRecorder) DeleteActionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActionByID", reflect.TypeOf((*MockDaler)(nil).DeleteActionByID), ctx, id)
}

// DeleteActions mocks base method.
func (m *MockDaler) DeleteActions(ctx context.Context, ids []primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActions", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActions indicates an expected call of DeleteActions.
func (mr *MockDalerMockRecorder) DeleteActions(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActions", reflect.TypeOf((*MockDaler)(nil).DeleteActions), ctx, ids)
}

// FindActionByID mocks base method.
func (m *MockDaler) FindActionByID(ctx context.Context, id primitive.ObjectID) (*dal.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActionByID", ctx, id)
	ret0, _ := ret[0].(*dal.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActionByID indicates an expected call of FindActionByID.
func (mr *MockDalerMockRecorder) FindActionByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActionByID", reflect.TypeOf((*MockDaler)(nil).FindActionByID), ctx, id)
}

// FindActionByOpt mocks base method.
func (m *MockDaler) FindActionByOpt(ctx context.Context, opt dal.FindActionOption) (*dal.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActionByOpt", ctx, opt)
	ret0, _ := ret[0].(*dal.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActionByOpt indicates an expected call of FindActionByOpt.
func (mr *MockDalerMockRecorder) FindActionByOpt(ctx, opt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActionByOpt", reflect.TypeOf((*MockDaler)(nil).FindActionByOpt), ctx, opt)
}

// ListActions mocks base method.
func (m *MockDaler) ListActions(ctx context.Context, opt dal.ListActionOption) (*[]dal.Action, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActions", ctx, opt)
	ret0, _ := ret[0].(*[]dal.Action)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListActions indicates an expected call of ListActions.
func (mr *MockDalerMockRecorder) ListActions(ctx, opt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActions", reflect.TypeOf((*MockDaler)(nil).ListActions), ctx, opt)
}

// UpdateAction mocks base method.
func (m *MockDaler) UpdateAction(ctx context.Context, action *dal.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAction", ctx, action)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAction indicates an expected call of UpdateAction.
func (mr *MockDalerMockRecorder) UpdateAction(ctx, action any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAction", reflect.TypeOf((*MockDaler)(nil).UpdateAction), ctx, action)
}
