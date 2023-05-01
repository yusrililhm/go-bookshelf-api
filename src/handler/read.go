package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func Read(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		data := model.Book{}

		encode, err := json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(encode)
		return
	}
	http.Error(w, "Error", http.StatusInternalServerError)
}