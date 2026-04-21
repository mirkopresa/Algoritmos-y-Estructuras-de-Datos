// Para un hash abierto, implementar una primitiva 
// func (hash *hashCerrado[K, V]) Claves() Lista[K] 
// que devuelva una lista con sus claves, sin utilizar el iterador interno.

package main

func (hash *hashAbierto[K, V]) Claves() Lista[K] {
	listaResultado := CrearListaEnlazada[K]()
	for _, celda := hash.tabla {
		for iter := celda.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			elemento := iter.VerActual()
			listaResultado.InsertarUltimo(elemento)
		}
	}
	for i := 0; i < len(hash.tabla); i++ {
		for iter := lista[i].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			elemento := iter.VerActual()
			listaResultado.InsertarUltimo(elemento)
		}
	}
    return listaResultado
}