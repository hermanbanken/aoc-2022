package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	// Read stacks
	var stacks [][]byte
	var cols int
	for scanner.Scan() {
		t := scanner.Text()
		if strings.TrimSpace(t) == "" || strings.HasPrefix(t, " 1") {
			break
		}

		cols = (len(t) + 1) / 4
		for len(stacks) < cols {
			stacks = append(stacks, nil)
		}
		for c := 0; c < cols; c++ {
			if t[c*4] == '[' {
				stacks[c] = append(stacks[c], t[c*4+1])
			}
		}
	}
	log.Println(stacks)

	// Read moves
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		fmt.Printf("%q\n", t)
		if t == "" {
			continue
		}
		p := strings.Split(t, " ")
		count, _ := strconv.Atoi(p[1])
		src, _ := strconv.Atoi(p[3])
		dst, _ := strconv.Atoi(p[5])
		move(count, src-1, dst-1, stacks)
	}

	for _, s := range stacks {
		fmt.Print(string(s[0:1]))
	}
	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func move(count, src, dst int, stacks [][]byte) {
	for c := 0; c < count; c++ {
		fmt.Println("moving", string(stacks[src][0:1]), "from", src, "to", dst)
		stacks[dst] = append([]byte{stacks[src][0]}, stacks[dst]...)
		stacks[src] = stacks[src][1:]
		fmt.Println(string(stacks[dst]), string(stacks[src]))
	}
}
