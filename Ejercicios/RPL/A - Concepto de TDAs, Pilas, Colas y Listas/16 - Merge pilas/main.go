// Dadas dos pilas de enteros positivos (con posibles valores repetidos)
// cuyos elementos fueron ingresados de menor a mayor, se pide implementar
// una función func MergePilas(pila1, pila2 Pila[int]) []int que devuelva un
// array ordenado de menor a mayor con todos los valores de ambas pilas sin repeticiones.
// Detallar y justificar la complejidad del algoritmo considerando que el tamaño de las pilas es N y M respectivamente.

package main

func cargarDesapiladosNoRepetidos(arr []int, pila Pila[int]) []int {
	desapilado := pila.Desapilar()
	if len(arr) == 0 || arr[len(arr)-1] != desapilado {
		arr = append(arr, desapilado)
	}
	return arr
}

func invertirArreglo(arr []int) []int {
	i, j := len(arr)-1, 0
	for i > j {
		arr[i], arr[j] = arr[j], arr[i]
		i--
		j++
	}
	return arr
}

func MergePilas(pila1, pila2 Pila[int]) []int {
	arr := make([]int, 0, 0)
	for !pila1.EstaVacia() || !pila2.EstaVacia() {
		if !pila1.EstaVacia() && !pila2.EstaVacia() && (pila1.VerTope() < pila2.VerTope()) {
			arr = cargarDesapiladosNoRepetidos(arr, pila2)
		} else if !pila1.EstaVacia() && !pila2.EstaVacia() && (pila1.VerTope() > pila2.VerTope()) {
			arr = cargarDesapiladosNoRepetidos(arr, pila1)
		} else if pila1.EstaVacia() {
			arr = cargarDesapiladosNoRepetidos(arr, pila2)
		} else {
			arr = cargarDesapiladosNoRepetidos(arr, pila1)
		}
	}
	arr = invertirArreglo(arr)
	return arr
}
