package main

import (
	"net/http"

	"github.com/vghessel/web_app/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
