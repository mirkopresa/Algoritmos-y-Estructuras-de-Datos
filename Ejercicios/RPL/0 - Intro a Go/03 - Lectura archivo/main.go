// Implementar una función que lea un archivo pasado por parámetro y
// cuenten la cantidad de caracteres alfanuméricos dentro del mismo (es decir, letras y números).

package main

import (
	"bufio"
	"os"
	"unicode"
)

func AlfanumericoEnArchivo(ruta string) int {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()

	contador := 0
	lector := bufio.NewScanner(archivo)

	for lector.Scan() {
		cadena := lector.Text()
		for _, caracter := range cadena {
			if unicode.IsDigit(caracter) || unicode.IsLetter(caracter) {
				contador++
			}
		}
	}

	return contador
}
