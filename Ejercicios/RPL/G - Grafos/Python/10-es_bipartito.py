# Escribir una función bool es_bipartito(grafo) que dado un grafo no dirigido devuelva true o false de acuerdo a si es bipartito o no. 
# Indicar y justificar el orden del algoritmo. ¿Qué tipo de recorrido utiliza?

from collections import deque

def es_bipartito(grafo) -> bool:
    visitados = set()
    orden = {}
    for v in grafo:
        if v not in visitados:
            orden[v] = 0
            visitados.add(v)
            bipartito = bfs(grafo, v, orden, visitados)
            if not bipartito:
                return False
    return True

def bfs(grafo, v, orden, visitados) -> bool:
    q = deque()
    q.append(v)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w in visitados:
                if orden[v] == orden[w]:
                    return False
            else:
                visitados.add(w)
                orden[w] = orden[v] + 1
                q.append(w)
    return True


