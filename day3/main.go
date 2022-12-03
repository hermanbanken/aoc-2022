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

	var sum1 = 0
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
	// mask = make([]bool, 52)

	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t == "" {
			break
		}

		idx := strings.IndexAny(t[0:len(t)/2], t[len(t)/2:])
		log.Println("prio", prio(t[idx]), "item", string(t[idx]))
		sum1 += int(prio(t[idx]))
	}
	fmt.Println("result1:", sum1)
	// fmt.Println("result2:", sum2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func prio(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 'a' + 1
	}
	return b - 'A' + 27
}
