package main

import (
	"aoc/lib"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	os.Args = []string{"", "input0.txt"}
	m, _ := read()
	m.facing = 2
	assert.Equal(t, Stance{lib.Coord{X: 4, Y: 4}, 1}, m.move())
	m.facing = 3
	assert.Equal(t, Stance{lib.Coord{X: 3, Y: 4}, 1}, m.move())
	m.facing = 0
	m.pos = lib.Coord{X: 12, Y: 2}
	assert.Equal(t, Stance{lib.Coord{X: 15, Y: 10}, 1}, m.move())
}
