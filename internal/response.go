package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/io-m/bankapp/internal/errors"
)

// ResponseJSON writes JSON object back to the connection. It requires struct where response can be written to
// as well as status code of the response
func ResponseJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	response, err := json.Marshal(data)
	if err != nil {
		log.Println("Could not marshal into response struct")
		return
	}
	w.Write(response)
}
// ResponseError writes JSON object for error response. It accepts customer Error struct and sends status code provided with
// that struct together with the message
func ResponseError(w http.ResponseWriter, e *errors.Error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(e.StatusCode)
	errorResponse, err := json.Marshal(e)
	if err != nil {
		log.Println("Could not marshal into response error struct")
		return
	}
	w.Write(errorResponse)
}