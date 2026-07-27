// Practica division y conquista para primer parcial/final

package main

import "fmt"

// Se cuenta con un arreglo de enteros ordenado de manera ascendente que contiene exactamente un número duplicado
// (es decir, todos los demás elementos son distintos, sin duplicados).
// Implementar una función que encuentre dicho número utilizando división y conquista.
// Indicar y justificar la complejidad del algoritmo, utilizando el Teorema Maestro.

func ElementoDuplicado(arr []int) int {
	return _duplicado(arr, 0, len(arr)-1)
}

func _duplicado(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return -1
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == arr[mitad+1] {
		return arr[mitad+1]
	}
	mitadIzq := _duplicado(arr, inicio, mitad)
	if mitadIzq == -1 {
		return _duplicado(arr, mitad+1, fin)
	} else {
		return mitadIzq
	}
}

// Se tiene un arreglo tal que [1, 1, 1, …, 0, 0, …] (es decir, unos seguidos de ceros).
// Se pide una función de orden O(log(n)) que encuentre el índice del primer 0.
// Si no hay ningún 0 (solo hay unos), debe devolver -1.

func IndicePrimeroCero(arr []int) int {
	return _primerCero(arr, 0, len(arr)-1)
}

func _primerCero(arr []int, inicio, fin int) int {
	if inicio == fin {
		if arr[inicio] == 0 {
			return inicio
		} else {
			return -1
		}
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == 1 {
		return _primerCero(arr, mitad+1, fin)
	} else {
		return _primerCero(arr, inicio, mitad)
	}
}

// Implementar un algoritmo que, por división y conquista, permita obtener la parte entera de la raíz cuadrada de un número n,
// en tiempo O(log n). Por ejemplo, para n = 10 debe devolver 3, y para n = 25 debe devolver 5.

func ParteEnteraRaiz(n int) int {
	return _raiz(n, 0, n)
}

func _raiz(numero, min, max int) int {
	if min > max {
		return max
	}
	mitad := (min + max) / 2
	if mitad*mitad == numero {
		return mitad
	} else if mitad*mitad > numero {
		return _raiz(numero, min, mitad-1)
	} else {
		return _raiz(numero, mitad+1, max)
	}
}

// Se tiene un arreglo de N >= 3 elementos en forma de pico, esto es: estrictamente creciente hasta una determinada posición p,
// y estrictamente decreciente a partir de ella (con 0 < p < N - 1).
// Por ejemplo, en el arreglo [1, 2, 3, 1, 0, -2] la posición del pico es p = 2.
// Se pide:
// Implementar un algoritmo de división y conquista de orden O(log n) que encuentre la posición p del pico:
// func PosicionPico(v []int, ini, fin int) int. La función será invocada inicialmente como: PosicionPico(v, 0, len(v)-1), y
// tiene como pre-condición que el arreglo tenga forma de pico.

func PosicionPico(arr []int) int {
	return _posicionPico(arr, 0, len(arr)-1)
}

func _posicionPico(arr []int, ini, fin int) int {
	if ini == fin {
		return ini
	}
	mitad := (ini + fin) / 2
	if arr[mitad] > arr[mitad+1] {
		return _posicionPico(arr, ini, mitad)
	} else {
		return _posicionPico(arr, mitad+1, fin)
	}
}

// Implementar una función que reciba un slice de enteros ordenado y un valor K y devuelva cuántas veces aparece ese
// valor en el mismo. Indicar y justificar la complejidad del algoritmo implementado.

func CuantasVeces(arr []int, k int) int {
	primeraPos := _primeraPosicion(arr, 0, len(arr)-1, k)
	if primeraPos == -1 {
		return 0
	}
	ultimaPos := _ultimaPosicion(arr, 0, len(arr)-1, k)
	return ultimaPos - primeraPos + 1
}

func _primeraPosicion(arr []int, inicio, fin, buscado int) int {
	if inicio > fin {
		return -1
	}
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == buscado {
		if mitad > 0 && arr[mitad-1] == buscado {
			return _primeraPosicion(arr, inicio, mitad-1, buscado)
		} else {
			return mitad
		}
	} else if arr[mitad] > buscado {
		return _primeraPosicion(arr, inicio, mitad-1, buscado)
	} else {
		return _primeraPosicion(arr, mitad+1, fin, buscado)
	}
}

func _ultimaPosicion(arr []int, inicio, fin, buscado int) int {
	if inicio == fin {
		return inicio
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == buscado && arr[mitad+1] != buscado {
		return mitad
	} else if arr[mitad] > buscado {
		return _ultimaPosicion(arr, inicio, mitad-1, buscado)
	} else {
		return _ultimaPosicion(arr, mitad+1, fin, buscado)
	}
}

// Implementar una función func esCuadradoPerfecto(n int) bool que por División y Conquista determine si el
// número n (un positivo entero) es un cuadrado perfecto. Un número es cuadrado perfecto si existe un número entero x
// tal que x^2 = n. Indicar y justificar la complejidad del algoritmo utilizando el Teorema Maestro.

func esCuadradoPerfecto(n int) bool {
	return _cuadradoPerfecto(n, 0, n)
}

func _cuadradoPerfecto(numero, min, max int) bool {
	if min > max {
		return false
	}
	mitad := (min + max) / 2
	if mitad*mitad == numero {
		return true
	} else if mitad*mitad > numero {
		return _cuadradoPerfecto(numero, min, mitad-1)
	} else {
		return _cuadradoPerfecto(numero, mitad+1, max)
	}
}

// Se tiene un arreglo ordenado ascendentemente el cual ha sufrido k rotaciones (el cual es desconocido), y se quiere hallar
// el menor elemento del mismo. Implementar una función hallarMenor(array []int) int que lo retorne, utilizando
// División y Conquista. ¿Cuál es la complejidad del algoritmo? Justificar utilizando el Teorema Maestro.

func hallarMenor(array []int) int {
	return _menor(array, 0, len(array)-1)
}

// Ej: [1,2,3,4,5,6] rotado 3 veces seria [4,5,6,7,1,2,3]
func _menor(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] > arr[fin] {
		return _menor(arr, mitad+1, fin)
	} else {
		return _menor(arr, inicio, mitad)
	}

}

// Implementar un algoritmo que reciba un arreglo de enteros desordenado y un número elem que, por división y
// conquista determine si elem se encuentra en el arreglo. Indicar y justificar adecuadamente la complejidad del algoritmo
// implementado.

func encontrarElem(arr []int, elem int) bool {
	return _encontrar(arr, 0, len(arr)-1, elem)
}

// [5, 3, 8, 1, 2], elem = 1
func _encontrar(arr []int, inicio, fin, elem int) bool {
	if inicio == fin {
		return arr[inicio] == elem
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] == elem {
		return true
	}
	mitadIzq := _encontrar(arr, inicio, mitad, elem)
	if !mitadIzq {
		return _encontrar(arr, mitad+1, fin, elem)
	} else {
		return mitadIzq
	}
}

// Implementar un algoritmo que, por División y Conquista, determine si un arreglo se encuentra
// ordenado ascendentenemente, o no. Indicar y justificar la complejidad del algoritmo implementado.

func EstaOrdenado(arr []int) bool {
	return _ordenado(arr, 0, len(arr)-1)
}

func _ordenado(arr []int, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}
	mitad := (inicio + fin) / 2
	mitadIzq := _ordenado(arr, inicio, mitad)
	mitadDer := _ordenado(arr, mitad+1, fin)
	if arr[mitad] > arr[mitad+1] {
		return false
	} else {
		return mitadIzq && mitadDer
	}
}

