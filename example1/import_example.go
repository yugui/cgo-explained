package main

import (
	//#cgo LDFLAGS: -lm
	//#include <math.h>
	"C"
	"fmt"
)

func printSqrt(n int) {
	fmt.Println(C.sqrt(C.double(n)))
}
