// Implementar, por división y conquista, una función que determine el mínimo de un arreglo.
// Indicar y justificar el orden.

package main

// BuscarMinimo devuelve el valor del minimo del arreglo, no su posicion
// Precondicion: el arreglo tiene al menos un elemento

// A = 2, B = 2, C = 0
// log en base 2 de 2 = 1, 1 > C, -> O(n^(log en base 2 de 2)) -> O(n)

func MinimoRecursivo(arr []int, inicio int, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	mitad := (inicio + fin) / 2
	minimo_izq := MinimoRecursivo(arr, inicio, mitad)
	minimo_der := MinimoRecursivo(arr, mitad+1, fin)
	if minimo_izq > minimo_der {
		return minimo_der
	} else {
		return minimo_izq
	}
}

func BuscarMinimo(arr []int) int {
	return MinimoRecursivo(arr, 0, len(arr)-1)
}
