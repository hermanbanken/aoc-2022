package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {
	r := reader()
	defer r.Close()

	var maxCalories []int = []int{}
	var elfCalories int = 0

	next := func() {
		if len(maxCalories) == 0 || elfCalories > maxCalories[0] {
			maxCalories = append(maxCalories, elfCalories)
			sort.Ints(maxCalories)
			if len(maxCalories) > 3 {
				maxCalories = maxCalories[len(maxCalories)-3:]
			}
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
		elfCalories += int(i)
	}
	next()

	fmt.Println("top:", last(maxCalories))
	fmt.Println("sum(3):", sum(maxCalories))

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

func last[T any](in []T) (out T) {
	if len(in) > 0 {
		return in[len(in)-1]
	}
	return
}

func sum[T constraints.Integer](in []T) (out T) {
	var sum T
	for _, c := range in {
		sum += c
	}
	return sum
}
