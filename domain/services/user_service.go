package services

import (
	"user-mapping/domain/dto"
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

func (s *UserServiceStruct) UserService() dto.BaseResponseDto[*response.AllUserResponse] {
	allUserResponse, err := s.iUserService.FetchAllUser()
	if err != nil {
		// Return failure response
		return dto.BaseResponseDto[*response.AllUserResponse]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: err.Error(),
			},
			Data: nil,
		}
	}

	// Return success response
	return dto.BaseResponseDto[*response.AllUserResponse]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: "Success",
		},
		Data: allUserResponse,
	}
}
func (s *UserServiceStruct) FetchUserProfileDetails(request request.FetchUserProfileRequestDto) dto.BaseResponseDto[*response.UserBasicDetailsResponse] {
	userProfile, err := s.iUserService.FetchUserProfile(request.EmployeeID)
	if err != nil {
		// Return failure response
		return dto.BaseResponseDto[*response.UserBasicDetailsResponse]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: err.Error(),
			},
			Data: nil,
		}
	}

	// Return success response
	return dto.BaseResponseDto[*response.UserBasicDetailsResponse]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: "Success",
		},
		Data: userProfile,
	}
}
