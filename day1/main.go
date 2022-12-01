package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main1() {
	r := reader()
	defer r.Close()

	var maxCalories int64 = 0
	var elfCalories int64 = 0

	next := func() {
		if elfCalories > maxCalories {
			maxCalories = elfCalories
		}
		elfCalories = 0
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			next()
			continue
		}
		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		elfCalories += i
	}
	next()
	fmt.Println(maxCalories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func reader() io.ReadCloser {
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
