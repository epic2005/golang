
package main

import (
	"fmt"
)

func mkslice() []int {
	s := make([]int, 0, 10)
	s = append(s, 200)
	return s
}

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 100
	return m
}


func main() {
	s := mkslice()
	fmt.Printf("slice 0 is : %d\n", s[0])

	m := mkmap()
	fmt.Printf("map a is: %d\n", m["a"])
}