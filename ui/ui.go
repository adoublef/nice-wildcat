//go:generate npm run build
package ui

import (
	"embed"
	"net/http"

	"github.com/adoublef/prelude/http/fs"
)

//go:embed all:out/*
var embedFS embed.FS
var fsys = fs.Must(embedFS, "out")

var Handler = http.FileServer(http.FS(fsys))
