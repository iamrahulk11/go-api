package services

import (
	"errors"
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

func (s *LoginServiceStruct) VerifyUserService(Login request.VerifyLoginRequestDto) (response.JWTResponse, error) {
	isValid, err := s.iLoginService.VerifyUserRepo(Login)
	if err != nil {
		return response.JWTResponse{}, err
	}

	if !isValid {
		return response.JWTResponse{}, errors.New("invalid credentials")
	}

	// User verified → create JWT
	token, err := s.jwtHelper.CreateToken(map[string]interface{}{
		"userId": Login.Username,
		"role":   "admin",
	})

	if err != nil {
		return response.JWTResponse{}, err
	}

	return response.JWTResponse{
		Token: token,
	}, nil
}
