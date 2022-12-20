package main

import (
	"aoc/lib"
	"fmt"
)

type Num struct {
	originalIdx int
	value       int
}

func main() {
	nums := []*Num{}
	idx := 0
	lib.EachLine(func(line string) {
		nums = append(nums, &Num{originalIdx: idx, value: lib.Int(line)})
		idx++
	})

	circ := []*Num{}
	for i := range nums {
		circ = append(circ, nums[i])
	}
	for i := range nums {
		num := nums[i]
		pos := indexOf(circ, num)
		newPos := (pos + num.value)
		if newPos >= len(nums) {
			newPos %= len(nums)
			newPos += 1
		}
		if newPos < 0 {
			// log.Println("neg", num.value, newPos, newPos%len(nums)+len(nums))
			newPos %= len(nums)
			newPos += len(nums) - 1
		}
		// log.Println(circ, num.value, "m", pos, "->", newPos)
		if pos != newPos {
			circ = add(remove(circ, pos), newPos, num)
		}
		// for i := range circ {
		// 	fmt.Print(circ[i].value, ", ")
		// }
		// fmt.Println()
	}

	for i := range circ {
		if circ[i].value == 0 {
			fmt.Println(circ[(i+1000)%len(circ)].value, circ[(i+2000)%len(circ)].value, circ[(i+3000)%len(circ)].value)
			fmt.Println(circ[(i+1000)%len(circ)].value + circ[(i+2000)%len(circ)].value + circ[(i+3000)%len(circ)].value)
		}
	}
	// not 1608 -8025 3733 => -2684
}

func indexOf(list []*Num, el *Num) int {
	for pp, pointer := range list {
		if pointer == el {
			return pp
		}
	}
	panic("not found")
}

func remove(list []*Num, idx int) []*Num {
	removed := append(list[0:idx], list[idx+1:]...)
	// log.Println(removed, "removed")
	return removed
}

func add(list []*Num, idx int, new *Num) []*Num {
	res := make([]*Num, len(list)+1)
	copy(res[0:idx], list[0:idx])
	res[idx] = new
	copy(res[idx+1:], list[idx:])
	// log.Println(res, "added")
	return res
}
