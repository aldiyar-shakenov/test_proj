package main

import "fmt"

func bubbleSort(arr []int32) []int32 {
	for i := range arr {
		if arr[i] < arr[i+1] {
			arr[i] = arr[i+1]
			arr[i+1] = arr[i]
		}
	}
	return arr
}

func main() {
	testArr := []int32{11, 10, 14, 15, 16, 19}
	fmt.Println(bubbleSort(testArr))

}
