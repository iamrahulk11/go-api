package services

import (
	"user-mapping/domain/dto"
	request "user-mapping/domain/dto/requests/login"
	response "user-mapping/domain/dto/response/login"
	"user-mapping/domain/interfaces"
	"user-mapping/helper"
	"user-mapping/internal/config"
)

type LoginServiceStruct struct {
	iLoginService interfaces.ILoginService
	jwtConfig     *config.JWTConfig
}

func NewLoginService(iLoginService interfaces.ILoginService, jwtConfiguration *config.JWTConfig) *LoginServiceStruct {
	return &LoginServiceStruct{
		iLoginService: iLoginService,
		jwtConfig:     jwtConfiguration,
	}
}

func (s *LoginServiceStruct) VerifyUserService(Login request.VerifyLoginRequestDto) dto.BaseResponseDto[response.JWTResponse] {
	isValid, err := s.iLoginService.VerifyUserRepo(Login)
	if err != nil {
		// Return failure response
		return dto.BaseResponseDto[response.JWTResponse]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: helper.DATA_NO_FOUND,
			},
		}
	}

	if !isValid {
		return dto.BaseResponseDto[response.JWTResponse]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: helper.INVALID_CREDENTIALS,
			},
		}
	}

	// User verified → create JWT
	token, err := helper.CreateToken(map[string]interface{}{
		"userId": Login.Username,
		"role":   "admin",
	},
		s.jwtConfig.Audience,
		s.jwtConfig.Issuer, s.jwtConfig.Secret, s.jwtConfig.ExpiresInMinute,
	)

	resp := response.JWTResponse{
		Token: token,
	}

	// Return success response
	return dto.BaseResponseDto[response.JWTResponse]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: helper.DATA_FOUND,
		},
		Data: resp,
	}
}
