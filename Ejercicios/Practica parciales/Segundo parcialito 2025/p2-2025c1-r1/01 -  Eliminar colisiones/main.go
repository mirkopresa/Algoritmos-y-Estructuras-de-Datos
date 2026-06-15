// Implementar una primitiva eliminarColisiones(clave K) []K para el hash, que elimine del hash todas las
// claves que colisionen con la clave pasada por parámetro en el estado actual (sin eliminar dicha clave del
// diccionario, si se encuentra), y devuelva dichas claves. Implementar tanto para el hash abierto como para el hash
// cerrado. Si no se implementa para alguno, el ejercicio no estará aprobable. Indicar y justificar la complejidad de
// la primitiva para ambos casos.

package main

// Hash cerrado

func (hash *hashCerrado[K, V]) eliminarColisiones(clave K) []K {
	eliminados := make([]K, 0)
	arrClave := convertirABytes(clave)
	indiceActual := fnvHashing(arrClave, hash.tam)
	indiceOriginal := indiceActual
	for {
		celda := hash.tabla[indiceActual]
		if celda.estado == OCUPADO {
			arr := convertirABytes(celda.clave)
			posibleIndice := fnvHashing(arr, hash.tam)
			if posibleIndice == indiceOriginal && celda.clave != clave {
				hash.tabla[indiceActual].estado = BORRADO
				eliminados = append(eliminados, hash.tabla[indiceActual].clave)
				hash.cantidad--
			}

		} else if celda.estado == VACIO {
			break
		}
		indiceActual = (indiceActual + 1) % hash.tam
		if indiceActual == indiceOriginal {
			break
		}
	}
	// habria que redimensionar aca
	return eliminados
}

// Hash abierto

func (hash *hashAbierto[K, V]) eliminarColisiones2(clave K) []K {
	eliminados := make([]K, 0)
	arrClave := convertirABytes(clave)
	indiceActual := fnvHashing(arrClave, hash.tam)
	lista := hash.tabla[indiceActual]
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		elemento := iter.VerActual()
		if elemento.clave != clave {
			elemento.estado = BORRADO
			eliminados = append(eliminados, elemento.clave)
			hash.cantidad--
		}
	}
	// habria que redimensionar aca
	return eliminados
}
