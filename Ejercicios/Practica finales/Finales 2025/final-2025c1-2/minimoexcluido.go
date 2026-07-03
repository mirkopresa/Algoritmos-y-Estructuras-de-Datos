// Implementar una función func minimoExcluido(arr []int) int que dado un arreglo de valores enteros (mayores o iguales a 0),
// obtenga el mínimo valor que no se encuentre en el arreglo. Indicar y justificar la complejidad del algoritmo (explicar en detalle este
// paso, porque es fácil que se te puedan pasar detalles importantes a explicar). ¿Es el mismo ejercicio del parcialito? Si.
// Por ejemplo:
// minimoExcluido([]int{0, 5, 1}) --> 2
// minimoExcluido([]int{3, 5, 1}) --> 0
// minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2}) --> 6
// minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2, 12345675433221345}) --> 6

package main

// [0, 1, 5]
func minimoExcluido(arr []int) int {
	// seria ordenar con merge sort, no lo voy a escribir completo
	ordenado := MergeSort(arr)
	k := 0
	for i := 0; i < len(ordenado)-1; i++ {
		if ordenado[i] == k {
			k++
		} else if ordenado[i] < k {
			continue
		} else {
			break
		}
	}
	return k
}
