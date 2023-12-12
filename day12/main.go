package main

import (
	"aoc/lib"
	"fmt"
	"strings"
)

func main() {
	lines := lib.Lines()
	sum := 0
	for _, line := range lines {
		fmt.Println()
		fmt.Println("puzzle", line)
		parts := strings.Fields(line)
		template := parts[0]
		chunks := lib.Map(strings.Split(parts[1], ","), lib.Int)
		result := countMatches(template, chunks)
		sum += result
		fmt.Printf("%d\n", result)
	}
	fmt.Println(sum)
}

func countMatches(template string, chunks []int) (countMatches int) {
	questions := []int{}
	for i, r := range template {
		if r == '?' {
			questions = append(questions, i)
		}
	}
	for option := 0; option < 1<<len(questions); option++ {
		digit := option
		instance := template
		for _, pos := range questions {
			if digit%2 == 0 {
				instance = instance[:pos] + "." + instance[pos+1:]
			} else {
				instance = instance[:pos] + "#" + instance[pos+1:]
			}
			digit = digit >> 1
		}
		if matches(instance, chunks) {
			countMatches += 1
			// fmt.Println(instance, "match")
		} else {
			// fmt.Println(instance, "no match")
		}
	}
	return
}

func matches(instance string, chunks []int) bool {
	// var chunkIdx = 0
	parts := lib.Filter(strings.Split(instance, "."), func(s string) bool { return s != "" })
	if len(parts) != len(chunks) {
		return false
	}
	for i, part := range parts {
		if len(part) == chunks[i] {
			continue
		}
		return false
	}
	return true
	// fmt.Println(strings.Join(strings.Split(instance, "."), ","))
	// i := 0
	// for {
	// 	if instance[i] == '#' {
	// 		for j := i + 1; j < chunks[chunkIdx]; j++ {
	// 			if instance[j] == '.' {
	// 				return false
	// 			}
	// 		}
	// 		if instance[i+chunks[chunkIdx]] != '.' {
	// 			return false
	// 		}
	// 		i += chunks[chunkIdx] + 1
	// 	} else {
	// 		i += 1
	// 	}
	// 	if i >= len(instance) {
	// 		return true
	// 	}
	// }
}
