package service

import (
	"fmt"
	"strconv"
)

func ForEachArray(arr [10]int) {
	fmt.Println("--------------------------------foreach array by len(array)-------------------------------")
	if len(arr) > 0 {
		for i := 0; i < len(arr); i++ {
			fmt.Printf("arr["+strconv.Itoa(i)+"]=%v", arr[i])
			fmt.Println("")
		}
	}
}

func RangeArray(arr [10]int) {
	fmt.Println("-------------------------------foreach array by range-------------------------------------")
	if len(arr) > 0 {
		for i, v := range arr {
			fmt.Printf("arr[%i]=%v", i, v)
			fmt.Println("")
		}
	}
}
