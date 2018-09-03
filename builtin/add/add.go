package add

/*
#include "add.h"
#include <stdio.h>
#include <stdlib.h>

int mul(int a, int b) {
    return a * b;
}
*/
import "C"

import (
	"unsafe"
)

func HelloWorld() {
	s := C.CString("Hello World!")
	C.puts(s)
	C.free(unsafe.Pointer(s))
}

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

func Mul(a, b int) int {
	return int(C.mul(C.int(a), C.int(b)))
}