// Implementar una función que, dado un arreglo ordenado y sin elementos repetidos de valores enteros no negativos,
// obtenga el mínimo valor que no se encuentre en el arreglo. Indicar y justificar adecuadamente la complejidad del
// algoritmo.

func MinimoExcluido(arr []int) int {
	return _minimo(arr, 0, len(arr)-1)
}

func _minimo(arr []int, inicio, fin int) int {
	if inicio == fin {
		if arr[inicio] == inicio {
			return inicio + 1
		} else {
			return inicio
		}
	}
	mitad := (inicio + fin) / 2
	if arr[mitad] != mitad {
		return _minimo(arr, inicio, mitad)
	} else {
		return _minimo(arr, mitad+1, fin)
	}
}

func main() {
	// caso1 := []int{1, 2, 3, 4, 4, 5, 6}
	// // Caso 2: El duplicado está en la primera posición (borde izquierdo)
	// caso2 := []int{7, 7, 8, 9, 10, 11, 12}
	// // Caso 3: El duplicado está en la última posición (borde derecho)
	// caso3 := []int{20, 21, 22, 23, 24, 25, 25}
	// // Caso 4: Arreglo más largo con inicio negativo
	// caso4 := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 5, 6, 7, 8}
	// // Caso 5: El arreglo más pequeño posible (mínimo 2 elementos para que haya un duplicado)
	// caso5 := []int{42, 42}
	// fmt.Println(ElementoDuplicado(caso1))
	// fmt.Println(ElementoDuplicado(caso2))
	// fmt.Println(ElementoDuplicado(caso3))
	// fmt.Println(ElementoDuplicado(caso4))
	// fmt.Println(ElementoDuplicado(caso5))

	// // Caso 1: Caso normal (el primer cero está en el medio)
	// caso1 := []int{1, 1, 1, 0, 0, 0}
	// // Caso 2: Solo hay unos (no hay ceros, debe devolver -1)
	// caso2 := []int{1, 1, 1, 1, 1}
	// // Caso 3: Solo hay ceros (el primer cero está en el índice 0)
	// caso3 := []int{0, 0, 0, 0}
	// // Caso 4: Un solo cero al final (borde derecho)
	// caso4 := []int{1, 1, 1, 1, 0}
	// // Caso 5: Un solo uno al principio (borde izquierdo)
	// caso5 := []int{1, 0, 0, 0, 0}
	// fmt.Println(IndicePrimeroCero(caso1)) // Debe devolver 3
	// fmt.Println(IndicePrimeroCero(caso2)) // Debe devolver -1
	// fmt.Println(IndicePrimeroCero(caso3)) // Debe devolver 0
	// fmt.Println(IndicePrimeroCero(caso4)) // Debe devolver 4
	// fmt.Println(IndicePrimeroCero(caso5)) // Debe devolver 1
	// fmt.Println(ParteEnteraRaiz(25))
	// fmt.Println(ParteEnteraRaiz(10))
	// // Caso 1: El ejemplo del enunciado (pico en el medio)
	// caso1 := []int{1, 2, 3, 1, 0, -2}
	// // Caso 2: Pico muy a la izquierda (índice 1)
	// caso2 := []int{10, 20, 15, 10, 5, 0}
	// // Caso 3: Pico muy a la derecha (índice len-2)
	// caso3 := []int{2, 4, 6, 8, 10, 9}
	// // Caso 4: Arreglo de tamaño mínimo (N=3)
	// caso4 := []int{1, 5, 2}
	// // Caso 5: Subida abrupta, bajada lenta
	// caso5 := []int{-5, 100, 99, 98, 97, 96}
	// // Caso 6: Cruzando el cero (negativos a positivos y vuelta)
	// caso6 := []int{-20, -10, 0, 5, 15, 8, -1}
	// fmt.Println(PosicionPico(caso1)) // Debe devolver 2
	// fmt.Println(PosicionPico(caso2)) // Debe devolver 1
	// fmt.Println(PosicionPico(caso3)) // Debe devolver 4
	// fmt.Println(PosicionPico(caso4)) // Debe devolver 1
	// fmt.Println(PosicionPico(caso5)) // Debe devolver 1
	// fmt.Println(PosicionPico(caso6)) // Debe devolver 4
	// Caso 1: K aparece varias veces en el medio
	// caso1 := []int{1, 2, 4, 4, 4, 5, 6}
	// k1 := 4
	// // Resultado esperado: 3

	// // Caso 2: K NO está en el arreglo
	// caso2 := []int{1, 2, 3, 5, 6}
	// k2 := 4
	// // Resultado esperado: 0

	// // Caso 3: K está en el borde izquierdo
	// caso3 := []int{7, 7, 8, 9, 10}
	// k3 := 7
	// // Resultado esperado: 2

	// // Caso 4: K está en el borde derecho
	// caso4 := []int{1, 2, 3, 9, 9, 9}
	// k4 := 9
	// // Resultado esperado: 3

	// // Caso 5: El arreglo está compuesto ÚNICAMENTE por K
	// caso5 := []int{5, 5, 5, 5, 5}
	// k5 := 5
	// // Resultado esperado: 5

	// // Caso 6: Arreglo de un solo elemento (y es K)
	// caso6 := []int{8}
	// k6 := 8
	// // Resultado esperado: 1

	// // Caso 7: Arreglo vacío
	// caso7 := []int{}
	// k7 := 1

	// caso8 := []int{1, 4, 5}
	// k8 := 4

	// fmt.Println(CuantasVeces(caso1, k1)) // Debe devolver 3
	// fmt.Println(CuantasVeces(caso2, k2)) // Debe devolver 0
	// fmt.Println(CuantasVeces(caso3, k3)) // Debe devolver 2
	// fmt.Println(CuantasVeces(caso4, k4)) // Debe devolver 3
	// fmt.Println(CuantasVeces(caso5, k5)) // Debe devolver 5
	// fmt.Println(CuantasVeces(caso6, k6)) // Debe devolver 1
	// fmt.Println(CuantasVeces(caso7, k7)) // Debe devolver 0
	// fmt.Println(CuantasVeces(caso8, k8)) // Debe devolver 1
	// fmt.Println(esCuadradoPerfecto(0))
	// fmt.Println(hallarMenor([]int{4, 5, 6, 7, 1, 2, 3}))
	// arrDesordenado := []int{5, 3, 8, 1, 2, 10}
	// fmt.Println(encontrarElem(arrDesordenado, 8))

	// arrOrdenado := []int{1, 2, 3, 4, 5}
	// fmt.Println(EstaOrdenado(arrOrdenado)) // Debe devolver true

	// arrDesordenado := []int{5, 3, 8, 1, 2}
	// fmt.Println(EstaOrdenado(arrDesordenado)) // Debe devolver false

	// arrUnElemento := []int{42}
	// fmt.Println(EstaOrdenado(arrUnElemento)) // Debe devolver true

	arrMinimoExcluido := []int{0, 5, 1}
	fmt.Println(MinimoExcluido(arrMinimoExcluido)) // Debe devolver 4
	arrMinimoExcluido2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(MinimoExcluido(arrMinimoExcluido2)) // Debe devolver 6
	arrMinimoExcluido3 := []int{1, 2, 3, 4, 5}
	fmt.Println(MinimoExcluido(arrMinimoExcluido3)) // Debe devolver 0
	arrMinimoExcluido4 := []int{0, 1, 2, 3, 4}
	fmt.Println(MinimoExcluido(arrMinimoExcluido4)) // Debe devolver 5
	arrMinimoExcluido5 := []int{0, 2, 3, 4, 5}
	fmt.Println(MinimoExcluido(arrMinimoExcluido5)) // Debe devolver 1
}
