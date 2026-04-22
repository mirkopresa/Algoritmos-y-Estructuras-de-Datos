// Implementar un algoritmo que, por División y Conquista, determine si un arreglo se encuentra
// ordenado ascendentenemente, o no. Indicar y justificar la complejidad del algoritmo implementado

package main

// A = 2, B = 2, C = 0; log en base 2 de 2 = 1 > C, O(n^(log base 2 de 2)) -> O(n)
func EstaOrdenadoRecursivo(arr []int, inicio int, fin int) bool {
	if inicio >= fin {
		return true
	}
	mitad := (inicio + fin) / 2
	mitad_izq := EstaOrdenadoRecursivo(arr, inicio, mitad)
	mitad_der := EstaOrdenadoRecursivo(arr, mitad+1, fin)
	if arr[mitad] <= arr[mitad+1] && mitad_izq && mitad_der {
		return true
	}
	return false
}
