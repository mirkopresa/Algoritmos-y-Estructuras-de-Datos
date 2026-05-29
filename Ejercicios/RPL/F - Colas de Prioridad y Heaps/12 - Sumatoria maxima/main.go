// Implementar un algoritmo que reciba un arreglo de n números, y un número k, y devuelva los k números
// dentro del arreglo cuya suma sería la máxima (entre todos los posibles subconjuntos de k elementos de dicho arreglo).
// La solución debe ser mejor que O(nlog⁡n) si k << n.
// Indicar y justificar la complejidad de la función implementada.

package main

func cmp(a, b int) int {
	return a - b
}

func Sumatoria(arreglo []int, k int) []int {
	heap := CrearHeapArr(arreglo, cmp) // O(n)
	resultado := make([]int, 0)
	// EL siguiente for es O(n log n) en el peor de los casos,
	// es decir, k >= n, pero en el caso de k << n, seria O(k log n) ya que estamos desencolando k veces
	for k > 0 && !heap.EstaVacia() {
		resultado = append(resultado, heap.Desencolar())
		k--
	}
	return resultado
}
