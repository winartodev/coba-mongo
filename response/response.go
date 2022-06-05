package response

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Err    string      `json:"error"`
}

func HttpResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	setResponse := response{
		Status: http.StatusText(http.StatusOK),
		Data:   data,
		Err:    "",
	}

	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func HttpResponseFailed(w http.ResponseWriter, r *http.Request, code int, err error) {
	setResponse := response{
		Status: http.StatusText(code),
		Data:   nil,
		Err:    err.Error(),
	}

	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
