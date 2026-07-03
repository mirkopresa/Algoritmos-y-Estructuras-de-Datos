// Se tiene un arreglo de enteros de tamaño conocido n, inicialmente ordenado de menor a mayor y sin elementos repetidos. Al mismo
// se le aplica un corrimiento de una cierta cantidad k de elementos. Esto es, todos los elementos que están a partir del índice k se
// los correrá hacia la izquierda k posiciones. Los elementos entre las posiciones 0 y k − 1, estarán ahora al final del arreglo (cada
// sub-segmento mantendrá el orden original de los elementos). Por ejemplo: v = [0, 1, 3, 5, 7, 8, 9] luego de correrlo con k = 3
// resulta en v = [5, 7, 8, 9, 0, 1, 3].
// [3, 5, 7, 8, 9, 0, 1]
// Implementar en Go una función que devuelva el valor de k (1 ≤ k ≤ n − 1) para un arreglo ya corrido en dicha cantidad k
// desconocida, utilizando un algoritmo de complejidad O(log(n)).
// La firma de la función es: func buscarK(v int[], ini int, fin int) int, y será llamada inicialmente con: buscarK(v, 0,
// n-1). Justificar la complejidad del algoritmo propuesto.

package main

import "fmt"

// Busca el minimo del arreglo y resta la longitud del arreglo con la posicion
// A = 1, B = 2, C = 0 -> log(2)(1) = 0, 0 = C -> O(n^c * log n) -> O(log n)
func buscarK(v []int, inicio, fin int) int {
	if inicio == fin {
		return len(v) - inicio
	}
	mitad := (inicio + fin) / 2
	if v[mitad] > v[fin] {
		return buscarK(v, mitad+1, fin)
	} else {
		return buscarK(v, inicio, mitad)
	}
}

func main() {
	fmt.Println(buscarK([]int{2, 3, 4, 5, 6, 7, 1}, 0, 6))
}
