package api

import (
	"encoding/json"
	"net/http"
	request "user-mapping/domain/dto/requests/user"
	"user-mapping/domain/services"
)

type UserHandlerStruct struct {
	services *services.UserServiceStruct
}

func UserHandler(services *services.UserServiceStruct) *UserHandlerStruct {
	return &UserHandlerStruct{services: services}
}

// VerifyUser now accepts the parsed request DTO
func (h *UserHandlerStruct) User(w http.ResponseWriter, r *http.Request) {
	result, err := h.services.UserService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *UserHandlerStruct) FetchUserProfile(w http.ResponseWriter, r request.FetchUserProfileRequestDto) {
	result, err := h.services.FetchUserProfileDetails(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
