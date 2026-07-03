// Dado un arreglo de números S, se define a un desordenamiento D a una permutación de los elementos tales que D[i] ̸= S[i]∀i (no
// hay ningún par de elementos que coincidan). El máximo desordenamiento es uno en el que D[i] > D[i + 1] para la mayor cantidad de
// posiciones i posibles. Implementar un algoritmo que permita obtener el máximo desordenamiento de un arreglo en O(n log n). Se
// puede asumir que todos los elementos del arreglo son diferentes. Justificar la complejidad del algoritmo implementado. Por ejemplo:
// [5, 4, 3, 2, 1] → [4, 5, 2, 1, 3]
// [1, 2, 3, 4, 5, 6] → [6, 5, 4, 3, 2, 1]
// [92, 3, 52, 13, 2, 31, 1] → [52, 92, 31, 3, 13, 1, 2]
// Recordar D es un desordenamiento, por lo que debe cumplirse que D[i] ̸= S[i]∀i. Ayuda: pensar en utilizar alguno de los TDAs
// implementados durante la cursada.

package main

// cuidado con los casos bordes
func MaxDesordenamiento(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	desordenamiento := make([]int, len(arr))
	heap := CrearHeapArr[int](arr, func(a, b int) int { return a - b })
	i := 0
	for !heap.EstaVacio() {
		desencolado := heap.Desencolar()
		if desencolado == arr[i] {
			if !heap.EstaVacio() {
				desencolado2 := heap.Desencolar()
				heap.Encolar(desencolado)
				desordenamiento[i] = desencolado2
			} else {
				desordenamiento[i] = desencolado
				desordenamiento[i], desordenamiento[i-1] = desordenamiento[i-1], desordenamiento[i]
			}
		} else {
			desordenamiento[i] = desencolado
		}
		i++
	}
	return desordenamiento
}
