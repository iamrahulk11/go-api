package request

// Optional: define your structs here
type VerifyLoginRequestDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
