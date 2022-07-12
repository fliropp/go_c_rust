package main

/*
#cgo LDFLAGS: ./go_rust/target/release/liblib.a
#include "./go_rust/dist/lib.h"
#include "./goC/funk.c"
*/
import "C"

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

type GData struct {
	size int
	list []int32
}

func main() {
	fmt.Println("BASIC FUNCTION")
	gBefore := time.Now()
	en := goLoop()
	fmt.Println(fmt.Sprintf("go loop took: %d (res: %d)", time.Since(gBefore), en))
	cBefore := time.Now()
	to := C.cLoop()
	fmt.Println(fmt.Sprintf("c loop took: %d (res: %d)", time.Since(cBefore), to))
	rBefore := time.Now()
	tre := C.rLoop()
	fmt.Println(fmt.Sprintf("rust loop took: %d (res: %d)", time.Since(rBefore), tre))

	fmt.Println("\nBUBBLE SORT")
	glaize, claize := genRandSlice(500)
	rlaize := make([]int32, len(glaize))
	copy(rlaize, glaize)

	//--------GO---------------
	fmt.Println("go_go:")
	runAsGo(glaize)

	//--------C----------------
	fmt.Println("\nc_go:")
	runAsC(claize)

	//--------RUST-------------
	fmt.Println("r_go:")
	runAsRust(rlaize)

}

func runAsGo(glaize []int32) {
	gints := GData{
		size: len(glaize),
		list: glaize,
	}
	fmt.Println(gints.list)
	gBefore := time.Now()
	gBubbleSort(gints)
	fmt.Println(fmt.Sprintf("go bubble sort took: %d", time.Since(gBefore)))
	fmt.Println(gints.list)
}

func runAsC(cints []C.int) {
	cd := C.createCData(&C.struct_CData{}, (*C.int)(&cints[0]), C.int(len(cints)))
	defer C.free(unsafe.Pointer(cd.list))
	lzt := (*[1 << 30]C.int)(unsafe.Pointer(cd.list))[:len(cints):len(cints)]
	fmt.Println(lzt)
	cBefore := time.Now()
	C.cBubbleSort(cd, C.int(len(cints)))
	fmt.Println(fmt.Sprintf("c bubble sort took: %d", time.Since(cBefore)))
	lzt = (*[1 << 30]C.int)(unsafe.Pointer(cd.list))[:len(cints):len(cints)]
	fmt.Println(lzt)
}

func runAsRust(slaize []int32) {
	prints := (*C.int)(C.malloc(C.size_t(len(slaize))))
	defer C.free(unsafe.Pointer(prints))
	rints := (*[1 << 30]C.int)(unsafe.Pointer(prints))[:len(slaize):len(slaize)]
	for i, d := range slaize {
		rints[i] = C.int(d)
	}
	rd := C.new_rd_data((*C.int)(&rints[0]), C.ulong(len(rints)))
	lzt := (*[1 << 30]C.int)(unsafe.Pointer(rd.list))[:len(rints):len(rints)]
	fmt.Println(lzt)

	rBefore := time.Now()
	C.rBubbleSort(&rd, C.ulong(len(rints)), C.ulong(len(rints)))
	fmt.Println(fmt.Sprintf("rust bubble sort took: %d", time.Since(rBefore)))
	lzt = (*[1 << 30]C.int)(unsafe.Pointer(rd.list))[:len(rints):len(rints)]
	fmt.Println(lzt)
}

func goLoop() int {
	s := 0
	for i := 0; i < 10001; i++ {
		s = s + i
	}
	return s
}

func gBubbleSort(gdata GData) {
	for i := 0; i < len(gdata.list)-1; i++ {
		for j := 0; j < len(gdata.list)-i-1; j++ {
			if gdata.list[j] > gdata.list[j+1] {
				gdata.list[j], gdata.list[j+1] = gdata.list[j+1], gdata.list[j]
			}
		}

	}
}

func genRandSlice(size int) ([]int32, []C.int) {
	rand.Seed(time.Now().UnixNano())
	gs := make([]int32, size)
	cs := make([]C.int, size)

	for i := 0; i < size; i++ {
		n := rand.Intn(100)
		gs[i] = int32(n)
		cs[i] = C.int(int32(n))
	}
	return gs, cs
}
