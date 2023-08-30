package test

import (
	"embed"
	// _ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {

	fmt.Println(version)

}

//go:embed wallpaper.jpg
var wallpaper []byte

func TestByteArray(t *testing.T) {
	err := os.WriteFile("wallpaper_next.jpg", wallpaper, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
var files embed.FS

func TestMultiFile(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

}

// go:embed files/*.txt
var path embed.FS

func TestPathMatch(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println("entry name", entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("content => ", string(content))

		}
	}

}
