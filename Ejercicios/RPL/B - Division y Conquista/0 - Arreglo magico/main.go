// Implementar un algoritmo en Go que reciba un arreglo de enteros de tamaño nn,
// ordenado ascendentemente y sin elementos repetidos, y determine en O(log n) si es mágico.
// Un arreglo es mágico si existe algún valor i tal que 0 <= i y arr[i] = i.
// Justificar el orden del algoritmo.

package main

// A = 1, B = 2, C = 0
// 1 x T(N/2) + O(n⁰)
// log en base 2 de 1 = C, entonces = O(n⁰log(n)) = o(log(n))

func MagicoRecursivo(arr []int, inicio int, fin int) bool {
	if inicio > fin {
		return false
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == mitad {
		return true
	} else if arr[mitad] > mitad {
		return MagicoRecursivo(arr, inicio, mitad-1)
	} else {
		return MagicoRecursivo(arr, mitad+1, fin)
	}
}

func ArregloEsMagico(arr []int) bool {
	return MagicoRecursivo(arr, 0, len(arr)-1)
}
