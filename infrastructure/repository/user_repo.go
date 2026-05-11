package repository

import (
	"fmt"
	response "user-mapping/domain/dto/response/user"
	"user-mapping/infrastructure"
)

type UserRepository struct {
	sqlWrapper *infrastructure.SQLWrapper
}

func NewUserRepository(sql_wrapper *infrastructure.SQLWrapper) *UserRepository {
	return &UserRepository{sqlWrapper: sql_wrapper}
}

func (u *UserRepository) FetchAllUser() (*response.AllUserResponse, error) {
	return &response.AllUserResponse{
		Username: "rahul",
	}, nil
}

func (u *UserRepository) FetchUserProfile(User_id string) ([]map[string]interface{}, error) {
	const query = `
		select employee_id username, employee_name name, employee_email email, employee_gender gender, is_active employement_status, employee_department department, employee_designation designation 
		from adm_user_details where employee_id = :user_id
	`

	params := map[string]interface{}{
		"user_id": User_id,
	}

	result, err := u.sqlWrapper.ExecuteQuery("localDB", query, params)
	if err != nil {
		return nil, fmt.Errorf("fetch user profile query failed: %w", err)
	}

	// defer rows.Close()
	// // handle no data
	// if !rows.Next() {
	// 	return nil, sql.ErrNoRows
	// }
	// // bind by column name
	// if err := rows.StructScan(&resp); err != nil {
	// 	return nil, fmt.Errorf("struct scan failed: %w", err)
	// }

	return result, nil
}
