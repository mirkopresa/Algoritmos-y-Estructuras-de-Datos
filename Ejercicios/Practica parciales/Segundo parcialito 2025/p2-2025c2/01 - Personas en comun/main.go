// A lo largo de su trayectoria, la empresa ha tenido varias rotaciones de equipos. Próximamente habrá una nueva, y se busca que los
// nuevos equipos incluyan personas que hayan compartido equipos antes, para facilitar la transición.
// Se dispone de una base de datos que registra, para N personas, en qué M equipos ha participado cada una. La base usa un hash donde
// la persona es la clave y el valor, una lista de equipos:
// {
// 'Ana': ['Frontend-Platform', 'Growth-Squad'],
// 'Beto': ['Backend-Services', 'Frontend-Platform', 'Mobile-Core'],
// 'Carla': ['Mobile-Core'],
// }
// Realizar una función personasEnComun que reciba el hash, y el nombre de una persona , y devuelva una lista con todas las personas
// que hayan trabajado en al menos uno de sus equipos listados. Indicar y justificar la complejidad del algoritmo implementado,
// expresada con las variables N y M del problema. Por ejemplo, si se consulta por ‘Beto’, la respuesta incluye a ‘Ana’ y ‘Carla’. Si se
// pregunta por ‘Beto’, la respuesta incluye a ‘Beto’.

package main

func personasEnComun(dict Diccionario[string, Lista[string]], persona string) Lista[string] {
	empresas := CrearHash[string, bool]()
	lista := dict.Obtener(persona)
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		dato := iter.VerActual()
		empresas.Guardar(dato, true)
	}
	res := CrearListaEnlazada[string]()
	for iter := dict.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		clave, lista := iter.VerActual()
		for iter2 := lista.Iterador(); iter2.HayAlgoMas(); iter2.Avanzar() {
			if empresas.Pertenece(iter2.VerActual()) {
				res.InsertarUltimo(clave)
				break
			}
		}
	}
	return res
}
