package web

import "mime/multipart"

type ImageCreateRequest struct {
	FormData []*multipart.FileHeader
}
