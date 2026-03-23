package api

import (
	"encoding/json"
	"net/http"
	"user-mapping/domain/dto"
	request "user-mapping/domain/dto/requests/login"
	"user-mapping/domain/services"
)

type LoginHandlerStruct struct {
	services *services.LoginServiceStruct
}

func LoginHandler(services *services.LoginServiceStruct) *LoginHandlerStruct {
	return &LoginHandlerStruct{services: services}
}

// VerifyUser now accepts the parsed request DTO
func (h *LoginHandlerStruct) VerifyUser(w http.ResponseWriter, loginReq request.VerifyLoginRequestDto) {
	result, err := h.services.VerifyUserService(loginReq)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.BaseResponseDto[any]{
			Result: dto.ResultResponseDto{
				Flag:        0,
				FlagMessage: err.Error(),
			},
		})
		return
	}

	// Success
	json.NewEncoder(w).Encode(dto.BaseResponseDto[any]{
		Result: dto.ResultResponseDto{
			Flag:        1,
			FlagMessage: "Success",
		},
		Data: result,
	})
}
