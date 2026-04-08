// Implementar un algoritmo que permita ordenar de menor a
// mayor en O(n) un arreglo casi ordenado de mayor a menor.

package main

import "fmt"

func Invertir(array []int) {
	i, j := 0, len(array)-1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
}

func Inserción(array []int) []int {
	for i, valor := range array[1:] {
		j := i
		for j >= 0 && valor < array[j] {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = valor
	}
	return array
}

func main() {
	arrayMayor := []int{15, 11, 9, 10, 7, 4, 3}
	Invertir(arrayMayor)
	fmt.Println(arrayMayor)
	fmt.Println(Inserción(arrayMayor))
}
