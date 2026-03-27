package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const (
	_RUTA1 = "archivo1.in"
	_RUTA2 = "archivo2.in"
)

// Función que recibe la ruta de un archivo y crea un arreglo con sus datos
func cargarArreglo(ruta string) []int {
	vector := make([]int, 0)

	archivo, err := os.Open(ruta)
	if err != nil {
		return vector
	}
	defer archivo.Close()

	lector := bufio.NewScanner(archivo)
	for lector.Scan() {
		numero, _ := strconv.Atoi(lector.Text())
		vector = append(vector, numero)
	}

	return vector
}

func imprimirArreglo(vector []int) {
	for _, value := range vector {
		fmt.Println(value)
	}
}

func imprimirSliceMayor(slice1, slice2 []int) {
	sliceMayor := slice1
	if ejercicios.Comparar(slice1, slice2) == -1 { // El segundo arreglo es mas grande que el primero
		sliceMayor = slice2
	}
	ejercicios.Seleccion(sliceMayor)
	imprimirArreglo(sliceMayor)
}

func main() {
	slice1, slice2 := cargarArreglo(_RUTA1), cargarArreglo(_RUTA2)
	imprimirSliceMayor(slice1, slice2)
}
