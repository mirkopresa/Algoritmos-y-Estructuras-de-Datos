// Se cuenta con un arreglo de enteros ordenado de manera ascendente que contiene exactamente un número
// duplicado (es decir, todos los demás elementos son distintos). Implementar una función que encuentre dicho
// número utilizando División y Conquista. Indicar y justificar la complejidad del algoritmo,
// utilizando el Teorema Maestro.

package main

import "fmt"

// A = 2, B = 2, C = 0, log en base 2 de 2 = 1 > C, O(n^c) -> O(n)
func EncontrarDuplicado(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return -1
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == arr[mitad+1] || (mitad > 0 && arr[mitad-1] == arr[mitad]) {
		return arr[mitad]
	}
	mitadIzq := EncontrarDuplicado(arr, inicio, mitad-1)
	if mitadIzq == -1 {
		return EncontrarDuplicado(arr, mitad+1, fin)
	} else {
		return mitadIzq
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 6}
	arrFinal := []int{1, 20, 40, 40}
	arrPrimero := []int{47, 47, 60, 120, 160}
	fmt.Println(EncontrarDuplicado(arr, 0, len(arr)-1))
	fmt.Println(EncontrarDuplicado(arrFinal, 0, len(arr)-1))
	fmt.Println(EncontrarDuplicado(arrPrimero, 0, len(arr)-1))
}
