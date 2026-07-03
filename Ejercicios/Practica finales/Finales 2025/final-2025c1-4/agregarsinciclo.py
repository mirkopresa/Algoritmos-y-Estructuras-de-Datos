# Implementar un algoritmo que dado un grafo dirigido y sin ciclos, y una lista de n pares de vértices (v, w) (siendo
# que no existe arista (v, w) ni (w, v) para ningún par de la lista), agregue una arista al grafo por cada par, con algún
# sentido (ya sea (v, w) o (w, v)) de tal forma que el mismo se mantenga sin ciclos. Se espera que el algoritmo ejecute
# lineal tanto en vértices y aristas, así como también en la cantidad de elementos de la lista. Es decir, O(V + E + n).
# Justificar la complejidad del algoritmo implementado. No hay ningún par repetido en ningún sentido.

# Grafo aciclico -> orden topologico

from collections import deque


def agregar_aristas(grafo, nuevas: list[tuple]):
    orden = orden_topologico(grafo)
    posiciones = {}
    i = 0
    for ver in orden:
        posiciones[ver] = i
        i += 1
    for par in nuevas:
        v, w = par
        if posiciones[v] < [w]:
            grafo.agregar_arista(v, w)
        else:
            grafo.agregar_arista(w, v)


def orden_topologico(grafo) -> list:
    grados_e = grados_entrada(grafo)
    cola = deque()
    for v in grafo:
        if grados_e[v] == 0:
            cola.append(v)
    orden = []
    while cola:
        v = cola.popleft()
        orden.append(v)
        for w in grafo.adyacentes(v):
            grados_e[w] -= 1
            if grados_e[w] == 0:
                cola.append(w)
    return orden


def grados_entrada(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados
