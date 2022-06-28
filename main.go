package main

/*
#cgo LDFLAGS: ./go_rust/target/release/libgo_rust.a
#include "./go_rust/dist/lib.h"
#include "./goC/funk.c"
*/
import "C"

import (
	"fmt"
	"math/rand"
	"time"
)

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

	fmt.Println("BUBBLE SORT")
	slaize, _ := genRandSlice(9)
	glaize := make([]int32, len(slaize))
	/*claize := make([]int32, len(slaize))
	rlaize := make([]int32, len(slaize))*/

	copy(glaize, slaize)
	/*copy(claize, slaize)
	copy(rlaize, slaize)*/

	fmt.Println("go_go:")
	fmt.Println(glaize)
	gBefore = time.Now()
	gBubbleSort(glaize)
	fmt.Println(fmt.Sprintf("go bubble sort took: %d", time.Since(gBefore)))
	fmt.Println(glaize)

	fmt.Println("c_go:")
	fmt.Println(slaize)
	cBefore = time.Now()
	var cints C.struct_CData
	for i, d := range slaize {
		cints.list[i] = C.int(d)
	}
	//C.cBubbleSort((*C.int)(&claize[0]), C.int(len(claize)))
	C.cBubbleSort(&cints, C.int(len(slaize)))
	fmt.Println(fmt.Sprintf("c bubble sort took: %d", time.Since(cBefore)))
	fmt.Println(cints.list)

	fmt.Println("r_go:")
	fmt.Println(slaize)
	var crints C.struct_Data
	for i, d := range slaize {
		crints.list[i] = C.int(d)
	}
	rBefore = time.Now()
	_ = C.rBubbleSort(&crints)
	fmt.Println(fmt.Sprintf("rust bubble sort took: %d", time.Since(rBefore)))
	fmt.Println(crints.list)

}

func goLoop() int {
	s := 0
	for i := 0; i < 10001; i++ {
		s = s + i
	}
	return s
}

func gBubbleSort(list []int32) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}

	}
}

func genRandSlice(size int) ([]int32, *[]C.int) {
	rand.Seed(time.Now().UnixNano())
	gs := make([]int32, size)
	cs := make([]C.int, size)

	for i := 0; i < size; i++ {
		gs[i] = int32(rand.Intn(100))
		cs[i] = C.int(int32(rand.Intn(100)))
	}
	return gs, &cs

}
