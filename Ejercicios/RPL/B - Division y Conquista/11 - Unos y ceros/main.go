// Se tiene un arreglo tal que [1, 1, 1, …, 0, 0, …] (es decir, unos seguidos de ceros).
// Se pide una función de orden O(log(n)) que encuentre el índice del primer 0.
// Si no hay ningún 0 (solo hay unos), debe devolver -1.

package main

import "fmt"

// A = 1, B = 2, C = 0
// log en base 2 de 1 = 0, 0 = C  O(n^0 * log(n)) -> O(log(n))
func PrimerCeroRecursivo(arr []int, inicio, fin int) int {
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == 1 {
		return PrimerCeroRecursivo(arr, mitad+1, fin)
	} else {
		return PrimerCeroRecursivo(arr, inicio, mitad)
	}
}

func IndicePrimeroCero(arr []int) int {
	if arr[len(arr)-1] == 1 {
		return -1
	}
	return PrimerCeroRecursivo(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{1, 1, 1, 1, 1, 0}
	fmt.Println(IndicePrimeroCero(arr))
}
