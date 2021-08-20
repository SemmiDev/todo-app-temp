package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, res interface{}) {
	err := json.NewDecoder(req.Body).Decode(res)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, res interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(res)
	PanicIfError(err)
}
