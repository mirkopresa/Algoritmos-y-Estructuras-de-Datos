# Se tiene un grafo no dirigido sin pesos, que representa una huerta de zanahorias. Cada vértice representa a una zanahoria y
# cada arista (v, w) indica que v y w se encuentran lo suficientemente cerca como para que si una se pudre, entonces podrá
# “contagiar a la otra”. Luego, esta nueva zanahoria contagiada infectará a las otras que se encuentren cercanas. En una unidad
# de tiempo, una zanahoria podrida infecta a todas las que tiene a su alcance. Suponiendo que existe una única componente
# conexa en dicho grafo, implementar una función que reciba un grafo de dichas características y una zanahoria podrida inicial,
# e indique cuál o cuáles son las últimas zanahorias en pudrirse. Indicar y justificar la complejidad de la función.

from collections import deque

def obtener_ultimas_zanahorias(grafo, inicio) -> list:
    visitados = set()
    orden = {}
    visitados.add(inicio)
    orden[inicio] = 0
    orden_maximo = 0
    q = deque()
    q.append(inicio)
    while q: # O(V + E)
        z = q.popleft()
        for w in grafo.adyacentes(z):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[z] + 1
                q.append(w)
        if orden[z] > orden_maximo:
            orden_maximo = orden[z]
    ultimas = []
    for v in orden: # O(V)
        if orden[v] == orden_maximo:
            ultimas.append(v)
    return ultimas
