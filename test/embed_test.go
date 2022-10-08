package test

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

// TESTING ERROR

//go:embed version.txt
var version string

func TestVersionString(t *testing.T) {
	fmt.Println(version)
}

// go:embed uwu.jpg
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("uwu_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

}

//go:embed test/files/a.txt
//go:embed test/files/b.txt
//go:embed test/files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed .../files/*
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
