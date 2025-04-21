package handler

import (
	"GST/billlingSystem/model"
	"GST/billlingSystem/service"
	"encoding/json"
	"net/http"
)

func Login(deps service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := model.User{}

		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			return
		}

		token, err := deps.BillerService.Login(user)
		if err != nil {
			return
		}
		loginRes := model.LoginResponse{}
		if token != "" {
			loginRes = model.LoginResponse{
				Message: "login successful",
				Token:   token,
			}
		} else {
			loginRes.Message = "Invalid user"
		}

		json_response, err := json.Marshal(loginRes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

		}
		w.Header().Add("Content-Type", "application/json")

		w.Write(json_response)

	}
}
