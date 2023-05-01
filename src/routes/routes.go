package routes

import (
	"net/http"

	"github.com/yusrililhm/go-bookshelf-api/src/controller"
)

func Routes()  {
	http.HandleFunc("/book", controller.Create)
	http.HandleFunc("/books", controller.Read)
}