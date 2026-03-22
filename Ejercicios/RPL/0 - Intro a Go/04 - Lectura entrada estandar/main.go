// Implementar un programa que lea la entrada del usuario (es decir, entrada estándar)
// hasta que se ingrese una 's', e imprima la cantidad de líneas leídas.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func IngresoPalabras() int {
	contador := 0
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		if lector.Text() == "s" {
			break
		}
		contador++
	}
	return contador
}

func main() {
	fmt.Println(IngresoPalabras())
}
