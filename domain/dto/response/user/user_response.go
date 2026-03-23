package response

type AllUserResponse struct {
	Username string `json:"username"`
}

type UserBasicDetailsResponse struct {
	Username         string  `db:"username" json:"username"`
	Name             string  `db:"name" json:"name"`
	Email            string  `db:"email" json:"email"`
	Gender           *string `db:"gender" json:"gender,omitempty"`
	EmploymentStatus *string `db:"employement_status" json:"employement_status,omitempty"`
	Department       *string `db:"department" json:"department,omitempty"`
	Designation      *string `db:"designation" json:"designation,omitempty"`
}

type UserProfileResponse struct {
	UserBasicDetailsResponse
	NotificationAllowed bool `db:"notification_allowed" json:"notification_allowed"`
}
