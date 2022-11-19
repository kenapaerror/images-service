package utils

import (
	"github.com/kenapaerror/images-service/model/entity"
	"github.com/kenapaerror/images-service/model/web"
)

func ToExampleResponse(example entity.Example) web.ExampleResponse {
	return web.ExampleResponse{
		Id:        example.Id,
		Name:      example.Name,
		Email:     example.Email,
		CreatedAt: example.CreatedAt,
		UpdatedAt: example.UpdatedAt,
	}
}

func ToImageResponse(image entity.Images) web.ImageResponse {
	return web.ImageResponse{
		Id:        image.Id,
		Path:      image.Path,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}
