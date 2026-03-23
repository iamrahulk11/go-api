package repository

import (
	request "user-mapping/domain/dto/requests/login"
	"user-mapping/domain/interfaces"
	sqlwrapper "user-mapping/infrastructure"
)

type LoginRepository struct {
	SQL *sqlwrapper.SQLWrapper
}

func NewLoginRepository(sqlWrapper *sqlwrapper.SQLWrapper) *LoginRepository {
	return &LoginRepository{
		SQL: sqlWrapper,
	}
}

var _ interfaces.ILoginService = (*LoginRepository)(nil)

func (r *LoginRepository) VerifyUserRepo(Login request.VerifyLoginRequestDto) (bool, error) {
	if Login.Username == "admin" && Login.Password == "123" {
		return true, nil
	}

	return false, nil
}
