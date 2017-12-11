package main

// #cgo CFLAGS: -Wall -Wextra -Wno-unused -I.
// #include <foo.h>
import "C"
import "fmt"

func main() {
	v := 2

	fmt.Println("-- in Go code -- about to call C")
	C.foo(C.int(v))
	fmt.Println("-- back in Go code")
}
