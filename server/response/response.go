package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResp : set status and encode data to json
// Example of usage: responses.JSON(w, http.StatusCreated, userCreated)
func JSONResp(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Print("JSONResp: error during encoding: ", err.Error())
	}
}

// ErrorResp : return status code and error specified
// Example of usage: responses.ErrorResp(w, http.StatusUnauthorized, errors.New("Unauthorized"))
func ErrorResp(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSONResp(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSONResp(w, http.StatusBadRequest, nil)
}
