package main

import (
	"github.com/goplus/llgo/c/sfml/system/time"
)

func main() {
	t := time.Time{}
	time.Init(&t)
	seconds := time.AsSeconds(&t)
	println("seconds", seconds)
}
