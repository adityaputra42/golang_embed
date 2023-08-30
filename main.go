package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed wallpaper.jpg
var wallpaper []byte

// go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("wallpaper_next.jpg", wallpaper, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println("entry name", entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("content => ", string(content))

		}
	}
}
