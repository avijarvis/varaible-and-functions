package main

import "fmt"

// Custom iterator
type Iterator struct {
	data []int
	pos  int
}

// Next returns the next value and a boolean indicating if iteration is complete
func (it *Iterator) Next() (int, bool) {
	if it.pos >= len(it.data) {
		return 0, false // No more elements
	}
	val := it.data[it.pos]
	it.pos++
	return val, true
}

func main() {
	it := &Iterator{data: []int{5, 10, 15}}

	for {
		val, ok := it.Next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
