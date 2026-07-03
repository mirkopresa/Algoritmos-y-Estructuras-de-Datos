// En un colegio de Villa General Belgrano, los alumnos deben formar de la siguiente manera: primero las niñas ordenadas por altura
// de menor a mayor. Luego los niños, también ordenados por altura de menor a mayor. Escribir una función en Go que reciba una
// cola de enteros (representando las alturas de los alumnos en centímetros), y devuelva true si es posible que estén bien formados
// (considerando sólo las alturas) o false en caso contrario. Se puede vaciar la cola sin necesidad de dejarla como se la recibió. Indicar
// y justificar la complejidad de la función implementada. Se dejan ejemplos de ejecución al dorso del examen.

// - Primero -> [ 125, 128, 129, 124, 134, 138, 140 ]: true
// - Primero -> [ 125, 128, 129, 133, 134, 138, 140 ]: true
// - Primero -> [ 125, 120, 129, 133, 124, 138, 140 ]: false

package main

// O(n), vemos todos los elementos de la cola y hacemos operaciones O(1)
func EstanFormados(formacion Cola[int]) bool {
	cantidad := 0
	for !formacion.EstaVacia() {
		desencolado := formacion.Desencolar()
		if !formacion.EstaVacia() {
			if desencolado > formacion.VerPrimero() {
				cantidad++
			}
		}
	}
	if cantidad >= 2 {
		return false
	}
	return true
}
