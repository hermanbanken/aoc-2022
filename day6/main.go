package main

import (
	"aoc/lib"
	"fmt"
	"io"
	"sort"
	"unicode"
)

func main() {
	r := lib.Reader()
	defer r.Close()

	var buf []byte = make([]byte, 14)
	var i = 0
	var waitingForEndOfLine = false
	for {
		read, err := r.Read(buf[0:1])
		if read == 0 || err == io.EOF {
			break
		}
		if waitingForEndOfLine && !unicode.IsSpace(rune(buf[0])) {
			continue
		}
		if unicode.IsSpace(rune(buf[0])) {
			i = 0
			waitingForEndOfLine = false
			fmt.Println("")
			continue
		}
		fmt.Print(string(buf[0]))
		i += read
		if i >= len(buf) && unique(buf) {
			waitingForEndOfLine = true
			fmt.Println(" ", i)
		}
		copy(buf[1:], buf[0:len(buf)-1])
	}
}

var room []byte

func unique(buf []byte) bool {
	if len(room) < len(buf) {
		room = make([]byte, len(buf))
	}
	copy(room, buf)

	sort.Sort(byteSort(room))
	var last byte
	for i, next := range room[0:len(buf)] {
		if i > 0 && next == last {
			return false
		}
		last = next
	}
	return true
}

type byteSort []byte

// Len implements sort.Interface
func (bs byteSort) Len() int {
	return len(bs)
}

// Less implements sort.Interface
func (bs byteSort) Less(i int, j int) bool {
	return bs[i] < bs[j]
}

// Swap implements sort.Interface
func (bs byteSort) Swap(i int, j int) {
	var tmp = bs[j]
	bs[j] = bs[i]
	bs[i] = tmp
}

var _ sort.Interface = byteSort([]byte{})
