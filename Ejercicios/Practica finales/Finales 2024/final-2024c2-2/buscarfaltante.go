// Tenemos un arreglo de números de 1 a n, ordenado. A dicho arreglo se le quita un elemento. Implementar un algoritmo
// que determine qué elemento falta en el arreglo. Indicar y justificar la complejidad del algoritmo implementado.

package main

func EncontrarElemento(arr []int) int {
	return _encontrar(arr, 0, len(arr)-1)
}

// A = 1, B = 2, C = 0, -> O(log n)
func _encontrar(arr []int, inicio, fin int) int {
	if inicio > fin {
		return inicio + 1
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == mitad+1 {
		return _encontrar(arr, mitad+1, fin)
	} else {
		return _encontrar(arr, inicio, mitad-1)
	}
}
