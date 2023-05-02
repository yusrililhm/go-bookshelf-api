package handler

import (
	"encoding/json"
	"net/http"
)

func ErrorJson(w http.ResponseWriter, err interface{}, code int)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}