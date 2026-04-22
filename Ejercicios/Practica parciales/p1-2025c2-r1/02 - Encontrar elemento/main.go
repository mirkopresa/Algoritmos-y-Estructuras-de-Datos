// Implementar un algoritmo que reciba un arreglo de enteros desordenado y un número elem que, por división y
// conquista determine si elem se encuentra en el arreglo.
// Indicar y justificar adecuadamente la complejidad del algoritmo implementado.

package main

import "fmt"

// A = 2, B = 2, C = 0, log en base 2 de 2 = 1 > C -> O(n^(log en base 2 de 2)) -> O(n)
func encontrarElementoEnDesordenado(arr []int, inicio, fin, elem int) bool {
	if inicio > fin {
		return false
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == elem {
		return true
	}
	mitadIzq := encontrarElementoEnDesordenado(arr, inicio, mitad-1, elem)
	mitadDer := encontrarElementoEnDesordenado(arr, mitad+1, fin, elem)
	return mitadIzq || mitadDer
}

func main() {
	arr := []int{5, 3, 10, 1, 4, 20, 40, 1, 4, 333, 2, 1}
	for _, elem := range arr {
		fmt.Println(encontrarElementoEnDesordenado(arr, 0, len(arr)-1, elem))
	}
}
