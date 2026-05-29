// ¿Puede utilizarse un Heap para implementar el TDA cola (en el que se extraen los elementos en
// el orden en que fueron insertados)? Implementarlo.

package colaheap

type colaSobreHeap[T any] struct {
	heap     ColaPrioridad[elemento[T]]
	contador int
}

type elemento[T any] struct {
	dato  T
	orden int
}

func primerOrden[T any](elem1, elem2 elemento[T]) int {
	return elem2.orden - elem1.orden
}

func CrearColaSobreColaPrioridad[T any]() Cola[T] {
	return &colaSobreHeap[T]{heap: CrearHeap[elemento[T]](primerOrden[T]), contador: 0}
}

func (cola *colaSobreHeap[T]) Encolar(dato T) {
	cola.heap.Encolar(elemento[T]{dato: dato, orden: cola.contador})
	cola.contador++
}

func (cola *colaSobreHeap[T]) Desencolar() T {
	if cola.heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	elem := cola.heap.Desencolar()
	return elem.dato
}

func (cola *colaSobreHeap[T]) EstaVacia() bool {
	return cola.heap.EstaVacia()
}

func (cola *colaSobreHeap[T]) VerPrimero() T {
	return cola.heap.VerMax().dato
}
