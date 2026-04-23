// Implementar una función que, dado un arreglo ordenado y sin elementos repetidos de valores enteros no negativos,
// obtenga el mínimo valor que no se encuentre en el arreglo. Indicar y justificar adecuadamente la complejidad del
// algoritmo.
// Por ejemplo:
// minimoExcluido([0, 1, 5]) --> 2
// minimoExcluido([1, 3, 5]) --> 0
// minimoExcluido([0, 1, 2, 3, 4, 5]) --> 6
// minimoExcluido([0, 1, 2, 3, 4, 5, 1234567]) --> 6

package main

func minimoExcluido(arr []int) int {
	return minimoRecursivo(arr, 0, len(arr)-1)
}

func minimoRecursivo(arr []int, inicio, fin int) int {
	if inicio > fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == mitad {
		return minimoRecursivo(arr, mitad+1, fin)
	}
	return minimoRecursivo(arr, inicio, mitad-1)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 1234567}
	println(minimoExcluido(arr))
	arr2 := []int{0, 1, 2, 3, 4, 5}
	println(minimoExcluido(arr2))
	arr3 := []int{1, 3, 5}
	println(minimoExcluido(arr3))
	arr4 := []int{0, 1, 5}
	println(minimoExcluido(arr4))
}
