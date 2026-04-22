// En un diccionario todas las claves tienen que ser diferentes, no así sus valores.
// Escribir en Go una primitiva para el hash cerrado func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int
// que, sin usar el iterador interno, dado un hash devuelva la cantidad de valores diferentes que almacena.
// Se puede suponer que el tipo V (el genérico de los valores) en este caso también es comparable, como lo son las claves.
// Indicar y justificar el orden del algoritmo.

package main

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int {
	dictValores := CrearHash[V, int]()
	for _, celda := range dicc.tabla {
		if celda.estado != OCUPADO {
			continue
		}
		if !dictValores.Pertenece(celda.dato) {
			dictValores.Guardar(celda.dato, 1)
		}
	}
	return dictValores.Cantidad()
}
