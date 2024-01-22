package main

import (
	"fmt"
	"runtime"
)

func NbCPU() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Number of logical CPUs: %d\n", numCPU)
}
