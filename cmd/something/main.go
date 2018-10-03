package main

import (
	"github.com/hyangah/testmod/m1/v3"
	"github.com/hyangah/testmod/m2/v2"
	"github.com/hyangah/testmod/m3"
)

func main() {
	println("m1: ", m1.Version)
	println("m2: ", m2.Version)
	println("m3: ", m3.Version)
}

