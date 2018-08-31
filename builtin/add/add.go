package add

/*
#include "add.h"

int mul(int a, int b) {
    return a * b;
}
*/
import "C"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

func Mul(a, b int) int {
	return int(C.mul(C.int(a), C.int(b)))
}
