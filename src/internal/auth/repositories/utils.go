package repositories

import (
	"embed"
	"fmt"
)

//go:embed queries/*
var fs embed.FS

func readFromFs(path string) string {
	content, err := fs.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("error loading query: %w", err))
	}

	return string(content)
}
