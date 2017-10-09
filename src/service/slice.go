package service

import (
	"fmt"
	"strconv"
)

func SliceTest(arr []int) {
	fmt.Println("arr.len=" + strconv.Itoa(len(arr)))
	fmt.Println("arr.cap=" + strconv.Itoa(cap(arr)))
	mySlice := arr[5:80]
	fmt.Println("slice cap=" + strconv.Itoa(cap(mySlice)))
	for _, v := range mySlice {
		fmt.Print(strconv.Itoa(v) + ",")
	}
}
