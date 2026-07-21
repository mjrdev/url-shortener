package request

type CreateUrlRequest struct {
	Url string `json:"url" validate:"required"`
}
