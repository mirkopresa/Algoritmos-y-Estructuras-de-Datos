// Implementar una función mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] que, dado un arreglo
// de Alumnos y un valor entero k, nos devuelva una lista de los k alumnos de mayor promedio (ordenada de
// mayor a menor). Indicar y justificar la complejidad del algoritmo implementado.
// Considerar que la estructura del alumno es:

package main

type Alumno struct {
	nombre string
	padron int
	notas  []int
}

// Funcion de comparacion O(m + m) -> O(m) siendo m la cantidad de notas de los alumnos
func cmpPromedios(a, b Alumno) int {
	suma1, suma2 := 0, 0
	cant := 0
	// O(m)
	for _, nota := range a.notas {
		suma1 += nota
		cant++
	}
	promedio1 := suma1 / cant
	cant = 0
	// O(m)
	for _, nota := range b.notas {
		suma2 += nota
		cant++
	}
	promedio2 := suma2 / cant
	return promedio1 - promedio2
}

// Complejidad final -> O(N*M + K*(M log N)) -> O(M*(N + K log N)),
// N cantidad de alumnos, M el costo de la funcion de comparacion, K la cantidad de alumnos en la lista
func mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] {
	resultado := CrearListaEnlazada[Alumno]()
	heap := CrearHeapArr(alumnos, cmpPromedios) // O(N*M), N siendo la cantidad de alumnos y hacemos N veces comparaciones O(M)
	for k > 0 && !heap.EstaVacia() {            // Hacemos desencolar K veces, entonces nos queda O(k * (m log n))
		resultado.InsertarUltimo(heap.Desencolar()) // Desencolar cuesta O(m log n)
		k--
	}
	return resultado
}
