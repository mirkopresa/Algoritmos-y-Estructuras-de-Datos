package cola_prioridad

type colaConPrioridad[T any] struct {
	array    []T
	cantidad int
	cmp      func(T, T) int
}

const (
	_CAPACIDAD_MINIMA = 10
	_REDUCCION        = 4
	_REDIMENSION      = 2
)

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{array: make([]T, _CAPACIDAD_MINIMA), cantidad: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arrNuevo := make([]T, len(arreglo))
	copy(arrNuevo, arreglo)
	return heapify(arrNuevo, funcion_cmp)
}

func heapify[T any](arreglo []T, funcion_cmp func(T, T) int) *colaConPrioridad[T] {
	heap := &colaConPrioridad[T]{array: arreglo, cantidad: len(arreglo), cmp: funcion_cmp}
	for i := heap.posicionPadre(heap.cantidad - 1); i >= 0; i-- {
		heap.downheap(i)
	}
	return heap
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := heapify(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		swap(&heap.array[0], &heap.array[i])
		heap.cantidad--
		heap.downheap(0)
	}
}

func (c *colaConPrioridad[T]) EstaVacia() bool {
	return c.cantidad == 0
}

func (c *colaConPrioridad[T]) Encolar(elem T) {
	if cap(c.array) == c.cantidad {
		c.redimensionar(cap(c.array) * _REDIMENSION)
	}
	c.array[c.cantidad] = elem
	c.cantidad++
	c.upheap(c.cantidad - 1)
}

func (c *colaConPrioridad[T]) VerMax() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.array[0]
}

func (c *colaConPrioridad[T]) Desencolar() T {
	desencolado := c.VerMax()
	swap(&c.array[0], &c.array[c.cantidad-1])
	c.cantidad--
	c.downheap(0)
	if c.cantidad <= cap(c.array)/_REDUCCION && cap(c.array) > _CAPACIDAD_MINIMA {
		c.redimensionar(cap(c.array) / _REDIMENSION)
	}
	return desencolado
}

func (c *colaConPrioridad[T]) Cantidad() int {
	return c.cantidad
}

func (c *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {
	if nuevaCapacidad < _CAPACIDAD_MINIMA {
		nuevaCapacidad = _CAPACIDAD_MINIMA
	}
	nuevoArray := make([]T, nuevaCapacidad)
	copy(nuevoArray, c.array)
	c.array = nuevoArray
}

func (c *colaConPrioridad[T]) posicionPadre(posicion int) int {
	return (posicion - 1) / 2
}

func (c *colaConPrioridad[T]) posicionHijoIzq(posicion int) int {
	return 2*posicion + 1
}

func (c *colaConPrioridad[T]) posicionHijoDer(posicion int) int {
	return 2*posicion + 2
}

func swap[T any](x *T, y *T) {
	*x, *y = *y, *x
}

func (c *colaConPrioridad[T]) upheap(posicionElem int) {
	if posicionElem == 0 {
		return
	}
	posicionPadre := c.posicionPadre(posicionElem)
	comparacion := c.cmp(c.array[posicionPadre], c.array[posicionElem])
	if comparacion < 0 {
		swap(&c.array[posicionPadre], &c.array[posicionElem])
		c.upheap(posicionPadre)
	}
}

func (c *colaConPrioridad[T]) downheap(posicionElem int) {
	posicionHijoIzq := c.posicionHijoIzq(posicionElem)
	if posicionHijoIzq >= c.cantidad {
		return // no hay ni hijo izquierdo, ni derecho
	}
	hijoMax := posicionHijoIzq
	posicionHijoDer := c.posicionHijoDer(posicionElem)
	if posicionHijoDer < c.cantidad {
		comparacionHijos := c.cmp(c.array[posicionHijoIzq], c.array[posicionHijoDer])
		if comparacionHijos < 0 {
			hijoMax = posicionHijoDer
		}
	}
	comparacion := c.cmp(c.array[hijoMax], c.array[posicionElem])
	if comparacion > 0 {
		swap(&c.array[hijoMax], &c.array[posicionElem])
		c.downheap(hijoMax)
	}
}
