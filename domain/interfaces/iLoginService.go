package interfaces

import request "user-mapping/domain/dto/requests/login"

type ILoginService interface {
	VerifyUserRepo(request.VerifyLoginRequestDto) (bool, error)
}
