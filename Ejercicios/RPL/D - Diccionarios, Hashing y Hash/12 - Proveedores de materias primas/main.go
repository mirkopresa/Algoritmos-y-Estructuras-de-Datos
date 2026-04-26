// Una fábrica de pastas de Lanús le pide a alumnos de Algoritmos y Programación II que
// le solucionen un problema: sus dos distribuidoras de materias primas le enviaron un hash cada una,
// dónde sus claves son los nombres de los productos, y sus valores asociados, sus precios.
// La fábrica de pastas le pide a los alumnos que le implementen una función que le devuelva un
// nuevo hash con la unión de todos esos productos, y en caso de que una misma materia prima
// se encuentre en ambos hashes, elegir la que tenga el precio más barato.
// Indicar y justificar el orden del algoritmo.

package main

func MergeProveedores(prov1, prov2 Diccionario[string, int]) Diccionario[string, int] {
	resultado := CrearHash[string, int]()
	for iter := prov1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, precio := iter.VerActual()
		precioAGuardar := precio
		if prov2.Pertenece(clave) {
			_, precio2 := prov2.Obtener(clave)
			if precio2 < precioAGuardar {
				precioAGuardar = precio2
			}
		}
		resultado.Guardar(clave, precioAGuardar)
	}
	for iter := prov2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, precio := iter.VerActual()
		if prov1.Pertenece(clave) {
			continue
		}
		resultado.Guardar(clave, precio)
	}
	return resultado
}
