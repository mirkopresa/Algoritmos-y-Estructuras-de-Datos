// Implementar una función (que utilice división y conquista) de orden O(n logn) que dado un arreglo de n números enteros
// devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces.
// Justificar el orden de la solución.
// Ejemplos:
// [1, 2, 1, 2, 3] -> false
// [1, 1, 2, 3] -> false
// [1, 2, 3, 1, 1, 1] -> true
// [1] -> true

package main

func MasDeLaMitad(arr []int) bool {
	_, masDe := _mitad(arr)
	return masDe
}

func _mitad(arr []int) (int, bool) {
	if len(arr) <= 1 {
		return arr[0], true
	}
	mitad := len(arr) / 2
	num1, mitadIzq := _mitad(arr[:mitad])
	num2, mitadDer := _mitad(arr[mitad:])
	if num1 == num2 {
		return num1, mitadIzq && mitadDer
	} else {
		return contar(num1, num2, arr)
	}
}

func contar(num1, num2 int, arr []int) (int, bool) {
	cantidad1, cantidad2 := 0, 0
	for _, elem := range arr {
		if elem == num1 {
			cantidad1++
		}
		if elem == num2 {
			cantidad2++
		}
	}
	mitad := len(arr) / 2
	if cantidad1 > mitad {
		return num1, true
	} else if cantidad2 > mitad {
		return num2, true
	} else {
		return -1, false
	}
}
