// Implementar, por división y conquista, una función que dado un arreglo
// sin elementos repetidos y casi ordenado (todos los elementos se encuentran ordenados, salvo uno),
// obtenga el elemento fuera de lugar.
// Indicar y justificar el orden.

package main

// A = 2, B = 2, C = 0, log en base 2 de 2 = 1 > C, entonces O(n^log(2 en base 2)) -> O(n)
func EncontrarElementoDesordenado(arr []int, inicio, fin int) int {
	if inicio == fin {
		return -1 // Esta ordenado
	}
	mitad := (inicio + fin) / 2
	// Caso ejemplos [1, 2, 5, 3, 4] o [1, 2, 3, -1, 4]
	if arr[mitad] > arr[mitad+1] {
		if mitad > 0 {
			if arr[mitad-1] < arr[mitad+1] {
				return arr[mitad]
			} else {
				return arr[mitad+1]
			}
		} else {
			return arr[mitad]
		}
	}
	mitadIzq := EncontrarElementoDesordenado(arr, inicio, mitad)
	if mitadIzq == -1 {
		return EncontrarElementoDesordenado(arr, mitad+1, fin)
	} else {
		return mitadIzq
	}
}

func ElementoDesordenado(arr []int) int {
	return EncontrarElementoDesordenado(arr, 0, len(arr)-1)
}
