// Implementar la primitiva Interseccion(otro *abb[K, V]) Lista[K] para el ABB que nos devuelva una lista ordenada con la
// intersección entre el árbol y el recibido por parámetro, que estén ocupando el mismo lugar en el árbol. Indicar y justificar la
// complejidad del algoritmo implementado. En el ejemplo a continuación, la intersección sería [4, 10, 18, 20].

package main

func (abb *abb[K, V]) Interseccion(otro *abb[K, V]) Lista[K] {
	resultado := CrearListaEnlazada[K]()
	abb._interseccion(otro, resultado)
	return resultado
}

// O(n) si n < m, O(m) si n > m
func (abb *abb[K, V]) _interseccion(otro *abb[K, V], lista Lista[K]) {
	if abb == nil || otro == nil {
		return
	}
	abb.izq._interseccion(otro.izq, lista)
	if abb.clave == otro.clave {
		lista.InsertarUltimo(abb.clave)
	}
	abb.der._interseccion(otro.der, lista)
}
