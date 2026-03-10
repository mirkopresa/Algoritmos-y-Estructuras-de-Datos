// Implementar una función en Go que determine la posición en la que se encuentra un determinado elemento en un arreglo (ambos pasados por parámetro).
// En caso de no encontrarse el elemento en el arreglo, devolver -1.

package main

import "fmt"

func BuscarElem(arr []int, elem int) int {
	// Otra manera de hacerlo
	// for index, value := range arr {
	// 	if elem == value {
	//		return index
	// 	}
	// }

	for i := 0; i < len(arr); i++ {
		if elem == arr[i] {
			return i
		}
	}
	return -1
}

// Main de prueba
func main() {
	array := []int{20, 50, 4, 77}
	fmt.Println(BuscarElem(array, 4))
}
