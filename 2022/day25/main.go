package main

import (
	"aoc/lib"
	"fmt"
	"math"
	"strconv"
)

func main() {
	sum := 0
	lib.EachLine(func(line string) {
		sum += toDec(line)
	})
	fmt.Println(toSnafu(sum))
}

var toDecMap = map[byte]int{
	'=': -2,
	'-': -1,
	'0': 0,
	'1': 1,
	'2': 2,
}

func toDec(snafu string) (out int) {
	for pow := len(snafu) - 1; pow >= 0; pow-- {
		digit := toDecMap[snafu[len(snafu)-1-pow]]
		base := int(math.Pow(5, float64(pow)))
		out += base * digit
	}
	return out
}

func toSnafu(dec int) (snafu string) {
	digits := []byte(strconv.FormatInt(int64(dec), 5))
	var carry int
	for i := len(digits) - 1; i >= 0; i-- {
		five, _ := strconv.Atoi(string(digits[i : i+1]))

		// previous carry
		if carry > 0 {
			five++
			if five == 5 {
				carry++
				digits[i] = '0'
				five = 0
			}
			carry--
		}

		if five == 1 {
			digits[i] = '1'
		}

		if five == 2 {
			digits[i] = '2'
		}

		// current digit
		if five == 3 {
			carry++
			digits[i] = '='
		}
		if five == 4 {
			carry++
			digits[i] = '-'
		}
	}
	if carry > 0 {
		digits = append([]byte(strconv.Itoa(carry)), digits...)
	}
	return string(digits)
}
