package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/yusrililhm/go-bookshelf-api/src/model"
)

func Create(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "applicatio/json")

	var err error
	
	if r.Method == "POST" {

		decode := json.NewDecoder(r.Body)

		// generate unique id

		id, err := exec.Command("uuidgen").Output()
		
		if err != nil {
			log.Fatal(err.Error())
		}
		
		// date manipulation

		update := time.Now()
		insertedAt := update.Format(time.RFC1123)
		updatedAt := insertedAt

		name := r.FormValue("name")
		year := r.FormValue("year")
		author := r.FormValue("author")
		summary := r.FormValue("summary")
		publisher := r.FormValue("publisher")
		pageCount := r.FormValue("pageCount")
		readPage := r.FormValue("readPage")
		reading := r.FormValue("reading")

		years, _ := strconv.Atoi(year)
		pageCounts, _ := strconv.Atoi(pageCount)
		readPages, _ := strconv.Atoi(readPage)
		readings, _ := strconv.ParseBool(reading)

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

		if payload.Name == "" {
			http.Error(w, "Gagal menambahkan buku. Mohon isi nama", 400)
			return
		}

		if payload.ReadPage > payload.PageCount {
			http.Error(w, "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount", 400)
			return
		}

		encode, err := json.Marshal(payload.Id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		message := map[string]interface{}{
			"status": "success",
			"message": "Buku berhasil ditambahkan",
			"data": map[string]interface{}{
				"id": encode,
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