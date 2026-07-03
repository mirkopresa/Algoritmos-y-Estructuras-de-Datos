// Implementar la primitiva Invertir(igualdad func (V, V) bool) Diccionario[V, Lista[K]] tanto para el Hash Abierto como
// para el Hash Cerrado. Dicha primitiva debe devolver un nuevo hash que tenga como claves los valores del original, y como valores
// debe tener una lista con las claves que tuvieran dichos valores en el original. Si no se implementa para ambas implementaciones, el
// ejercicio no estará aprobable. Indicar y justificar la complejidad del algoritmo implementado.

package main

func (hash *hashCerrado[K, V]) Invertir() Diccionario[V, Lista[K]] {
	invertido := CrearHash[V, Lista[K]]()
	for _, celda := range hash.tabla {
		if celda.estado != OCUPADO {
			continue
		}
		if !invertido.Pertenece(celda.dato) {
			lista := CrearListaEnlazada[K]()
			lista.InsertarUltimo(celda.clave)
			invertido.Guardar(celda.dato, lista)
		} else {
			lista := invertido.Obtener(celda.dato)
			lista.InsertarUltimo(celda.clave)
		}
	}
	return invertido
}

func (hash *hashAbierto[K, V]) Invertir() Diccionario[V, Lista[K]] {
	invertido := CrearHash[V, Lista[K]]()
	for _, lista := range hash.tabla {
		for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
			par := iter.VerActual()
			if !invertido.Pertenece(par.dato) {
				lista2 := CrearListaEnlazada[K]()
				lista2.InsertarUltimo(par.clave)
				invertido.Guardar(par.dato, lista)
			} else {
				lista2 := invertido.Obtener(par.dato)
				lista2.InsertarUltimo(par.clave)
			}
		}
	}
	return invertido
}
