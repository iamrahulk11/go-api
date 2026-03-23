package request

// Optional: define your structs here
type FetchUserProfileRequestDto struct {
	Employee_id string `schema:"employee_id" validate:"required"`
}
