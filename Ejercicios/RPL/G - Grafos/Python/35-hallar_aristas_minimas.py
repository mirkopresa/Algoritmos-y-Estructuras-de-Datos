# Implementar un algoritmo que reciba un grafo no dirigido y determine la cantidad mínima de aristas que 
# debería agregársele para que el grafo sea conexo.
# Obviamente, si el grafo ya es conexo el algoritmo debe devolver 0. Indicar y justificar la complejidad del algoritmo implementado.

from collections import deque

def hallar_aristas_minimas(grafo):
    visitados = set()
    cantidad = 0
    if len(grafo) == 0:
        return 0
    for v in grafo:
        if v not in visitados:
            bfs(grafo, v, visitados)
            cantidad += 1
    return cantidad - 1

def bfs(grafo, origen, visitados):
    visitados.add(origen)
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                q.append(w)