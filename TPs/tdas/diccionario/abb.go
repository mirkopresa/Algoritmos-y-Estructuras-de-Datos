package diccionario

import TDAPila "tdas/pila"

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type nodoAbb[K comparable, V any] struct {
	clave K
	dato  V
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
}

func CrearABB[K comparable, V any](cmp funcCmp[K]) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato, izq: nil, der: nil}
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	padre, nodo := encontrarPadreYNodo(nil, ab.raiz, ab.cmp, clave)
	if nodo == nil {
		ab.cantidad++
		nuevoNodo := crearNodo(clave, dato)
		if padre == nil {
			ab.raiz = nuevoNodo
		} else {
			comparacion := ab.cmp(clave, padre.clave)
			if comparacion < 0 {
				padre.izq = nuevoNodo
			} else {
				padre.der = nuevoNodo
			}
		}
	} else {
		nodo.dato = dato
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	_, nodo := encontrarPadreYNodo(nil, ab.raiz, ab.cmp, clave)
	if nodo == nil {
		return false
	}
	return true
}

func (ab *abb[K, V]) Obtener(clave K) V {
	_, nodo := encontrarPadreYNodo(nil, ab.raiz, ab.cmp, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (ab *abb[K, V]) Borrar(clave K) V {
	padre, nodo := encontrarPadreYNodo(nil, ab.raiz, ab.cmp, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := nodo.dato
	ab.borrarNodo(padre, nodo)
	ab.cantidad--
	return dato
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	iterarNodo(ab.raiz, f, nil, nil, ab.cmp)
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarNodo(ab.raiz, visitar, desde, hasta, ab.cmp)
}

type iterAbb[K comparable, V any] struct {
	ab    *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter := &iterAbb[K, V]{
		ab:    ab,
		pila:  pila,
		desde: desde,
		hasta: hasta,
	}
	iter.apilarAnteriores(iter.ab.raiz)
	return iter
}

func (iter *iterAbb[K, V]) HayAlgoMas() bool {
	return iter.pila.EstaVacia() == false
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.pila.VerTope()
	return nodoActual.clave, nodoActual.dato
}

func (iter *iterAbb[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.pila.Desapilar()
	if nodoActual.der != nil {
		iter.apilarAnteriores(nodoActual.der)
	}
}

// Funciones auxiliares

func encontrarPadreYNodo[K comparable, V any](padre, nodo *nodoAbb[K, V], cmp funcCmp[K], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return padre, nil
	}
	comparacion := cmp(clave, nodo.clave)
	if comparacion < 0 {
		padre = nodo
		return encontrarPadreYNodo(padre, nodo.izq, cmp, clave)
	} else if comparacion > 0 {
		padre = nodo
		return encontrarPadreYNodo(padre, nodo.der, cmp, clave)
	} else {
		return padre, nodo
	}
}

func (ab *abb[K, V]) borrarNodo(padre, nodo *nodoAbb[K, V]) {
	if nodo.izq != nil && nodo.der != nil {
		anteriorPadre, anteriorNodo := ab.buscarAnterior(nodo, nodo.izq)
		nodo.clave = anteriorNodo.clave
		nodo.dato = anteriorNodo.dato
		ab.borrarNodo(anteriorPadre, anteriorNodo)
		return
	}
	var hijo *nodoAbb[K, V]
	if nodo.izq != nil {
		hijo = nodo.izq
	} else {
		hijo = nodo.der
	}
	if padre == nil {
		ab.raiz = hijo
	} else if padre.izq == nodo {
		padre.izq = hijo
	} else {
		padre.der = hijo
	}
}

func (ab *abb[K, V]) buscarAnterior(padre, hijo *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	for hijo.der != nil {
		padre = hijo
		hijo = hijo.der
	}
	return padre, hijo
}

func iterarNodo[K comparable, V any](nodo *nodoAbb[K, V], f func(K, V) bool, desde *K, hasta *K, comparar func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || comparar(nodo.clave, *desde) > 0 {
		if !iterarNodo(nodo.izq, f, desde, hasta, comparar) {
			return false
		}
	}
	dentroDesde := desde == nil || comparar(nodo.clave, *desde) >= 0
	dentroHasta := hasta == nil || comparar(nodo.clave, *hasta) <= 0

	if dentroDesde && dentroHasta && !f(nodo.clave, nodo.dato) {
		return false
	}

	if hasta == nil || comparar(nodo.clave, *hasta) < 0 {
		return iterarNodo(nodo.der, f, desde, hasta, comparar)
	}
	return true
}

func (iter *iterAbb[K, V]) apilarAnteriores(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.hasta != nil && iter.ab.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izq
		} else if iter.desde != nil && iter.ab.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
		} else {
			iter.pila.Apilar(nodo)
			nodo = nodo.izq
		}
	}
}
