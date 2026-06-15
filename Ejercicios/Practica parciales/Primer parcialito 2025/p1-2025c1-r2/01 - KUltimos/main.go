// Implementar una primitiva KUltimos[T any](k int) []T para el TDA Cola que dada una cola devuelva un
// arreglo con los últimos k elementos que saldrían de la cola (los del fondo), en el orden en que saldrían de esta.
// En el caso de tener menos de k elementos encolados, devolver los existentes en la cola. Al finalizar la ejecución
// de la función la cola debe quedar en el mismo estado que antes de invocar a la primitiva.
// Indicar y justificar la complejidad del algoritmo.

package main

// O(n) siendo n los elementos de la cola, la vamos a terminar recorriendo siempre
func (c colaEnlazada[T]) KUltimos(k int) []T {
	rapido := c.primero
	for i := 0; i < k && rapido != nil; i++ {
		rapido = rapido.siguiente
	}
	res := make([]T, 0)
	lento := c.primero
	if rapido != nil {
		for rapido != nil {
			lento = lento.siguiente
			rapido = rapido.siguiente
		}
		for lento != nil {
			res = append(res, lento.dato)
			lento = lento.siguiente
		}
	} else {
		for lento != nil {
			res = append(res, lento.dato)
			lento = lento.siguiente
		}
	}
	return res
}
