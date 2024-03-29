// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-vacancy-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-vacancy-api/internal/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddVacancies mocks base method.
func (m *MockRepo) AddVacancies(arg0 context.Context, arg1 []models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVacancies", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVacancies indicates an expected call of AddVacancies.
func (mr *MockRepoMockRecorder) AddVacancies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVacancies", reflect.TypeOf((*MockRepo)(nil).AddVacancies), arg0, arg1)
}

// CreateVacancy mocks base method.
func (m *MockRepo) CreateVacancy(arg0 context.Context, arg1 models.Vacancy) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVacancy", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVacancy indicates an expected call of CreateVacancy.
func (mr *MockRepoMockRecorder) CreateVacancy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVacancy", reflect.TypeOf((*MockRepo)(nil).CreateVacancy), arg0, arg1)
}

// DescribeVacancy mocks base method.
func (m *MockRepo) DescribeVacancy(arg0 context.Context, arg1 uint64) (models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeVacancy", arg0, arg1)
	ret0, _ := ret[0].(models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVacancy indicates an expected call of DescribeVacancy.
func (mr *MockRepoMockRecorder) DescribeVacancy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVacancy", reflect.TypeOf((*MockRepo)(nil).DescribeVacancy), arg0, arg1)
}

// ListVacancies mocks base method.
func (m *MockRepo) ListVacancies(arg0 context.Context, arg1, arg2 uint64) ([]models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVacancies", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVacancies indicates an expected call of ListVacancies.
func (mr *MockRepoMockRecorder) ListVacancies(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVacancies", reflect.TypeOf((*MockRepo)(nil).ListVacancies), arg0, arg1, arg2)
}

// RemoveVacancy mocks base method.
func (m *MockRepo) RemoveVacancy(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVacancy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveVacancy indicates an expected call of RemoveVacancy.
func (mr *MockRepoMockRecorder) RemoveVacancy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVacancy", reflect.TypeOf((*MockRepo)(nil).RemoveVacancy), arg0, arg1)
}

// UpdateVacancy mocks base method.
func (m *MockRepo) UpdateVacancy(arg0 context.Context, arg1 models.Vacancy) (models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVacancy", arg0, arg1)
	ret0, _ := ret[0].(models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVacancy indicates an expected call of UpdateVacancy.
func (mr *MockRepoMockRecorder) UpdateVacancy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVacancy", reflect.TypeOf((*MockRepo)(nil).UpdateVacancy), arg0, arg1)
}
