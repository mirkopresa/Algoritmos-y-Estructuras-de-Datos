// Se tiene un arreglo ordenado ascendentemente el cual ha sufrido krotaciones (el cual es desconocido), y se quiere hallar
// el menor elemento del mismo. Implementar una función hallarMenor(array []int) int que lo retorne, utilizando
// División y Conquista. ¿Cuál es la complejidad del algoritmo? Justificar utilizando el Teorema Maestro.

package main

import "fmt"

func hallarMenor(array []int) int {
	return menorRecursivo(array, 0, len(array)-1)
}

// A = 1, B = 2, C = 0, log base 2 de 1 = 0, 0 = C, O(n^c*log(n)) -> O(log(n))
func menorRecursivo(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	mitad := (inicio + fin) / 2
	// Si es menor, lo incluimos en la proxima llamada
	// Si es mayor, no lo incluimos ya que estamos buscando el minimo
	if arr[mitad] <= arr[fin] {
		return menorRecursivo(arr, inicio, mitad)
	} else {
		return menorRecursivo(arr, mitad+1, fin)
	}
}

func main() {
	arr := []int{5, 1, 2, 3, 4}
	fmt.Println(hallarMenor(arr))
}
