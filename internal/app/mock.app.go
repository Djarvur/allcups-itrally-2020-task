// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package app is a generated GoMock package.
package app

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	game "github.com/Djarvur/allcups-itrally-2020-task/internal/app/game"
)

// MockAppl is a mock of Appl interface
type MockAppl struct {
	ctrl     *gomock.Controller
	recorder *MockApplMockRecorder
}

// MockApplMockRecorder is the mock recorder for MockAppl
type MockApplMockRecorder struct {
	mock *MockAppl
}

// NewMockAppl creates a new mock instance
func NewMockAppl(ctrl *gomock.Controller) *MockAppl {
	mock := &MockAppl{ctrl: ctrl}
	mock.recorder = &MockApplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppl) EXPECT() *MockApplMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method
func (m *MockAppl) HealthCheck(arg0 Ctx) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockApplMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockAppl)(nil).HealthCheck), arg0)
}

// Start mocks base method
func (m *MockAppl) Start(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockApplMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockAppl)(nil).Start), arg0)
}

// Balance mocks base method
func (m *MockAppl) Balance(arg0 Ctx) (int, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Balance indicates an expected call of Balance
func (mr *MockApplMockRecorder) Balance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockAppl)(nil).Balance), arg0)
}

// Licenses mocks base method
func (m *MockAppl) Licenses(arg0 Ctx) ([]game.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Licenses", arg0)
	ret0, _ := ret[0].([]game.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Licenses indicates an expected call of Licenses
func (mr *MockApplMockRecorder) Licenses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Licenses", reflect.TypeOf((*MockAppl)(nil).Licenses), arg0)
}

// IssueLicense mocks base method
func (m *MockAppl) IssueLicense(arg0 Ctx, wallet []int) (game.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueLicense", arg0, wallet)
	ret0, _ := ret[0].(game.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueLicense indicates an expected call of IssueLicense
func (mr *MockApplMockRecorder) IssueLicense(arg0, wallet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueLicense", reflect.TypeOf((*MockAppl)(nil).IssueLicense), arg0, wallet)
}

// ExploreArea mocks base method
func (m *MockAppl) ExploreArea(arg0 Ctx, area game.Area) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExploreArea", arg0, area)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExploreArea indicates an expected call of ExploreArea
func (mr *MockApplMockRecorder) ExploreArea(arg0, area interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExploreArea", reflect.TypeOf((*MockAppl)(nil).ExploreArea), arg0, area)
}

// Dig mocks base method
func (m *MockAppl) Dig(arg0 Ctx, licenseID int, pos game.Coord) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dig", arg0, licenseID, pos)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dig indicates an expected call of Dig
func (mr *MockApplMockRecorder) Dig(arg0, licenseID, pos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dig", reflect.TypeOf((*MockAppl)(nil).Dig), arg0, licenseID, pos)
}

// Cash mocks base method
func (m *MockAppl) Cash(arg0 Ctx, treasure string) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cash", arg0, treasure)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cash indicates an expected call of Cash
func (mr *MockApplMockRecorder) Cash(arg0, treasure interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cash", reflect.TypeOf((*MockAppl)(nil).Cash), arg0, treasure)
}

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// LoadStartTime mocks base method
func (m *MockRepo) LoadStartTime() (*time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadStartTime")
	ret0, _ := ret[0].(*time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadStartTime indicates an expected call of LoadStartTime
func (mr *MockRepoMockRecorder) LoadStartTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadStartTime", reflect.TypeOf((*MockRepo)(nil).LoadStartTime))
}

// SaveStartTime mocks base method
func (m *MockRepo) SaveStartTime(t time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveStartTime", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveStartTime indicates an expected call of SaveStartTime
func (mr *MockRepoMockRecorder) SaveStartTime(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveStartTime", reflect.TypeOf((*MockRepo)(nil).SaveStartTime), t)
}
