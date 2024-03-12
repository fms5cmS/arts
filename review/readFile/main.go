package main

import "runtime"

var maxGoroutines int

func main() {
	maxGoroutines = runtime.NumCPU()
}
