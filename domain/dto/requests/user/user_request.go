package request

// Optional: define your structs here
type FetchUserProfileRequestDto struct {
	EmployeeID string `form:"employee_id" validate:"required"`
}
