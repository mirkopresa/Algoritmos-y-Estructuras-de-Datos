# En nuestra huerta de zanahorias se está construyendo un sistema de riego. La plantación completa es muy extensa y entre
# algunas zonas no pueden ubicarse mangueras, por lo que no es tarea sencilla determinar cómo conectar las mangueras para
# hacer llegar el agua a todas las zonas.
# Se tiene un grafo pesado en el que se representan como vértices estas zonas, y las aristas determinan la distancia entre zona y
# zona (en caso de poder conectarse entre sí). Definir una función que, dado el grafo mencionado, devuelva la mejor forma de
# conectar las distintas zonas, se requiere:
# Minimizar la cantidad de metros de manguera a utilizar.
# Que ninguna zona quede desconectada del resto de la red.
# Indicar y justificar la complejidad de la función.

def mst_kruskal(huerta):
    aristas = sorted(obtener_aristas(huerta))
    conjuntos = UnionFind(huerta.obtener_vertices()) # imaginate que esta implementado
    arbol = Grafo(False, huerta.obtener_vertices()) # imaginate que esta implementado
    for a in aristas:
        v, w, peso = a
        if conjuntos.find(v) != conjuntos.find(w):
            arbol.agregar_arista(v, w, peso)
            conjuntos.union(v, w)
    return arbol

def obtener_aristas(huerta):
    visitados = set()
    aristas = []
    for z in huerta:
        for w in huerta.adyacentes(z):
            if w not in visitados:
                aristas.append((z, w, huerta.peso_arista(z, w)))
        visitados.add(z)
    return aristas