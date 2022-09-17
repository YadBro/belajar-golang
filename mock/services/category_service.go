package services

import (
	"belajar-golang-unit-test/mock/entities"
	"belajar-golang-unit-test/mock/repositories"
	"errors"
)

type CategoryService struct {
	Repository repositories.CategoryRepository
}

func (service CategoryService) Get(Id int) (*entities.Category, error) {
	category := service.Repository.FindById(Id)

	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}
