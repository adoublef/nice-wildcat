//go:generate pnpm run build
package ui

import (
	"embed"
	"net/http"

	fs "github.com/adoublef/nice-wildcat"
)

//go:embed all:out/*
var embedFS embed.FS
var fsys = fs.Must(embedFS, "out")

var Handler = http.FileServer(http.FS(fsys))
