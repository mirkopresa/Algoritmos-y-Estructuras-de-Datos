// Implementar un algoritmo que reciba un arreglo de elementos genéricos y
// una función de comparación (*) de dicho tipo, y devuelva el elemento más grande dentro de dicho arreglo.
// Se puede asumir que el arreglo tiene al menos un elemento.

package main

func Maximo[T any](arr []T, comparacion func(T, T) int) T {
	max := arr[0]
	for _, value := range arr[1:] {
		if comparacion(max, value) < 0 {
			max = value
		}
	}
	return max
}
