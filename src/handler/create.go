package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/yusrililhm/go-bookshelf-api/src/model"
	"github.com/yusrililhm/go-bookshelf-api/src/res"
)

func Create(w http.ResponseWriter, r *http.Request)  {
	
	var err error
	
	if r.Method == "POST" {
		
		w.Header().Set("Content-Type", "application/json")

		decode := json.NewDecoder(r.Body)

		// generate unique id

		id, err := exec.Command("uuidgen").Output()
		
		if err != nil {
			log.Fatal(err.Error())
		}
		
		// date manipulation

		t := time.Now()
		insertedAt := t.Format(time.RFC1123)
		updatedAt := insertedAt

		// body request

		name := r.FormValue("name")
		year := r.FormValue("year")
		author := r.FormValue("author")
		summary := r.FormValue("summary")
		publisher := r.FormValue("publisher")
		pageCount := r.FormValue("pageCount")
		readPage := r.FormValue("readPage")
		reading := r.FormValue("reading")

		// convert type from string

		years, _ := strconv.Atoi(year)
		pageCounts, _ := strconv.Atoi(pageCount)
		readPages, _ := strconv.Atoi(readPage)
		readings, _ := strconv.ParseBool(reading)

		// finished = pageCount == readPage

		finished := pageCounts == readPages

		payload := model.Book{
			Id: string(id),
			Name: name,
			Year: years,
			Author: author,
			Summary: summary,
			Publisher: publisher,
			PageCount: pageCounts,
			ReadPage: readPages,
			Reading: readings,
			Finished: finished,
			InsertedAt: insertedAt,
			UpdateAt: updatedAt,
		}

		if err := decode.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		model.Books = []model.Book{}

		model.Books = append(model.Books, payload)

		// error jika nama tidak dilampirkan

		if payload.Name == "" {
			message := map[string]interface{}{
				"status": "fail",
				"message": "Gagal menambahkan buku. Mohon isi nama buku",
			}

			res.ErrorJson(w, message, 400)
			return
		}

		// error jika readCount > pageCount

		if payload.ReadPage > payload.PageCount {
			message := map[string]interface{}{
				"status": "fail",
				"message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
			}

			res.ErrorJson(w, message, 400)
			return
		}

		// response jika buku berhasil ditambahkan

		message := map[string]interface{}{
			"status": "success",
			"message": "Buku berhasil ditambahkan",
			"data": map[string]interface{}{
				"id": payload.Id,
			},
		}

		encodes, err := json.Marshal(message)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(encodes)
		return
	}
	http.Error(w, err.Error(), 500)
}