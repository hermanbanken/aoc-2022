package main

import (
	"aoc/lib"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	// Read stacks
	var file *File
	var isListing = false
	for scanner.Scan() {
		t := scanner.Text()
		if strings.TrimSpace(t) == "" {
			break
		}

		if isListing && strings.HasPrefix(t, "$") {
			isListing = false
		}

		if isListing {
			if strings.HasPrefix(t, "dir ") {
				file.Descend(strings.TrimPrefix(t, "dir "))
			} else {
				parts := strings.Split(t, " ")
				f := file.Descend(parts[1])
				f.Size, _ = strconv.Atoi(parts[0])
				f.Dir = false
			}
		}

		// cd command
		if strings.HasPrefix(t, "$ cd ") {
			dir := strings.TrimPrefix(t, "$ cd ")
			if strings.HasPrefix(dir, "/") {
				file = &File{Name: "/", Dir: true}
			} else if dir == ".." {
				file = file.Up()
			} else {
				file = file.Descend(dir)
			}
		}
		if t == "$ ls" {
			isListing = true
		}
	}

	file = file.Parents()[0]
	fmt.Println(file.String(0))
	file.Walk(func(f *File) {
		if !f.Dir {
			for _, p := range f.Parents() {
				p.Size += f.Size
			}
		}
	})

	// Part 1
	sum := 0
	file.Walk(func(f *File) {
		if f.Dir && f.Size <= 100000 {
			sum += f.Size
		}
	})
	fmt.Println("sum ", sum)

	// Part 2
	available := 70000000 - file.Size
	needed := 30000000
	deleted := needed - available
	log.Println(available, needed, deleted)

	sizes := []int{}
	file.Walk(func(f *File) {
		if f.Dir {
			sizes = append(sizes, f.Size)
		}
	})
	sort.Ints(sizes)

	for _, s := range sizes {
		if s < deleted {
			continue
		}
		fmt.Println(s)
		break
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type File struct {
	Name   string
	Dir    bool
	Size   int
	Nested []*File
	Parent *File
}

func (f *File) Up() *File {
	return f.Parent
}

func (f *File) Descend(child string) *File {
	for _, f := range f.Nested {
		if f.Name == child {
			return f
		}
	}
	nested := &File{Name: child, Dir: true, Parent: f}
	f.Nested = append(f.Nested, nested)
	return nested
}

func (f *File) Walk(fn func(f *File)) {
	fn(f)
	for _, c := range f.Nested {
		c.Walk(fn)
	}
}

func (f *File) Parents() []*File {
	if f.Parent != nil {
		return append(f.Parent.Parents(), f.Parent)
	}
	return nil
}

func (f *File) String(indent int) string {
	out := strings.Repeat("  ", indent) + "- " + f.Name
	if f.Dir {
		out += " (dir)"
	} else {
		out += fmt.Sprintf(" (file, size=%d)", f.Size)
	}
	out += "\n"
	for _, n := range f.Nested {
		out += n.String(indent + 1)
	}
	return out
}
