// Escribir una primitiva para la pila (dinámica) cuya firma es
// func (pila pilaDinamica[T]) Transformar(aplicar func(T) T) Pila[T]
// que devuelva una nueva pila cuyos elementos sean los resultantes de
// aplicarle la función aplicar a cada elemento de la pila original.
// Los elementos en la nueva pila deben tener el orden que tenían en la pila original,
// y la pila original debe quedar en el mismo estado al inicial.
// Indicar y justificar la complejidad de la primitiva.

package main

func (pila *pilaDinamica[T]) Transformar(aplicar func(T) T) *pilaDinamica[T] {
	pilaAuxiliar := &pilaDinamica[T]{make([]T, pila.cantidad), pila.cantidad}
	for i := 0; i < pila.cantidad; i++ {
		elemento := aplicar(pila.datos[i])
		pilaAuxiliar.datos = append(pilaAuxiliar.datos, elemento)
	}
	return pilaAuxiliar
}
