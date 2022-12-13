package lib

import (
	"bufio"
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

func Lines() (lines []string) {
	r := Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func EachLine(fn func(line string)) {
	r := Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fn(scanner.Text())
	}
}
