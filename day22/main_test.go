package main

import (
	"aoc/lib"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	origin := lib.Coord{X: 50, Y: 100}
	assert.Equal(t, lib.Coord{X: 49, Y: 50}, rotate(lib.Coord{X: 0, Y: 101}, origin, math.Pi/2))
	assert.Equal(t, lib.Coord{X: 49, Y: 149}, rotate(lib.Coord{X: 99, Y: 101}, origin, math.Pi/2))
}

func TestMove(t *testing.T) {
	os.Args = []string{"", "input0.txt"}
	m, _ := read(true)

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
	assert.Equal(t, Stance{lib.Coord{X: 14, Y: 8}, 1}, m.nextPart2())

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

	// Puzzle 2
	os.Args = []string{"", "input1.txt"}
	m, _ = read(true)

	m.facing = 3
	m.pos = lib.Coord{X: 99, Y: 0}
	assert.Equal(t, Stance{lib.Coord{X: 0, Y: 199}, 0}, m.nextPart2())
	m.facing = 2
	m.pos = lib.Coord{X: 0, Y: 199}
	assert.Equal(t, Stance{lib.Coord{X: 99, Y: 0}, 1}, m.nextPart2())

	m.facing = 3
	m.pos = lib.Coord{X: 50, Y: 0}
	assert.Equal(t, Stance{lib.Coord{X: 0, Y: 150}, 0}, m.nextPart2())
	m.facing = 2
	m.pos = lib.Coord{X: 0, Y: 150}
	assert.Equal(t, Stance{lib.Coord{X: 50, Y: 0}, 1}, m.nextPart2())

	m.facing = 3
	m.pos = lib.Coord{X: 100, Y: 0}
	assert.Equal(t, Stance{lib.Coord{X: 0, Y: 199}, 3}, m.nextPart2())
	m.facing = 1
	m.pos = lib.Coord{X: 0, Y: 199}
	assert.Equal(t, Stance{lib.Coord{X: 100, Y: 0}, 1}, m.nextPart2())

	m.facing = 3
	m.pos = lib.Coord{X: 149, Y: 0}
	assert.Equal(t, Stance{lib.Coord{X: 49, Y: 199}, 3}, m.nextPart2())
	m.facing = 1
	m.pos = lib.Coord{X: 49, Y: 199}
	assert.Equal(t, Stance{lib.Coord{X: 149, Y: 0}, 1}, m.nextPart2())
}
