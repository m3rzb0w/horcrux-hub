package main

import (
	"embed"
	"io/fs"
)

//go:embed front/dist
var embedFrontend embed.FS

func getFrontendAssets() fs.FS {
	f, err := fs.Sub(embedFrontend, "front/dist")
	if err != nil {
		panic(err)
	}
	return f
}
