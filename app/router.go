package app

import (
	"embed"
	"html/template"
	"learn-7-restful-api/controller"
	"learn-7-restful-api/exception"
	"learn-7-restful-api/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed page/not-found.html
var notFoundTemplates embed.FS

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(notFoundTemplates, "page/not-found.html")

		helper.PanicIfError(err)

		t.ExecuteTemplate(w, "not-found.html", nil)
	})

	return router
}
