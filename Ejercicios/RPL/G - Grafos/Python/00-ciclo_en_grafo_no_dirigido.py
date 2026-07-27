# Implementar un algoritmo que, dado un grafo no dirigido, nos devuelva un ciclo dentro del mismo, si es que los tiene. 
# Indicar el orden del algoritmo.

# Ejemplo

#     Para el grafo {A: [B], B: [A, C], C: [B]} el resultado sería lista vacía: []
#     Para el grafo {A: [B, C], B: [A, C], C: [B,A]} el resultado podría ser, entre otros, [A,B,C] 
#     ya que existe un camino que recorra A -> B -> C -> A

from collections import deque


def encontrar_ciclo(g) -> list | None:
    '''
    Devuelve una lista de vertices que conforman el ciclo. En el segundo ejemplo, 
    debería devolver [A, B, C] (o [B, C, A], etc...). 
    Si no hay ciclo, debe devolver None. 
    '''
    padres = {}
    visitados = set()
    for v in g:
        if v not in visitados:
            ciclo = ciclo_bfs(g, v, padres, visitados)
            if ciclo is not None:
                return ciclo
    return None

def ciclo_bfs(grafo, origen, padres, visitados) -> list | None:
    visitados.add(origen)
    padres[origen] = None
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                visitados.add(w)
                q.append(w)
            elif padres[w] != v:
                return reconstruir_ciclo(padres, v, w)
    return None

def reconstruir_ciclo(padres, inicio, fin) -> list:
    actual = fin
    camino = []
    while actual is not inicio:
        camino.append(actual)
        actual = padres[actual]
    camino.append(inicio)
    return camino.reverse()