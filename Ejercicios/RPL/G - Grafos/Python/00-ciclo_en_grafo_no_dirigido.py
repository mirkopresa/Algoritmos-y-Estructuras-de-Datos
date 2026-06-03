# Implementar un algoritmo que, dado un grafo no dirigido, nos devuelva un ciclo dentro del mismo, si es que los tiene. 
# Indicar el orden del algoritmo.

# Ejemplo

#     Para el grafo {A: [B], B: [A, C], C: [B]} el resultado sería lista vacía: []
#     Para el grafo {A: [B, C], B: [A, C], C: [B,A]} el resultado podría ser, entre otros, [A,B,C] 
#     ya que existe un camino que recorra A -> B -> C -> A


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
            visitados.add(v)
            padres[v] = None
            ciclo = ciclo_dfs(g, v, padres, visitados)
            if ciclo is not None:
                return ciclo
    return None

def ciclo_dfs(grafo, v, padres, visitados):
    for w in grafo.adyacentes(v):
        if w in visitados:
            if w != padres[v]:
                return reconstruir_ciclo(padres, w, v)
            else:
                padres[w] = v
                ciclo = ciclo_dfs(grafo, w, padres, visitados)
                if ciclo is not None:
                    return ciclo
    return None

def reconstruir_ciclo(padres, inicio, fin):
    v = fin
    camino = []
    while v != inicio:
        camino.append(v)
        v = padres[v]
    camino.append(inicio)
    return camino[::-1]