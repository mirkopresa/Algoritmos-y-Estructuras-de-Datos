# Implementar un algoritmo que reciba un grafo dirigido y nos devuelva la cantidad de componentes débilmente
# conexas de este. Indicar y justificar la complejidad del algoritmo implementado. 

from grafo import Grafo
from collections import deque

def cantidad_componentes_debiles(grafo) -> int:
    grafo_no_dirigido = Grafo(es_dirigido = False, vertices_init = grafo.obtener_vertices())
    for v in grafo:
        for w in grafo.adyacentes(v):
            if not grafo_no_dirigido.estan_unidos(v, w):
                grafo_no_dirigido.agregar_arista(v, w, 0)
    componentes = []
    visitados = set()
    for v in grafo:
        if v not in visitados:
            componentes.append(bfs(grafo_no_dirigido, v, visitados))
    return len(componentes)

def bfs(grafo, origen, visitados) -> list:
    vertices = []
    visitados.add(origen)
    vertices.append(origen)
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                vertices.append(w)
                q.append(w)
    return vertices