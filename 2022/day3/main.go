package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()

	var sum = 0
	scanner := bufio.NewScanner(r)
	/*var mask []bool
	i := 0
	// log.Println(len(t), len(t)/2)
	for ; i < len(t)/2; i++ {
		mask[prio(t[i])] = true
	}
	// log.Println(mask)
	for ; i < len(t); i++ {
		if mask[prio(t[i])] {
			sum += int(prio(t[i]))
			log.Println("prio", prio(t[i]), "item", string(t[i]))
			break
		}
	}*/

	var masks = [][]bool{make([]bool, 53), make([]bool, 53)}
	line := 0
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			break
		}

		if line%3 == 2 {
			for _, letter := range t {
				if masks[0][prio(letter)] && masks[1][prio(letter)] {
					log.Println("match", string(letter))
					sum += int(prio(letter))
					break
				}
			}
			masks = [][]bool{make([]bool, 53), make([]bool, 53)}
		} else {
			for _, letter := range t {
				masks[line%3][prio(letter)] = true
			}
		}
		line++
	}
	fmt.Println("result:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func prio(b rune) rune {
	if b >= 'a' && b <= 'z' {
		return b - 'a' + 1
	}
	return b - 'A' + 27
}
