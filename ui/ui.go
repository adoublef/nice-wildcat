package ui

import (
	"embed"
	"net/http"

	fs "github.com/adoublef/nice-wildcat"
)

//go:embed all:dist/*
var embedFS embed.FS
var fsys = fs.Must(embedFS, "dist")

var Handler = http.FileServer(http.FS(fsys))
