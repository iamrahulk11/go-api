package repository

import (
	request "user-mapping/domain/dto/requests/login"
	"user-mapping/infrastructure"
)

type LoginRepository struct {
	db *infrastructure.SQLWrapper
}

func NewLoginRepository(db *infrastructure.SQLWrapper) *LoginRepository {
	return &LoginRepository{db: db}
}

func (s LoginRepository) VerifyUserRepo(Login request.VerifyLoginRequestDto) (bool, error) {
	if Login.Username == "admin" && Login.Password == "123" {
		return true, nil
	}

	return false, nil
}
