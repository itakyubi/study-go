package main

/*
#cgo LDFLAGS: -L . -lapi -lstdc++
#cgo CFLAGS: -I ./lib
#include "api.h"
*/
import "C"

func main() {
	C.receive(C.CString("123123321"))
	C.init(5)
	C.receive(C.CString("123123321"))
}
