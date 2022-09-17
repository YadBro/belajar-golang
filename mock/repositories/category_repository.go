package repositories

import "belajar-golang-unit-test/mock/entities"

type CategoryRepository interface {
	FindById(Id int) *entities.Category
}
