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

	m.facing = 0 // right
	assert.Equal(t, Stance{lib.Coord{X: 9, Y: 0}, 0}, m.nextPart2())

	m.facing = 1 // down
	assert.Equal(t, Stance{lib.Coord{X: 8, Y: 1}, 1}, m.nextPart2())

	m.facing = 2 // left, ort, anti-clockwise
	assert.Equal(t, Stance{lib.Coord{X: 4, Y: 4}, 1}, m.nextPart2())

	m.facing = 3 // up
	assert.Equal(t, Stance{lib.Coord{X: 3, Y: 4}, 1}, m.nextPart2())

	m.facing = 0 // right, ort, clockwise
	m.pos = lib.Coord{X: 11, Y: 5}
	assert.Equal(t, Stance{lib.Coord{X: 15, Y: 8}, 1}, m.nextPart2())

	m.facing = 3 // up, ort, anti-clockwise
	m.pos = lib.Coord{X: 12, Y: 8}
	assert.Equal(t, Stance{lib.Coord{X: 11, Y: 7}, 2}, m.nextPart2())

	m.facing = 3 // up, ort, clockwise
	m.pos = lib.Coord{X: 6, Y: 4}
	assert.Equal(t, Stance{lib.Coord{X: 8, Y: 2}, 0}, m.nextPart2())

	m.facing = 1 // down
	m.pos = lib.Coord{X: 8, Y: 11}
	assert.Equal(t, Stance{lib.Coord{X: 3, Y: 7}, 3}, m.nextPart2())

	m.facing = 0 // right at edge, going to bottom right
	m.pos = lib.Coord{X: 11, Y: 2}
	assert.Equal(t, Stance{lib.Coord{X: 15, Y: 9}, 2}, m.nextPart2())
}
