# Ejercicio 3: AlgoConnect (error 404: chiste not found)
# La exitosa empresa AlgoConnect está desarrollando dos aplicaciones:
# AlgoFriends: Es una red social en donde los usuarios pueden subir imágenes y entablar relaciones. Cualquier usuario puede conectarse con cualquier otro.
# AlgoBuy: Es una plataforma de ventas, en donde cada usuario se registra como comprador o vendedor. El único motivo de conexión posible entre usuarios es hacer consultas sobre un producto, es por esto que los compradores sólo pueden interactuar con vendedores, y viceversa.
# Cada aplicación cuenta con su propio grafo, en ambos casos los grafos son no dirigidos, no pesados, en donde los vértices son los usuarios y las aristas representan que dos usuarios se han conectado.
# Lamentablemente hubo un error en el software que borró las etiquetas de estos grafos, por lo que debemos re-etiquetarlos para saber si cada grafo corresponde a AlgoFriends o a AlgoBuy. Implementar una función que dado dos grafos nos devuelva dos Strings que correspondan a la asignación de cada uno de los grafos. Debe devolver "AlgoFriends", "AlgoBuy" si el primero correspondería al de AlgoFriends y el segundo al de AlgoBuy, "AlgoBuy", "AlgoFriends", si es lo contrario, o directamente None, None si no es posible determinar. Indicar y justificar la complejidad de la función.

def identificar_grafos(grafo_1, grafo_2):
    g1_bipartito = es_bipartito(grafo_1)
    g2_bipartito = es_bipartito(grafo_2)
    if g1_bipartito and not g2_bipartito:
        return "AlgoBuy", "AlgoFriends"
    if not g1_bipartito and g2_bipartito:
        return "AlgoFriends", "AlgoBuy"
    return None, None

def es_bipartito(grafo) -> bool:
    visitados = set()
    orden = {}
    for v in grafo:
        if v not in visitados:
            if not bfs_bipartito(grafo, v, visitados, orden):
                return False
    return True

def bfs_bipartito(grafo, origen, visitados, orden) -> bool:
    visitados.add(origen)
    orden[origen] = 0
    cola = Cola() # considerar implementada
    cola.encolar(origen)
    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                cola.encolar(w)
                orden[w] = orden[v] + 1
                visitados.add(w)
            elif orden[v] == orden[w]:
                return False
    return True
