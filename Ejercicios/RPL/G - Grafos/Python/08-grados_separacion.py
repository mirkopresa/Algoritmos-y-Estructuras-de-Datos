# La teoría de los 6 grados de separación dice que cualquiera en la Tierra puede estar conectado a cualquier
# otra persona del planeta a través de una cadena de conocidos que no tiene más de cinco intermediarios (conectando a
# ambas personas con solo seis enlaces). Suponiendo que se tiene un grafo G en el que cada vértice es una persona y
# cada arista conecta gente que se conoce (el grafo es no dirigido):

# Implementar un algoritmo para comprobar si se cumple tal teoría para todo el conjunto de personas 
# representadas en el grafo G. Indicar el orden del algoritmo.

# Suponiendo que en el grafo G no habrán altas ni bajas de vértices, pero podrían haberla de aristas 
# (la gente se va conociendo), explicar las ventajas y desventajas que tendría implementar al grafo G 
# con una matriz de adyacencia.

from collections import deque

def seis_grados(grafo) -> bool:
    for v in grafo:
        se_cumple = bfs(grafo, v)
        if not se_cumple:
            return False
    return True

def bfs(grafo, origen):
    visitados = set()
    padres = {}
    orden = {}
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        if orden[v] > 6:
            return False
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.encolar(w)
    return len(orden) == len(grafo)