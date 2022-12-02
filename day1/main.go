package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := lib.Reader()
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

	fmt.Println("top:", lib.Last(maxCalories))
	fmt.Println("sum(3):", lib.Sum(maxCalories))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
