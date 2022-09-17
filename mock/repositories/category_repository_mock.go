package repositories

import (
	"belajar-golang-unit-test/mock/entities"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(Id int) *entities.Category {
	arguments := repository.Mock.Called(Id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entities.Category)
		return &category
	}
}
