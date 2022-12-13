package main

import (
	"aoc/lib"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {
	r := lib.Reader()
	defer r.Close()
	scanner := bufio.NewScanner(r)

	var index = 1
	var sumIndexes = 0
	for scanner.Scan() {
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()

		if rightOrder(a, b) {
			sumIndexes += index
		}
		index++
	}

	fmt.Println(sumIndexes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func rightOrder(astr, bstr string) (res bool) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e, astr, bstr)
			res = false
		}
	}()

	var a, b interface{}
	err := json.Unmarshal([]byte(astr), &a)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bstr), &b)
	if err != nil {
		panic(err)
	}

	fmt.Println(astr)
	fmt.Println(bstr)
	v := isLess(a, b)
	fmt.Println(v)
	fmt.Println()
	return v <= 0
}

func isLess(a interface{}, b interface{}) int {
	fa, AisFloat := a.(float64)
	fb, BisFloat := b.(float64)
	if AisFloat && BisFloat {
		return compareF(int(fa), int(fb))
	}
	la := asList(a)
	lb := asList(b)

	for len(lb) > 0 {
		if len(la) == 0 {
			return -1
		}
		fmt.Println("compare", la[0], lb[0])
		if result := isLess(la[0], lb[0]); result != 0 {
			return result
		}
		la = la[1:]
		lb = lb[1:]
	}
	if len(la) == 0 {
		return 0
	}
	fmt.Println("a continues")
	return 1
}

func compareF(a, b int) int {
	fmt.Println("compareF", a, b)
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func itemRest(elem interface{}) (*int, []interface{}) {
	switch v := elem.(type) {
	case float64:
		vf := int(v)
		return &vf, nil
	case []interface{}:
		if len(v) > 0 {
			d, rest := itemRest(v[0])
			rest = append(rest, v[1:]...)
			return d, rest
		}
		return nil, nil
	default:
		panic(reflect.TypeOf(v))
	}
}

func intOrDefault(v *int) int {
	if v == nil {
		return -1
	}
	return *v
}

func rightOrderL(a, b interface{}) int {
	ahead, arest := itemRest(a)
	bhead, brest := itemRest(b)
	fmt.Printf("comparing %+v %+v\n", intOrDefault(ahead), intOrDefault(bhead))
	if ahead == nil && bhead != nil {
		return -1
	} else if ahead != nil && bhead == nil {
		return 1
	} else if ahead == nil && bhead == nil {
		fmt.Println("same length", arest, brest, a, b)
		return 0
	}

	if *ahead < *bhead {
		return -1
	}
	if *ahead > *bhead {
		return 1
	}
	return rightOrderL(arest, brest)
}

func asList(elem interface{}) []interface{} {
	switch v := elem.(type) {
	case []interface{}:
		return v
	default:
		return []interface{}{v}
	}
}

type kindPair struct {
	a reflect.Kind
	b reflect.Kind
}

/*
	for {

		var ai, bi = 0, 0
		alist := asList(a)
		blist := asList(b)

		if ai >= len(alist) {
			break
		}
		if bi >= len(blist) {
			break
		}
		if alist[ai] < blist[bi] {

		}

		kinds := kindPair{a: reflect.TypeOf(a).Kind(), b: reflect.TypeOf(b).Kind()}
		switch kinds {
		case kindPair{a: reflect.Int, b: reflect.Int}:
		case kindPair{a: reflect.Slice, b: reflect.Slice}:

		case kindPair{a: reflect.Int, b: reflect.Slice}:
		case kindPair{a: reflect.Slice, b: reflect.Int}:
		}
	}*/
