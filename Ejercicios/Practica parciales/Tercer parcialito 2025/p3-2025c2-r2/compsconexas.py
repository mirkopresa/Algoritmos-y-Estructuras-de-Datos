# Implementar un algoritmo que encuentre todas las componentes conexas de un grafo no dirigido. Indicar y justificar la
# complejidad del algoritmo.

def componentes(grafo):
    visitados = set()
    resultado = []
    for v in grafo:
        if v not in visitados:
            aux = bfs(grafo, v, visitados)
            resultado.append(aux)
    return resultado

def bfs(grafo, origen, visitados):
    visitados.add(origen)
    comps = []
    cola = Cola() # ignorar
    cola.encolar(origen)
    while not cola.esta_vacia():
        v = cola.desencolar()
        comps.append(v)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.encolar(w)
    return comps