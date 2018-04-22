package main

import (
	"fmt"
)

type flags byte

const (
	read flags = 1 << iota
	write
	exec
	open
	put
)

func main() {
	f := read | open | put
	fmt.Printf("%b\n", f)
}