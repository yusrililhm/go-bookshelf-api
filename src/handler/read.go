package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func Read(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		data := model.Books

		for _, v := range data {
			message := map[string]interface{}{
				"status": "success",
				"data": map[string]any{
					"book": map[string]any{
						"id": v.Id,
						"name": v.Name,
						"publisher": v.Publisher,
					},
				},
			}
	
			encode, err := json.Marshal(message)
	
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(encode)
		}
		return
	}
	http.Error(w, "Error", 500)
}