package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func Read(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		data := []model.Book{
			{
				Name: "momo",
    		Year: 2022,
    		Author: "momotaros",
    		Summary: "lorem ipsum dolor sit amet",
    		Publisher: "lorem",
    		PageCount: 100,
    		ReadPage: 50,
    		Reading: false,
			},
			{
				Name: "sana",
    		Year: 2022,
    		Author: "momotaros",
    		Summary: "lorem ipsum dolor sit amet",
    		Publisher: "lorem",
    		PageCount: 100,
    		ReadPage: 50,
    		Reading: false,
			},
		}

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