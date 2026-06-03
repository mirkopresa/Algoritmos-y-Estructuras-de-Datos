# Implementar un algoritmo que determine si un grafo no dirigido es conexo o no. 
# Indicar la complejidad del algoritmo si el grafo está implementado con una matriz de adyacencia.

def es_conexo(grafo):
    if len(grafo) == 0:
        return True
    visitados = set()
    vertice_aleatorio = grafo.vertice_aleatorio()
    visitados.add(vertice_aleatorio)
    dfs(grafo, vertice_aleatorio, visitados)  
    return len(visitados) == len(grafo)

def dfs(grafo, v, visitados):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            dfs(grafo, w, visitados)