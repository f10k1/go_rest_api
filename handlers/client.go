package handlers

import (
	"encoding/json"
	"net/http"
	"rest_api/database"
)

func HandleClientProfile(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		GetClientProfile(writer, request)
	case http.MethodPatch:
		UpdateClientProfile(writer, request)
	default:
		http.Error(writer, "Method is not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(writer http.ResponseWriter, request *http.Request) {
	clientProfile := request.Context().Value("clientProfile").(database.ClientProfile)

	writer.Header().Set("Content-Type", "application/json")

	response := database.ClientProfile{
		Email: clientProfile.Email,
		Name:  clientProfile.Name,
		Id:    clientProfile.Id,
	}

	json.NewEncoder(writer).Encode(response)
}

func UpdateClientProfile(writer http.ResponseWriter, request *http.Request) {
	clientProfile := request.Context().Value("clientProfile").(database.ClientProfile)

	var payloadData database.ClientProfile
	if err := json.NewDecoder(request.Body).Decode(&payloadData); err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	clientProfile.Email = payloadData.Email
	clientProfile.Name = payloadData.Name
	database.Database[clientProfile.Id] = clientProfile

	writer.WriteHeader(http.StatusOK)
}
