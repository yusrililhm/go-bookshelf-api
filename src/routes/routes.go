package routes

import (
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/handler"
)

func Routes()  {
	http.HandleFunc("/book", handler.Create)
	http.HandleFunc("/books", handler.Read)
}