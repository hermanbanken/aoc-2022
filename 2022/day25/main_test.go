package main

import (
	"aoc/lib"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDec(t *testing.T) {
	pairs := [][]string{
		{"1=", "3"},
		{"12", "7"},
		{"21", "11"},
		{"111", "31"},
		{"112", "32"},
		{"122", "37"},
		{"1=-0-2", "1747"},
		{"12111", "906"},
		{"2=0=", "198"},
		{"2=01", "201"},
		{"20012", "1257"},
		{"1=-1=", "353"},
		{"1-12", "107"},
	}
	for _, p := range pairs {
		input := p[0]
		expected := lib.Int(p[1])
		t.Run(fmt.Sprintf("toDec(%s)", input), func(t *testing.T) {
			assert.Equal(t, expected, toDec(input))
		})
	}
}

func TestToSnafu(t *testing.T) {
	pairs := [][]string{
		{"1", "1"},
		{"2", "2"},
		{"3", "1="},
		{"4", "1-"},
		{"5", "10"},
		{"6", "11"},
		{"7", "12"},
		{"8", "2="},
		{"9", "2-"},
		{"10", "20"},
		{"15", "1=0"},
		{"20", "1-0"},
		{"2022", "1=11-2"},
		{"12345", "1-0---0"},
		{"314159265", "1121-1110-1=0"},
	}
	for _, p := range pairs {
		input := p[0]
		expected := p[1]
		t.Run(fmt.Sprintf("toSnafu(%s)", input), func(t *testing.T) {
			assert.Equal(t, expected, toSnafu(lib.Int(input)))
		})
	}
}
