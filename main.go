package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

var (
	source   string
	target   string
	oldValue string
	newValue string
)

func main() {
	flag.StringVar(&source, "source", "", "Absolute path for source textual file")
	flag.StringVar(&target, "target", "", "Absolute path for target textual file")
	flag.StringVar(&oldValue, "old", "", `old content`)
	flag.StringVar(&newValue, "new", "", "new content")
	flag.Parse()

	if len(source) < 1 {
		log.Fatal(`"source" must be provided!`)
	}

	if len(target) < 1 {
		log.Fatal(`"target" must be provided!`)
	}

	if len(oldValue) < 1 {
		log.Fatal(`"old" must be provided!`)
	}

	if len(newValue) < 1 {
		log.Fatal(`"new" must be provided!`)
	}

	content, err := ReadFileAsString(source)
	if err != nil {
		log.Fatal("source file not found")
	}

	err = WriteIOReaderToFile(target, strings.NewReader(strings.ReplaceAll(content, oldValue, newValue)))
	if err != nil {
		log.Fatal("target file cannot be written")
	}
}

// ReadFile read a file
func ReadFile(path string, onFileRead func(*os.File) error) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	return onFileRead(file)
}

// CreateFile create file
func CreateFile(path string, onFileCreate func(*os.File) error) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return onFileCreate(file)
}

// ReadFileAsString read file as string
func ReadFileAsString(path string) (string, error) {
	buf := new(bytes.Buffer)

	err := ReadFile(path, func(file *os.File) error {
		_, err := buf.ReadFrom(file)
		return err
	})

	if err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func WriteIOReaderToFile(path string, reader io.Reader) error {
	return CreateFile(path, func(file *os.File) error {
		_, err := io.Copy(file, reader)
		return err
	})
}
