package res

import (
	"encoding/json"
	"net/http"
)

func ErrorJson(w http.ResponseWriter, message map[string]any, code int)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(message)
}