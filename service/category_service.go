package service

import (
	"context"
	"learn-7-restful-api/model/web"
)

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
