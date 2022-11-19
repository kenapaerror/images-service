package main

import (
	"fmt"
	"github.com/kenapaerror/images-service/app"
	"github.com/kenapaerror/images-service/controller"
	"github.com/kenapaerror/images-service/exception"
	"github.com/kenapaerror/images-service/helper"
	"github.com/kenapaerror/images-service/repository"
	"github.com/kenapaerror/images-service/service"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	envErr := godotenv.Load(".env")
	helper.PanicIfError(envErr)

	webPort := os.Getenv("PORT")

	log.Printf("starting service on port %s\n", webPort)

	db := app.NewDB()
	validate := validator.New()

	repository := repository.NewImageRepositoryImpl()

	service := service.NewImageServiceImpl(repository, db, validate)

	controller := controller.NewImageControllerImpl(service)

	router := httprouter.New()

	router.POST("/api/image", controller.Create)
	router.DELETE("/api/image/:imageId", controller.Delete)
	router.GET("/api/image/:imageId", controller.FindById)

	router.PanicHandler = exception.ErrorHandler

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
