// Se tienen k arreglos de enteros previamente ordenados y se quiere obtener
// un arreglo ordenado que contenga a todos los elementos de los k arreglos. Sabiendo que cada arreglo
// tiene tamaño h, definimos como n a la sumatoria de la cantidad de elementos de todos los arreglos,
// es decir, n = k \times h.

// Escribir en Go una función func KMerge(arr [][]int) que reciba los k arreglos y
// devuelva uno nuevo con los n elementos ordenados entre sí. La función debe ser de orden
// O(n log k). Justificar el orden del algoritmo propuesto.

package main

func comparacionPrimerElemento(arr1, arr2 []int) int {
	if arr1[0] < arr2[0] {
		return 1
	}
	if arr1[0] > arr2[0] {
		return -1
	}
	return 0
}

func KMerge(arr [][]int) []int {
	resultado := make([]int, 0)
	heap := CrearHeapArr(arr, comparacionPrimerElemento) // O(k) (creamos un heap de k arreglos)
	for !heap.EstaVacia() {                              // O(n) siendo n la cantidad de elementos de todos los arreglos
		arreglo := heap.Desencolar()              // O(log k) siendo k la cantidad de arreglos
		resultado = append(resultado, arreglo[0]) // O(1)
		if len(arreglo[1:]) > 0 {                 // O(1)
			heap.Encolar(arreglo[1:]) // O(log k) siendo k la cantidad de arreglos
		}
	}
	// Complejidad final O(k + n * ((log k) + (log k))) -> O(k + n*logk) -> O(n*logk) (el segundo termino domina en el crecimiento)
	return resultado
}
