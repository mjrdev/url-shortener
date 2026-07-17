package request

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required,min=8,max=100"`
}
