package api

import (
	"encoding/json"
	"net/http"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
