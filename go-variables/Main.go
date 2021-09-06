package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	c = iota
)

func main() {
	fileSize := 400000000.
	fmt.Printf("%.2fGB", fileSize/GB)
}
