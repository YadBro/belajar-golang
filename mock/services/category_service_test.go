package services

import (
	"belajar-golang-unit-test/mock/entities"
	"belajar-golang-unit-test/mock/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repositories.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", 1).Return(nil)

	category, error := categoryService.Get(1)
	assert.Nil(t, category)
	assert.NotNil(t, error)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entities.Category{
		Id:   2,
		Name: "yadi",
	}
	categoryRepository.Mock.On("FindById", 2).Return(category)

	result, error := categoryService.Get(2)
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
