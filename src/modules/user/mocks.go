package user

import (
	"github.com/adiet95/go-order-api/src/database"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindByEmail(email string) (*database.User, error) {
	args := m.mock.Called(email)
	return args.Get(0).(*database.User), nil
}

func (m *RepoMock) RegisterEmail(data *database.User) (*database.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*database.User), nil
}
