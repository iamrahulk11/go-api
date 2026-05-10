package helper

const (
	DATA_FOUND          = "Data found."
	DATA_NO_FOUND       = "Data not found."
	INVALID_CREDENTIALS = "Invalid Credentials."

	// Login
	UsernameRequired  = "Username is required"
	PasswordRequired  = "Password is required"
	PasswordMinLength = "Password must be at least 6 characters"

	// User profile
	EmployeeIDRequired = "Employee ID is required"
)

// Custom validation messages map
var CustomMessages = map[string]string{
	"Username.required": "Username is required",
	"Password.required": "Password is required",
}
