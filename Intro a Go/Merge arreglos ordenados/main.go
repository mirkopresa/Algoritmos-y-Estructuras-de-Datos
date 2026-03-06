package main

import "fmt"

func Merge(arr1, arr2 []int) []int {
	arr3 := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] >= arr2[j] {
			arr3 = append(arr3, arr2[j])
			j++
		} else {
			arr3 = append(arr3, arr1[i])
			i++
		}
	}
	arr3 = append(arr3, arr1[i:]...)
	arr3 = append(arr3, arr2[j:]...)
	return arr3
}

func main() {
	arr1 := []int{1, 2, 3, 40}
	arr2 := []int{4, 5, 6}
	fmt.Println(Merge(arr1, arr2))
}
