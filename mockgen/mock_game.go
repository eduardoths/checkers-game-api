// Code generated by MockGen. DO NOT EDIT.
// Source: src/interfaces/game.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	structs "github.com/eduardoths/checkers-game/src/structs"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockGameService is a mock of GameService interface.
type MockGameService struct {
	ctrl     *gomock.Controller
	recorder *MockGameServiceMockRecorder
}

// MockGameServiceMockRecorder is the mock recorder for MockGameService.
type MockGameServiceMockRecorder struct {
	mock *MockGameService
}

// NewMockGameService creates a new mock instance.
func NewMockGameService(ctrl *gomock.Controller) *MockGameService {
	mock := &MockGameService{ctrl: ctrl}
	mock.recorder = &MockGameServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameService) EXPECT() *MockGameServiceMockRecorder {
	return m.recorder
}

// Move mocks base method.
func (m *MockGameService) Move(gameID uuid.UUID, from int, movements []int) (*structs.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Move", gameID, from, movements)
	ret0, _ := ret[0].(*structs.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Move indicates an expected call of Move.
func (mr *MockGameServiceMockRecorder) Move(gameID, from, movements interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockGameService)(nil).Move), gameID, from, movements)
}

// NewGame mocks base method.
func (m *MockGameService) NewGame(playerOne, playerTwo *structs.Player) (*structs.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGame", playerOne, playerTwo)
	ret0, _ := ret[0].(*structs.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGame indicates an expected call of NewGame.
func (mr *MockGameServiceMockRecorder) NewGame(playerOne, playerTwo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGame", reflect.TypeOf((*MockGameService)(nil).NewGame), playerOne, playerTwo)
}

// MockGameRepository is a mock of GameRepository interface.
type MockGameRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameRepositoryMockRecorder
}

// MockGameRepositoryMockRecorder is the mock recorder for MockGameRepository.
type MockGameRepositoryMockRecorder struct {
	mock *MockGameRepository
}

// NewMockGameRepository creates a new mock instance.
func NewMockGameRepository(ctrl *gomock.Controller) *MockGameRepository {
	mock := &MockGameRepository{ctrl: ctrl}
	mock.recorder = &MockGameRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameRepository) EXPECT() *MockGameRepositoryMockRecorder {
	return m.recorder
}

// FindGame mocks base method.
func (m *MockGameRepository) FindGame(id uuid.UUID) (*structs.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindGame", id)
	ret0, _ := ret[0].(*structs.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindGame indicates an expected call of FindGame.
func (mr *MockGameRepositoryMockRecorder) FindGame(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindGame", reflect.TypeOf((*MockGameRepository)(nil).FindGame), id)
}