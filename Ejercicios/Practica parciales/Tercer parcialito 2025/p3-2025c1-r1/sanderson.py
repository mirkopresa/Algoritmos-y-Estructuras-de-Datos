# Un ayudante de esta materia quiere comenzar a leer los libros de Sanderson. Este autor tiene distintas sagas de libros
# interconectadas en un mismo universo, sumado a distintos libros independientes. Una opción podría ser leerlos orden de salida
# cronológica, pero hay algo más importante, la consideración del autor. En la web de Sanderson se puede encontrar, por cada
# libro, un listado de cuáles son los libros que es necesario leer con anterioridad.
# a. Describir detalladamente cómo se podría representar esta información en un grafo que nos sea útil para resolver este
# problema.
# b. Implementar una función que, dado un grafo definido como en el punto a, devuelva un orden posible que nuestro ayudante
# pueda seguir. Indicar y justificar la complejidad de la función.

def grados_entrada(grafo):
    grados = {}
    for v in grafo:
        grados[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados

def orden_topologico(grafo):
    g_entrada = grados_entrada(grafo)
    cola = Cola() #imagina que esta implementada
    for v in grafo:
        if g_entrada[v] == 0:
            cola.encolar(v)
    orden_libros = []
    while not cola.esta_vacia():
        libro = cola.desencolar()
        orden_libros.append(libro)
        for c in grafo.adyacentes(libro):
            g_entrada[c] -= 1
            if g_entrada[c] == 0:
                cola.encolar(c)
    return orden_libros