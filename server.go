package server

import (
        "fmt"
        "net/http"
		"mymodule/urls"
)

func WebpagesHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Webpages")
}

func BoardsHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Boards")
}
