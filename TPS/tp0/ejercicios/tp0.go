package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	maximo := 0
	for i := 1; i < len(vector); i++ {
		if vector[i] > vector[maximo] {
			maximo = i
		}
	}
	return maximo
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	longitud := min(len(vector1), len(vector2))
	for i := 0; i < longitud; i++ {
		if vector1[i] > vector2[i] {
			return 1
		} else if vector1[i] < vector2[i] {
			return -1
		}
	}
	if len(vector1) > len(vector2) {
		return 1
	} else if len(vector1) < len(vector2) {
		return -1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector) - 1; i > 0; i-- {
		maximo := Maximo(vector[:i+1])
		Swap(&vector[i], &vector[maximo])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func sumaRecursiva(vector []int, largo int) int {
	if largo == 0 {
		return 0
	}
	return vector[largo-1] + sumaRecursiva(vector, largo-1)
}

func Suma(vector []int) int {
	return sumaRecursiva(vector, len(vector))
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func capicuaRecursivo(cadena string, inicio int, fin int) bool {
	if inicio == fin || inicio > fin {
		return true
	}
	if cadena[inicio] == cadena[fin-1] {
		return capicuaRecursivo(cadena, inicio+1, fin-1)
	}
	return false
}

func EsCadenaCapicua(cadena string) bool {
	return capicuaRecursivo(cadena, 0, len(cadena))
}
