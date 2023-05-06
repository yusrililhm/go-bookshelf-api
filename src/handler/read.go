package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/model"
	"github.com/yusrililhm/go-bookshelf-api/src/res"
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

func GetBookById(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		data := model.Books

		id := r.FormValue("id")

		for _, v := range data {
			if v.Id == id {
				message := map[string]any{
					"status": "success",
					"data": map[string]any{
						"book": map[string]any{
							"id": v.Id,
              "name": v.Name,
              "year": v.Year,
              "author": v.Author,
              "summary": v.Summary,
              "publisher": v.Publisher,
              "pageCount": v.PageCount,
              "readPage": v.ReadPage,
              "finished": v.Finished,
              "reading": v.Reading,
              "insertedAt": v.InsertedAt,
              "updatedAt": v.UpdateAt,
						},
					},
				}

				encode, err := json.Marshal(message)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(encode)
				return
			}	
		}
		message := map[string]any{
			"status": "fail",
			"message": "Buku tidak dapat ditemukan",
		}
		res.ErrorJson(w, message, 404)
		return
	}
}