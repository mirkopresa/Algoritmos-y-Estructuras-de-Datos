// Para un hash cerrado, implementar una primitiva
// func (hash *hashCerrado[K, V]) Claves() Lista[K]
// que devuelva una lista con sus claves, sin utilizar el iterador interno.

package main

func (hash *hashCerrado[K, V]) Claves() Lista[K] {
	listaResultado := CrearListaEnlazada[K]()
	for _, celda := range hash.tabla {
		if celda.estado != OCUPADO {
			continue
		}
		listaResultado.InsertarUltimo(celda.clave)
	}
	return listaResultado
}
