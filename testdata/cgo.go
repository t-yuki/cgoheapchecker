package cgo

// #include <stdlib.h>
import "C"
import "unsafe"

var ptr unsafe.Pointer

func CallMalloc() {
	ptr = C.malloc(512)
}

func CallFree() {
	C.free(ptr)
}
