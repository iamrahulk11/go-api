package interfaces

type IUserService interface {
	FetchAllUser() ([]map[string]interface{}, error)
	FetchUserProfile(User_id string) ([]map[string]interface{}, error)
}
