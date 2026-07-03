// Implementar un algoritmo que, utilizando división y conquista, determine si un arreglo se encuentra ordenado. Indicar y justificar
// adecuadamente la complejidad del algoritmo implementado.

package main

func EstaOrdenado(arr []int) bool {
	return _ordenado(arr, 0, len(arr)-1)
}

func _ordenado(arr []int, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}
	mitad := (inicio + fin) / 2
	izq := _ordenado(arr, inicio, mitad)
	der := _ordenado(arr, mitad+1, fin)
	if arr[mitad] > arr[mitad+1] {
		return false
	}
	return izq && der
}
