package add

//#cgo CFLAGS: -I./
//#cgo LDFLAGS: -L${SRCDIR}/ -ladd
//
//#include "add.h"
import "C"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
