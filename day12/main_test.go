package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanBeGroup(t *testing.T) {
	assert.True(t, problem{cache: map[[2]int]int{}, template: "####", chunks: []int{4}}.canBeGroup(0, 0))
	assert.False(t, problem{cache: map[[2]int]int{}, template: "####", chunks: []int{5}}.canBeGroup(0, 0))
	assert.True(t, problem{cache: map[[2]int]int{}, template: ".##.", chunks: []int{2}}.canBeGroup(1, 0))
	assert.False(t, problem{cache: map[[2]int]int{}, template: ".##.", chunks: []int{3}}.canBeGroup(1, 0))
	assert.False(t, problem{cache: map[[2]int]int{}, template: ".##.", chunks: []int{2}}.canBeGroup(2, 0))
}

func TestVariations(t *testing.T) {
	p := problem{}
	p.parse("???.### 1,1,3", false)
	assert.Equal(t, 1, p.variations(0, 0))
	p.parse("???.### 1,1,3", true)
	assert.Equal(t, 1, p.variations(0, 0))

	assert.True(t, problem{cache: map[[2]int]int{}, template: "?.###", chunks: []int{1, 3}}.canBeGroup(2, 1))
	p.parse("?.### 1,3", false)
	assert.Equal(t, 1, p.variations(2, 1))
	assert.Equal(t, 1, p.variations(0, 0))
}
