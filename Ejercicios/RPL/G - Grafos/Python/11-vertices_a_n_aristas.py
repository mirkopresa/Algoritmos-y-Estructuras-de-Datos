# Implementar un algoritmo que reciba un grafo dirigido, un vértice V y un número N, y devuelva una lista 
# con todos los vértices que se encuentren a exactamente N aristas de distancia del vértice V. 
# Indicar el tipo de recorrido utilizado y el orden del algoritmo. Justificar.

from collections import deque

def a_n_aristas(grafo, v, n) -> list | None:
    'Devolver una lista con los vértices que cumplen la propiedad'
    # Recorrido utilizado: BFS, complejidad O(V + E)
    # En el peor de los casos, la complejidad va a ser O(V + E) ya que si n es muy grande, podriamos llegar
    # a estar recorriendo casi todos o todos los vertices
    # Luego recorrer el diccionario es O(V) en el peor de los casos por la misma razon
    visitados = set()
    orden = {}
    visitados.add(v)
    orden[v] = 0
    q = deque()
    q.append(v)
    while q:
        actual = q.popleft()
        if orden[actual] == n:
            break
        for w in grafo.adyacentes(actual):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[actual] + 1
                q.append(w)
    lista = []
    for vertice in orden:
        if orden[vertice] == n:
            lista.append(vertice)
    if lista is not None:
        return lista
    else:
        return None