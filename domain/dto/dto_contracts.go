package dto

import (
	request "user-mapping/domain/dto/requests/login"
	profileRequest "user-mapping/domain/dto/requests/user"
)

type VerifyLoginRequest = request.VerifyLoginRequestDto
type FetchUserProfileRequest = profileRequest.FetchUserProfileRequestDto
