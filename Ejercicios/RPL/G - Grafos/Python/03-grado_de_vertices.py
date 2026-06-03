# Implementar una función que determine el:

# a. El grado de todos los vértices de un grafo no dirigido.

# b. El grado de salida de todos los vértices de un grafo dirigido.

# c. El grado de entrada de todos los vértices de un grafo dirigido.

# Nota: Las funciones deberan devolver un diccionario con clave vertice y valor grado.

def grados(g):
    grados = {}
    for v in g:
        grados[v] = len(g.adyacentes(v))
    return grados

def grados_salida(g):
    return grados(g)

def grados_entrada(g):
    grados_e = {}
    lista_vertices = g.obtener_vertices()
    for v in lista_vertices:
        grados_e[v] = 0
    for v in g:
        for w in g.adyacentes(v):
            grados_e[w] += 1
    return grados_e
