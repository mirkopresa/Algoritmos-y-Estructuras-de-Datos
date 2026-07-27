from heap import Heap
from grafo import Grafo
from collections import deque

import sys

sys.setrecursionlimit(100000)


def dijkstra(grafo: Grafo, inicio, fin):
    distancia = {}
    padre = {}
    for v in grafo:
        distancia[v] = float("inf")

    distancia[inicio] = 0
    padre[inicio] = None

    heap = Heap(min_cmp)
    heap.encolar((0, inicio))

    while not heap.esta_vacio():
        _, v = heap.desencolar()
        if v == fin:
            break
        for w in grafo.adyacentes(v):
            distancia_actual = distancia[v] + grafo.peso_arista(v, w)
            if distancia_actual < distancia[w]:
                distancia[w] = distancia_actual
                padre[w] = v
                heap.encolar((distancia[w], w))
    return padre, distancia


def mst_prim(grafo: Grafo) -> Grafo:
    arbol = Grafo(es_dirigido=False, vertices_init=grafo.obtener_vertices())
    v = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    heap = Heap(min_cmp)
    for w in grafo.adyacentes(v):
        peso = grafo.peso_arista(v, w)
        heap.encolar((peso, v, w))

    while not heap.esta_vacio():
        peso, v, w = heap.desencolar()
        if w not in visitados:
            arbol.agregar_arista(v, w, peso)
            visitados.add(w)
            for u in grafo.adyacentes(w):
                if u not in visitados:
                    peso_arista = grafo.peso_arista(w, u)
                    heap.encolar((peso_arista, w, u))

    return arbol


def orden_topologico(grafo: Grafo) -> list:
    g_entrada = grados_entrada(grafo)
    cola = deque()
    for v in grafo:
        if g_entrada[v] == 0:
            cola.append(v)
    orden = []
    while cola:
        v = cola.popleft()
        orden.append(v)
        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.append(w)

    return orden


def ciclo_euleriano(grafo: Grafo) -> bool:
    if not es_conexo(grafo):
        return False
    for v in grafo:
        if len(grafo.adyacentes(v)) % 2 != 0:
            return False
    return True


def hierzoler(grafo: Grafo, inicio):
    aristas_visitadas = set()
    camino = []
    if not ciclo_euleriano(grafo):
        return None
    dfs_hierzoler(grafo, inicio, aristas_visitadas, camino)
    camino.reverse()
    return camino


def dfs_hierzoler(grafo: Grafo, v, aristas_visitadas: set, camino: list) -> None:
    for w in grafo.adyacentes(v):
        if (v, w) not in aristas_visitadas:
            aristas_visitadas.add((v, w))
            aristas_visitadas.add((w, v))
            dfs_hierzoler(grafo, w, aristas_visitadas, camino)
    camino.append(v)


def es_conexo(grafo: Grafo) -> bool:
    if len(grafo) == 0:
        return True
    visitados = set()
    vertice_aleatorio = grafo.vertice_aleatorio()
    visitados.add(vertice_aleatorio)
    dfs(grafo, vertice_aleatorio, visitados)
    return len(visitados) == len(grafo)


def dfs(grafo: Grafo, v, visitados: set) -> None:
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            dfs(grafo, w, visitados)


def reconstruir_camino(padre: dict, fin) -> list:
    camino = []
    actual = fin
    while actual is not None:
        camino.append(actual)
        actual = padre[actual]
    camino.reverse()
    return camino


def grados_entrada(grafo: Grafo) -> dict:
    g_entrada = {}
    for v in grafo:
        g_entrada[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            g_entrada[w] += 1
    return g_entrada


def min_cmp(a: tuple, b: tuple) -> int:
    if b[0] < a[0]:
        return -1
    elif b[0] > a[0]:
        return 1
    else:
        return 0
