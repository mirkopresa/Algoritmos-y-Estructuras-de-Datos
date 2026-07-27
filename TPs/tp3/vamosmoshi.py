#!/usr/bin/python3
import sys
from agencia import Agencia


def cargar_comandos(agencia: Agencia):
    for linea in sys.stdin:
        linea = linea.strip()
        linea = linea.split(" ", 1)
        comando = linea[0]
        parametros_sin_procesar = linea[1]
        parametros_sin_procesar = parametros_sin_procesar.split(",")
        parametros = [parametro.strip() for parametro in parametros_sin_procesar]
        if comando == "ir":
            inicio = parametros[0]
            fin = parametros[1]
            archivo = parametros[2]
            resultado, resultado_tiempo = agencia.ir(inicio, fin, archivo)
            print(resultado)
            if resultado_tiempo != "":
                print(resultado_tiempo)
        elif comando == "itinerario":
            archivo = parametros[0]
            resultado = agencia.itinerario(archivo)
            print(resultado)
        elif comando == "viaje":
            origen = parametros[0]
            archivo = parametros[1]
            resultado, resultado_tiempo = agencia.viaje(origen, archivo)
            print(resultado)
            if resultado_tiempo != "":
                print(resultado_tiempo)
        elif comando == "reducir_caminos":
            archivo = parametros[0]
            resultado = agencia.reducir_caminos(archivo)
            print(resultado)
        else:
            return


def main():
    if len(sys.argv) < 2:
        sys.exit(1)
    agencia = Agencia(sys.argv[1])
    cargar_comandos(agencia)


if __name__ == "__main__":
    main()
