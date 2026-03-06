package main

import "fmt"

func BuscarElem(arr []int, elem int) int {
	// Otra manera de hacerlo
	// for index, value := range arr {
	// 	if elem == value {
	//		return index
	// 	}
	// }

	for i := 0; i < len(arr); i++ {
		if elem == arr[i] {
			return i
		}
	}
	return -1
}

func main() {
	array := []int{20, 50, 4, 77}
	fmt.Println(BuscarElem(array, 7))
}
