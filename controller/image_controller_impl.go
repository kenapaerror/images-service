package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kenapaerror/images-service/helper"
	"github.com/kenapaerror/images-service/model/web"
	"github.com/kenapaerror/images-service/service"
	"net/http"
	"os"
)

type ImageControllerImpl struct {
	ImageService service.ImageService
}

func NewImageControllerImpl(imageService service.ImageService) ImageController {
	return &ImageControllerImpl{
		ImageService: imageService,
	}
}

func (controller *ImageControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	request.ParseMultipartForm(10 * 1024 * 1024)

	imageCreateRequest := web.ImageCreateRequest{}

	imageCreateRequest.FormData = request.MultipartForm.File["image"]

	exampleResponse := controller.ImageService.Create(request.Context(), imageCreateRequest)
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
		Data:    exampleResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *ImageControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	imageId := params.ByName("imageId")

	controller.ImageService.Delete(request.Context(), imageId)
	response := web.WebResponse{
		Error:   false,
		Message: "OK",
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *ImageControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	imageId := params.ByName("imageId")

	exampleResponse := controller.ImageService.FindById(request.Context(), imageId)
	fileBytes, err := os.ReadFile("public/" + exampleResponse.Path)
	helper.PanicIfError(err)

	writer.Write(fileBytes)
}
