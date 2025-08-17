package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	arr := [5]int{1, 2, 3, 0, 0}

	s := slice{
		array: unsafe.Pointer(&arr[0]), 
		len:   3,                        
		cap:   len(arr),                 
	}

	fmt.Printf("slice struct: %+v\n", s)

	realArr := (*[5]int)(s.array)
	fmt.Println("Underlying array:", *realArr)

	fmt.Println("First 3 elements:", (*realArr)[:s.len])
}
