# El diámetro de una red es el máximo de las distancias mínimas entre todos los vértices de la misma. 
# Implementar un algoritmo que permita obtener el diámetro de una red, para el caso de un grafo no dirigido y no pesado. 
# Indicar el orden del algoritmo propuesto.

from collections import deque

def diametro(grafo) -> int:
    diametro_maximo = 0
    for v in grafo:
        diametro = bfs(grafo, v)
        if diametro > diametro_maximo:
            diametro_maximo = diametro
    return diametro_maximo

def bfs(grafo, origen) -> int:
    visitados = set()
    orden = {}
    visitados.add(origen)
    orden[origen] = 0
    orden_max = 0
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                orden[w] = orden[v] + 1
                q.append(w)
                if orden[w] > orden_max:
                    orden_max = orden[w]
    return orden_max
