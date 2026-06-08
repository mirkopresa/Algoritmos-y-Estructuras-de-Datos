# Un autor decidió escribir un libro con varias tramas que se puede leer de forma no lineal. Es decir, por ejemplo, 
# después del capítulo 1 puede leer el 2 o el 73; pero la historia no tiene sentido si se abordan estos últimos antes que el 1.

# Siendo un aficionado de la computación, el autor ahora necesita un orden para publicar su obra, y decidió modelar este 
# problema como un grafo dirigido, en dónde los capítulos son los vértices y sus dependencias las aristas. 
# Así existen, por ejemplo, las aristas (v1, v2) y (v1, v73).

# Escribir un algoritmo que devuelva un orden en el que se puede leer la historia sin obviar ningún capítulo. 
# Indicar la complejidad del algoritmo.

from collections import deque

def grados_entrada(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados 


def obtener_orden(grafo):
    g_entrada = grados_entrada(grafo)
    q = deque()
    for v in grafo:
        if g_entrada[v] == 0:
            q.append(v)

    orden = []
    while q:
        v = q.popleft()
        orden.append(v)
        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                q.append(w)
    if orden is not None:
        return orden
    return None