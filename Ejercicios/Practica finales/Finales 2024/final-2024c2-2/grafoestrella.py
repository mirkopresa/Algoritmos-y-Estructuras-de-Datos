# Implementar un algoritmo que reciba un grafo no dirigido y determine si el mismo tiene forma de estrella. Es decir,
# si todos los vértices, salvo 1, se conectan al mismo vértice, mientras ese único vértice se conecta con todos los demás.
# Indicar y justificar la complejidad del algoritmo si se implementara el grafo con una lista de adyacencia (diccionario de
# diccionarios), y también si se hiciera con una matriz de adyacencia.


def es_estrella(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    cant_mayor_1 = 0
    vertice_estrella = False
    for v in grados:
        grado = grados[v]
        if grado > 1:
            if grado == len(grafo) - 1:
                vertice_estrella = True
            cant_mayor_1 += 1
    return cant_mayor_1 == 1 and vertice_estrella
