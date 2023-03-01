package managers

import (
	"encoding/json"
	"errors"
	"net/http"

	"handler/function/types"
)

// print to http response not found
func ThrowNotFoundError(w http.ResponseWriter) {
	err := errors.New("path not handled")
	retunResponseError(w, err, http.StatusNotFound)
}

// print to http response internal server error
func ThrowInternalServerError(w http.ResponseWriter, err error) {
	retunResponseError(w, err, http.StatusInternalServerError)
}

// print to http response bad request error
func ThrowBadRequest(w http.ResponseWriter, err error) {
	retunResponseError(w, err, http.StatusBadRequest)
}

// print to http response unauthorized error
func ThrowUnauthorize(w http.ResponseWriter, err error) {
	retunResponseError(w, err, http.StatusUnauthorized)
}

// print error to response
func retunResponseError(w http.ResponseWriter, err error, statusCode int) {
	// create response
	ApplicationErrorResponse := types.ApplicationErrorResponse{}
	ApplicationErrorResponse.Status = http.StatusUnauthorized
	ApplicationErrorResponse.Message = err.Error()
	// convert to json
	json, _ := json.Marshal(err)
	// return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
