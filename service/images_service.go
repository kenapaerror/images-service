package service

import (
	"context"
	"github.com/kenapaerror/images-service/model/web"
)

type ImageService interface {
	Create(ctx context.Context, request web.ImageCreateRequest) []web.ImageResponse
	Delete(ctx context.Context, imageId string)
	FindById(ctx context.Context, imageId string) web.ImageResponse
}
