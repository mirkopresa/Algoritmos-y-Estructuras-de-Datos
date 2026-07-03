// Implementar una primitiva para el hash cerrado filtro(func(V) bool) que elimine del hash todas las claves del mismo que
// tengan asociado un valor para el cuál la función devuelva false. Indicar y justificar la complejidad de la primitiva implementada.

package main

func (h *hashCerrado[K, V]) Filtro(f func(V) bool) {
	// ojo, acceder a traves del indice para modificar correctamente en memoria
	for i, _ := range h.tabla {
		if h.tabla[i].estado == OCUPADO && !f(h.tabla[i].dato) {
			h.tabla[i].estado == BORRADO
			h.cantidad--
			h.borrados++
		}
	}
	factor := float64(h.cantidad+h.borrados) / float64(h.tam)
	if factor <= 0.35 {
		h.redimensionar(h.tam / 2)
	}
}
