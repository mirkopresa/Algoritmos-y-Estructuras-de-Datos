// Dado un arreglo de enteros y un número K, se desea que todos los elementos
// del arreglo sean mayores a K. Aquellos números que sean menores o iguales a K
// deberían combinarse de la siguiente forma: buscar los dos números más chicos del
// vector, sacarlos y generar uno nuevo de la forma
// Nuevo número = número-más-chico + 2 x segundo-más-chico. Por ejemplo, si K = 10 y
// los números más chicos del arreglo son 3 y 4: 3 + 2 * 4 = 11. Los números combinados
// pueden volver a ser combinados con otros en caso de ser necesario (en el ejemplo
// no lo es), aplicando la misma lógica hasta que el número resultante sea mayor a K.

// Implementar una función que reciba un arreglo de enteros, su largo y un número `K`,
// y devuelva una lista con los elementos que quedarían luego de aplicar las
// modificaciones. El arreglo original debe quedar en el estado original. El orden de
// la lista no es importante. En caso de no poderse combinar todos los elementos para
// que todos los elementos sea mayores a `K`, devolver `nil`. Determinar y justificar
// la complejidad del algoritmo implementado.

// Ejemplo: Entrada: `[11, 14, 8, 19, 42, 3, 1, 9]; K = 13`:

// [11, 14, 8, 19, 42, 3, 1, 9] - (1,3)  -> 1 + 2*3 = 7
// [11, 14, 8, 19, 42, 9, 7]    - (7,8)  -> 7 + 8*2 = 23
// [11, 14, 19, 42, 9, 23]      - (9,11) -> 9 + 11*2 = 31
// [14, 19, 42, 23, 31]

package main

func Combinar(arr []int, k int) []int {
	heap := CrearHeapArr(arr, func(a, b int) int { return b - a })
	resultado := make([]int, 0)
	for !heap.EstaVacia() && heap.VerMax() <= k {
		primerNumero := heap.Desencolar()
		var segundoNumero int
		if !heap.EstaVacia() && heap.VerMax() <= k {
			segundoNumero = heap.Desencolar()
		} else {
			return nil
		}
		numeroNuevo := primerNumero + 2*segundoNumero
		heap.Encolar(numeroNuevo)
	}
	for !heap.EstaVacia() {
		resultado = append(resultado, heap.Desencolar())
	}
	return resultado
}
