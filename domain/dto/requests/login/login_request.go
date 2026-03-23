package request

// Optional: define your structs here
type VerifyLoginRequestDto struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}
