package main

import (
	"net/http"

	"github.com/adoublef/nice-wildcat/ui"
)

func main() {
	http.Handle("/", ui.Handler)
	http.ListenAndServe(":8080", nil)
}
