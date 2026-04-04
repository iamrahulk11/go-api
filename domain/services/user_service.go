package services

import (
	request "user-mapping/domain/dto/requests/user"
	response "user-mapping/domain/dto/response/user"
	"user-mapping/domain/interfaces"
	"user-mapping/helper"
)

type UserServiceStruct struct {
	iUserService interfaces.IUserService
	jwtHelper    *helper.JWT
}

func NewUserService(jwtHelper *helper.JWT, iUserService interfaces.IUserService) *UserServiceStruct {
	return &UserServiceStruct{
		jwtHelper:    jwtHelper,
		iUserService: iUserService,
	}
}

func (s *UserServiceStruct) UserService() (*response.AllUserResponse, error) {
	allUserResponse, err := s.iUserService.FetchAllUser()
	if err != nil {
		return allUserResponse, err
	}

	return allUserResponse, nil
}
func (s *UserServiceStruct) FetchUserProfileDetails(request request.FetchUserProfileRequestDto) (*response.UserBasicDetailsResponse, error) {
	allUserResponse, err := s.iUserService.FetchUserProfile(request.EmployeeID)
	if err != nil {
		return allUserResponse, err
	}

	return allUserResponse, nil
}
