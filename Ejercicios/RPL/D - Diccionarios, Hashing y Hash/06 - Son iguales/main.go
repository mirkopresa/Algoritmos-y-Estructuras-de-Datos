// Asumiendo que se tiene disponible una implementación completa del TDA Hash,
// se desea implementar una función que decida si dos Hash dados representan o no el mismo Diccionario.
// Considere para la solución que es de interés la mejor eficiencia temporal posible.
// Indique, para su solución, eficiencia en tiempo y espacio.
// Nota: Dos tablas de hash representan el mismo diccionario si tienen la misma cantidad de elementos;
// todas las claves del primero están en el segundo; todas las del segundo, en el primero;
// y los datos asociados a cada una de esas claves son iguales (se pueden comparar los valores con “==”).

package main

// Complejidad temporal O(n), siendo n la cantidad de elementos de los diccionarios (ya que son iguales)
// Complejidad espacial O(1), no estamos generando memoria extra
func SonIguales[K comparable, V comparable](d1, d2 Diccionario[K, V]) bool {
	if d1.Cantidad() != d2.Cantidad() {
		return false
	}
	iter1, iter2 := d1.Iterador(), d2.Iterador()
	for iter1.HaySiguiente() && iter2.HaySiguiente() {
		claveD1, datoD1 := iter1.VerActual()
		claveD2, datoD2 := iter2.VerActual()
		if claveD1 != claveD2 || datoD1 != datoD2 {
			return false
		}
		iter1.Siguiente()
		iter2.Siguiente()
	}
	return true
}

// O hecho de otra manera

func SonIguales[K comparable, V comparable](d1, d2 Diccionario[K, V]) bool {
	if d1.Cantidad() != d2.Cantidad() {
		return false
	}
	iter1 := d1.Iterador()
	for iter1.HaySiguiente() {
		claveD1, datoD1 := iter1.VerActual()
		if !d2.Pertenece(claveD1) || d2.Pertenece(claveD1) && d2.Obtener(claveD1) != datoD1 {
			return false
		}
		iter1.Siguiente()
	}
	return true
}
