package web

type ImageResponse struct {
	Id        string `json:"id"`
	Path      string `json:"path"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
