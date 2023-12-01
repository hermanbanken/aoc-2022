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

func TestReadBoth(t *testing.T) {
	os.Args = []string{"", "input0.txt"}
	_, _, _ = read()

	// Puzzle 2
	os.Args = []string{"", "input1.txt"}
	_, _, _ = read()
}

func TestT(t *testing.T) {
	assert.Equal(t, lib.Coord{X: 3, Y: 0}, transp(lib.Coord{X: 0, Y: 0}, 1, 4))
	assert.Equal(t, lib.Coord{X: 3, Y: 3}, transp(lib.Coord{X: 0, Y: 0}, 2, 4))
	assert.Equal(t, lib.Coord{X: 0, Y: 3}, transp(lib.Coord{X: 0, Y: 0}, 3, 4))
	assert.Equal(t, lib.Coord{X: 0, Y: 0}, transp(lib.Coord{X: 0, Y: 0}, 4, 4))

	assert.Equal(t, lib.Coord{X: 2, Y: 1}, transp(lib.Coord{X: 1, Y: 1}, 1, 4))
	assert.Equal(t, lib.Coord{X: 2, Y: 2}, transp(lib.Coord{X: 1, Y: 1}, 2, 4))
	assert.Equal(t, lib.Coord{X: 1, Y: 2}, transp(lib.Coord{X: 1, Y: 1}, 3, 4))
	assert.Equal(t, lib.Coord{X: 1, Y: 1}, transp(lib.Coord{X: 1, Y: 1}, 4, 4))
}

func TestOffset(t *testing.T) {
	p := lib.Coord{X: 1, Y: 150}
	offset := lib.Coord{X: 50, Y: 0}

	norm := p.AddR(offset.MultR(-1))
	sideCoord := norm.DivR(50)
	sidePos := norm.ModR(50)

	assert.Equal(t, lib.Coord{X: -1, Y: 3}, sideCoord)
	assert.Equal(t, lib.Coord{X: 1, Y: 0}, sidePos)
}
