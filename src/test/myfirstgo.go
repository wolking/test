package main

import (
	"fmt"
	"service"
)

func main() {
	var array = make([]int, 10, 30)
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	array[5] = 6
	array[6] = 7
	//service.ForEachArray(array)
	//service.RangeArray(array)
	service.SliceTest(array)
	fmt.Println("")
}
