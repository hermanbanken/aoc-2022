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

	fmt.Println("moving", string(stacks[src][0:count]), "from", src, "to", dst)
	var dstcopy = make([]byte, count+len(stacks[dst]))
	copy(dstcopy[0:], stacks[src][0:count])
	copy(dstcopy[count:], stacks[dst])
	stacks[dst] = dstcopy
	stacks[src] = stacks[src][count:]
	fmt.Println(string(stacks[dst]), string(stacks[src]))

}
