# Un árbol es un grafo no dirigido que cumple con las siguientes propiedades:

# a. E = V - 1

# b. Es acíclico

# c. Es conexo

# Por teorema, si un grafo cumple dos de estas tres condiciones, será árbol (y por consiguiente, cumplirá la tercera). 
# Haciendo uso de ésto (y únicamente de ésto), se pide implementar una función que reciba un grafo no dirigido 
# y determine si se trata de un árbol, o no. 
# Indicar el orden de la función implementada.

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

def contar_aristas(grafo):
    contador = 0
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in visitados:
                contador += 1
        visitados.add(v)
    return contador


def es_arbol(g) -> bool:
    return es_conexo(g) and (contar_aristas(g) == len(g)-1)