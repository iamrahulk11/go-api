package services

import (
	"user-mapping/domain/dto"
	request "user-mapping/domain/dto/requests/login"
	response "user-mapping/domain/dto/response/login"
	"user-mapping/domain/interfaces"
	"user-mapping/helper"
)

type LoginServiceStruct struct {
	iLoginService interfaces.ILoginService
	jwtHelper     *helper.JWT
}

func NewLoginService(jwtHelper *helper.JWT, iLoginService interfaces.ILoginService) *LoginServiceStruct {
	return &LoginServiceStruct{
		jwtHelper:     jwtHelper,
		iLoginService: iLoginService,
	}
}

func (s *LoginServiceStruct) VerifyUserService(Login request.VerifyLoginRequestDto) dto.BaseResponseDto[*response.JWTResponse] {
	isValid, err := s.iLoginService.VerifyUserRepo(Login)
	if err != nil {
		// Return failure response
		return dto.BaseResponseDto[*response.JWTResponse]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: helper.DATA_NO_FOUND,
			},
			Data: nil,
		}
	}

	if !isValid {
		return dto.BaseResponseDto[*response.JWTResponse]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: helper.INVALID_CREDENTIALS,
			},
			Data: nil,
		}
	}

	// User verified → create JWT
	token, err := s.jwtHelper.CreateToken(map[string]interface{}{
		"userId": Login.Username,
		"role":   "admin",
	})

	resp := response.JWTResponse{
		Token: token,
	}

	// Return success response
	return dto.BaseResponseDto[*response.JWTResponse]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: helper.DATA_FOUND,
		},
		Data: &resp,
	}
}
