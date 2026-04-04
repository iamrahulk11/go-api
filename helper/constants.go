package helper

const (
	// Login
	UsernameRequired  = "Username is required"
	PasswordRequired  = "Password is required"
	PasswordMinLength = "Password must be at least 6 characters"

	// User profile
	EmployeeIDRequired = "Employee ID is required"
)

// Custom validation messages map
var customMessages = map[string]string{
	"Username.required": "Username is required",
	"Password.required": "Password is required",
}
