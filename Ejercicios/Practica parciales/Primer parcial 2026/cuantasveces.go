// Implementar una función que reciba un slice de enteros ordenado y un valor K y devuelva cuántas veces aparece ese
// valor en el mismo. Indicar y justificar la complejidad del algoritmo implementado.

package main

func CuantasVeces(arr []int, k int) int {
	primeraPos := _primeraPosicion(arr, 0, len(arr)-1, k)
	ultimaPos := _ultimaPosicion(arr, 0, len(arr)-1, k)
	if primeraPos == -1 || ultimaPos == -1 {
		return 0
	}
	return ultimaPos - primeraPos + 1
}

func _primeraPosicion(arr []int, inicio, fin, buscado int) int {
	if inicio > fin {
		return -1
	}
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == buscado {
		if arr[mitad+1] == buscado {
			return _primeraPosicion(arr, inicio, mitad, buscado)
		} else {
			return mitad
		}
	} else if arr[mitad] > buscado {
		return _primeraPosicion(arr, inicio, mitad-1, buscado)
	} else {
		return _primeraPosicion(arr, mitad+1, fin, buscado)
	}
}

func _ultimaPosicion(arr []int, inicio, fin, buscado int) int {
	if inicio > fin {
		return -1
	}
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == buscado && arr[mitad+1] != buscado {
		return mitad
	} else if arr[mitad] > buscado {
		return _ultimaPosicion(arr, inicio, mitad-1, buscado)
	} else {
		return _ultimaPosicion(arr, mitad+1, fin, buscado)
	}
}
