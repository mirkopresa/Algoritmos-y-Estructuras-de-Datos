// ¿Puede utilizarse un Heap para implementar el TDA Pila
// (en el que se extraen los elementos en el orden inverso en que fueron insertados)?
// Implementarlo.

package pilaheap

type pilaSobreHeap[T any] struct {
	heap     ColaPrioridad[elemento[T]]
	contador int
}

type elemento[T any] struct {
	dato  T
	orden int
}

func primerOrden[T any](elem1, elem2 elemento[T]) int {
	return elem1.orden - elem2.orden
}

func CrearPilaSobreColaPrioridad[T any]() Pila[T] {
	return &pilaSobreHeap[T]{heap: CrearHeap[elemento[T]](primerOrden[T]), contador: 0}
}

func (pila *pilaSobreHeap[T]) Apilar(dato T) {
	pila.heap.Encolar(elemento[T]{dato: dato, orden: pila.contador})
	pila.contador++
}

func (pila *pilaSobreHeap[T]) Desapilar() T {
	if pila.heap.EstaVacia() {
		panic("La pila esta vacia")
	}
	elem := pila.heap.Desencolar()
	return elem.dato
}

func (pila *pilaSobreHeap[T]) EstaVacia() bool {
	return pila.heap.EstaVacia()
}

func (pila *pilaSobreHeap[T]) VerTope() T {
	if pila.heap.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.heap.VerMax().dato
}
