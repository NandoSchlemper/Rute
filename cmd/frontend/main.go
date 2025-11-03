package main

import (
	"fmt"
	"net/http"
	"rute/cmd/frontend/components"

	"github.com/a-h/templ"
)

func main() {
	component := components.Game()

	http.Handle("/", templ.Handler(component))
	fmt.Printf("Templates iniciados na porta: http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}
