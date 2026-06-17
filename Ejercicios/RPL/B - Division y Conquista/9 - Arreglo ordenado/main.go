// Implementar, por división y conquista, una función que dado un arreglo y su largo,
// determine si el mismo se encuentra ordenado.
// Indicar y justificar el orden.

package main

import "fmt"

// A = 2, B = 2, C = 0; log en base 2 de 2 = 1 > C, O(n^(log base 2 de 2)) -> O(n)
func EstaOrdenadoRecursivo(arr []int, inicio int, fin int) bool {
	if inicio >= fin {
		return true
	}
	mitad := (inicio + fin) / 2
	mitadIzq := EstaOrdenadoRecursivo(arr, inicio, mitad)
	mitadDer := EstaOrdenadoRecursivo(arr, mitad+1, fin)
	if arr[mitad] > arr[mitad+1] {
		return false
	} else {
		return mitadIzq && mitadDer
	}
}

func EstaOrdenado(arr []int) bool {
	return EstaOrdenadoRecursivo(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{3, 4}
	fmt.Println(EstaOrdenado(arr))
}
