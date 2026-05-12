package services

import (
	"user-mapping/domain/dto"
	request "user-mapping/domain/dto/requests/user"
	"user-mapping/domain/interfaces"
	"user-mapping/helper"
)

type UserServiceStruct struct {
	iUserService interfaces.IUserService
}

func NewUserService(iUserService interfaces.IUserService) *UserServiceStruct {
	return &UserServiceStruct{
		iUserService: iUserService,
	}
}

func (s *UserServiceStruct) UserService() dto.BaseResponseDto[[]map[string]interface{}] {
	allUserResponse, err := s.iUserService.FetchAllUser()

	if err != nil {
		return dto.BaseResponseDto[[]map[string]interface{}]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: helper.DATA_NO_FOUND,
			},
			Data: nil,
		}
	}

	// Return success response
	return dto.BaseResponseDto[[]map[string]interface{}]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: helper.DATA_FOUND,
		},
		Data: allUserResponse,
	}
}
func (s *UserServiceStruct) FetchUserProfileDetails(request request.FetchUserProfileRequestDto) dto.BaseResponseDto[[]map[string]interface{}] {
	userProfile, err := s.iUserService.FetchUserProfile(request.EmployeeID)
	if err != nil {
		// Return failure response
		return dto.BaseResponseDto[[]map[string]interface{}]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: helper.DATA_NO_FOUND,
			},
			Data: nil,
		}
	}

	// Return success response
	return dto.BaseResponseDto[[]map[string]interface{}]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: helper.DATA_FOUND,
		},
		Data: userProfile,
	}
}
