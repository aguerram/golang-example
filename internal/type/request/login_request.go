package request

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=255,min=2"`
	Password string `json:"password" validate:"required,max=255,min=2"`
}
