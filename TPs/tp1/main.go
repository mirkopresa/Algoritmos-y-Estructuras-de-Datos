package main

import (
	"bufio"
	"os"
)

func main() {
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		linea := lector.Text()
		vectorInfix := LineaAVector(linea)
		vectorPostfix := ConvertirInfixToPostfix(vectorInfix)
		ImprimirResultado(vectorPostfix)
	}
}
