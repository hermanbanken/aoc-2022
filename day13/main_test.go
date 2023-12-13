package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test13(t *testing.T) {
	data, err := os.ReadFile("input1.txt")
	assert.NoError(t, err)
	pz := puzzles(strings.Split(string(data), "\n"))
	assert.Equal(t, 100, solve(1, pz[1]))
}
