package request

// Optional: define your structs here
type VerifyLoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
