package interfaces

import response "user-mapping/domain/dto/response/user"

type IUserService interface {
	FetchAllUser() (*response.AllUserResponse, error)
	FetchUserProfile(User_id string) (*response.UserBasicDetailsResponse, error)
}
