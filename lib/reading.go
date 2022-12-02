package lib

import (
	"io"
	"log"
	"os"
)

func Reader() io.ReadCloser {
	if len(os.Args) < 2 {
		log.Fatal("Supply filename as first argument")
	}

	if os.Args[1] == "-" {
		return os.Stdin
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return file
}
