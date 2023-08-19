package main

import (
	"learn-7-restful-api/app"
	"learn-7-restful-api/controller"
	"learn-7-restful-api/helper"
	"learn-7-restful-api/middleware"
	"learn-7-restful-api/repository"
	"learn-7-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	db := app.NewDB("mysql", "golang:golang009@tcp(localhost:3306)/belajar_golang_restful_api")

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
