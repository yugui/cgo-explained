package main

// #include "use_exported.h"
import "C"

func main() {
	C.print_go_version()
}
