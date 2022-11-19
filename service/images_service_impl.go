package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/kenapaerror/images-service/exception"
	"github.com/kenapaerror/images-service/helper"
	"github.com/kenapaerror/images-service/model/entity"
	"github.com/kenapaerror/images-service/model/web"
	"github.com/kenapaerror/images-service/repository"
	"github.com/kenapaerror/images-service/utils"
	"io"
	"os"
	"strings"
)

type ImageServiceImpl struct {
	ImageRepository repository.ImageRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewImageServiceImpl(imageRepository repository.ImageRepository, DB *sql.DB, validate *validator.Validate) ImageService {
	return &ImageServiceImpl{
		ImageRepository: imageRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *ImageServiceImpl) Create(ctx context.Context, request web.ImageCreateRequest) []web.ImageResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var imageResponses []web.ImageResponse

	for _, image := range request.FormData {
		file, _ := image.Open()

		tempFile, err := os.CreateTemp("public", "image-*.jpg")
		helper.PanicIfError(err)
		defer tempFile.Close()

		fileBytes, err := io.ReadAll(file)
		helper.PanicIfError(err)

		tempFile.Write(fileBytes)

		fileName := tempFile.Name()
		newFileName := strings.Split(fileName, "\\")

		image := entity.Images{
			Id:        utils.GenerateId(),
			Path:      newFileName[1],
			CreatedAt: utils.CurrentMillis(),
			UpdatedAt: utils.CurrentMillis(),
		}

		image = service.ImageRepository.Create(ctx, tx, image)
		imageResponses = append(imageResponses, utils.ToImageResponse(image))
	}

	return imageResponses
}

func (service *ImageServiceImpl) Delete(ctx context.Context, imageId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ImageRepository.Delete(ctx, tx, image)

	os.Remove("public/" + image.Path)
}

func (service *ImageServiceImpl) FindById(ctx context.Context, imageId string) web.ImageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	image, err := service.ImageRepository.FindById(ctx, tx, imageId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return utils.ToImageResponse(image)
}
