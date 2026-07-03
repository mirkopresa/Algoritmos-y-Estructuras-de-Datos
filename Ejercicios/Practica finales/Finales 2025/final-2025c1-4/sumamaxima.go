// Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números dentro del
// arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo). Indicar y
// justificar la complejidad de la función implementada.

package main

func SumaMaxima(arr []int, k int) []int {
	heap := CrearHeapArr[int](arr, func(a, b int) int { return a - b })
	resultado := make([]int, 0)
	for k > 0 && !heap.EstaVacia() {
		resultado = append(resultado, heap.Desencolar())
		k--
	}
	return resultado
}
