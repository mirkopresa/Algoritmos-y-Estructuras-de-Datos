package diccionario

import (
	"fmt"
)

type estado int

const (
	_VACIO estado = iota
	_OCUPADO
	_BORRADO
)

const (
	_CAPACIDAD_INICIAL  = 11
	_FACTOR_REDIMENSION = 0.70
	_FACTOR_REDUCCION   = 0.35
	_REDIMENSION        = 2
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fnvHashing(clave []byte, largo int) int {
	var h uint64 = 14695981039346656037
	for _, c := range clave {
		h *= 1099511628211
		h ^= uint64(c)
	}
	return int(h % uint64(largo))
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := crearTabla[K, V](_CAPACIDAD_INICIAL)
	return &hashCerrado[K, V]{tabla: tabla, cantidad: 0, tam: _CAPACIDAD_INICIAL, borrados: 0}
}

func (d *hashCerrado[K, V]) Guardar(clave K, dato V) {
	factor := obtenerFactor(d.cantidad, d.borrados, d.tam)
	if factor >= _FACTOR_REDIMENSION {
		d.redimensionar(d.tam * _REDIMENSION)
	}
	encontrado, indice := d.obtenerIndice(clave)
	if encontrado {
		d.tabla[indice].dato = dato
	} else {
		d.tabla[indice].clave = clave
		d.tabla[indice].dato = dato
		d.tabla[indice].estado = _OCUPADO
		d.cantidad++
	}
}

func (d *hashCerrado[K, V]) Pertenece(clave K) bool {
	encontrado, _ := d.obtenerIndice(clave)
	return encontrado
}

func (d *hashCerrado[K, V]) Obtener(clave K) V {
	encontrado, indice := d.obtenerIndice(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	return d.tabla[indice].dato
}

func (d *hashCerrado[K, V]) Borrar(clave K) V {
	factor := obtenerFactor(d.cantidad, d.borrados, d.tam)
	if factor <= _FACTOR_REDUCCION && d.tam > _CAPACIDAD_INICIAL {
		d.redimensionar(d.tam / _REDIMENSION)
	}
	encontrado, indice := d.obtenerIndice(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	d.tabla[indice].estado = _BORRADO
	d.cantidad--
	d.borrados++
	return d.tabla[indice].dato
}

func (d *hashCerrado[K, V]) Cantidad() int {
	return d.cantidad
}

func (d *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for i := 0; i < d.tam; i++ {
		if d.tabla[i].estado == _OCUPADO {
			if !f(d.tabla[i].clave, d.tabla[i].dato) {
				break
			}
		}
	}
}

type iterDiccionario[K comparable, V any] struct {
	dict   *hashCerrado[K, V]
	indice int
}

func (d *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterDiccionario[K, V]{dict: d, indice: 0}
	if iter.dict.tabla[iter.indice].estado != _OCUPADO {
		iter.proximoOcupado()
	}
	return iter
}

func (iter *iterDiccionario[K, V]) HayAlgoMas() bool {
	return iter.indice < iter.dict.tam
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	celda := iter.dict.tabla[iter.indice]
	return celda.clave, celda.dato
}

func (iter *iterDiccionario[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iter.indice++
	iter.proximoOcupado()
}

// Funciones auxiliares

func (d *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {
	tablaNueva := crearTabla[K, V](nuevaCapacidad)
	tablaVieja := d.tabla
	d.tabla, d.tam, d.cantidad, d.borrados = tablaNueva, nuevaCapacidad, 0, 0
	for _, elem := range tablaVieja {
		if elem.estado == _OCUPADO {
			d.Guardar(elem.clave, elem.dato)
		}
	}
}

func (iter *iterDiccionario[K, V]) proximoOcupado() {
	for iter.HayAlgoMas() && iter.dict.tabla[iter.indice].estado != _OCUPADO {
		iter.indice++
	}
}

func (d *hashCerrado[K, V]) obtenerIndice(clave K) (bool, int) {
	arrClave := convertirABytes(clave)
	indice := fnvHashing(arrClave, d.tam)
	for {
		if d.tabla[indice].clave == clave && d.tabla[indice].estado == _OCUPADO {
			return true, indice
		} else if d.tabla[indice].estado == _VACIO {
			return false, indice
		}
		indice = (indice + 1) % d.tam
	}
}

func obtenerFactor(cantidad, borrados, tam int) float64 {
	return float64(cantidad+borrados) / float64(tam)
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], capacidad)
}
