// Implementar una primitiva para el árbol binario que determine si cumple la propiedad de Heap. Indicar y justificar la complejidad
// del algoritmo implementado.

package main

type arbol struct {
	izq   *arbol
	der   *arbol
	clave int
}

func (ab *arbol) EsHeap() bool {
	if ab == nil {
		return true
	}
	heapIzq := ab.izq.EsHeap()
	heapDer := ab.der.EsHeap()
	if ab.izq != nil {
		if ab.clave < ab.izq.clave {
			return false
		}
	}
	if ab.der != nil {
		if ab.clave < ab.der.clave {
			return false
		}
	}
	return heapIzq && heapDer
}
