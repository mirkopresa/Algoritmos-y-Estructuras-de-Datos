// Implementar una primitiva KUltimos[T any](k int) []T para el TDA Cola que dada una cola devuelva un
// arreglo con los últimos k elementos que saldrían de la cola (los del fondo), en el orden en que saldrían de esta.
// En el caso de tener menos de k elementos encolados, devolver los existentes en la cola. Al finalizar la ejecución
// de la función la cola debe quedar en el mismo estado que antes de invocar a la primitiva.
// Indicar y justificar la complejidad del algoritmo.

package main

func (c colaEnlazada[T]) KUltimos(k int) []T {
	resultado := make([]T, 0)
	actual := c.primero
	contador := 0
	for actual != nil {
		contador++
		actual = actual.siguiente
	}
	if contador < k {
		nodo := c.primero
		for nodo != nil {
			resultado = append(resultado, nodo.dato)
			nodo = nodo.siguiente
		}
	} else {
		nodo = c.primero
		elementosAIgnorar := contador - k
		for elementosAIgnorar > 0 {
			nodo = nodo.siguiente
			elementosAIgnorar--
		}
		for nodo != nil {
			resultado = append(resultado, nodo.dato)
			nodo = nodo.siguiente
		}
	}
	return resultado
}

// El algoritmo tiene una complejidad de O(n) siendo n los elementos de la cola, ya que primero estamos
// contando todos los elementos en el primer for, luego, si hay menos que k, hacemos un for guardando
// todos los elementos de la cola (n)
// En el caso que no, el primer for nos skipea k-1 elementos y terminamos recorriendo el resto de la cola
// lo cual termina siendo n de nuevo, ya que estamos recorriendo toda la cola de nuevo si o si
// El resto de operaciones son constantes
