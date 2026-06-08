# Juan, un fanático de The Red Hot Chili Peppers, y un gran ayudante de diversas materias, está obligando a sus corregidos
# a instruirse en el mundo de la música y les insiste que escuche a los Red Hot. Los alumnos después de un par de reuniones
# escuchandolo quejarse de lo poco que saben de música, un poco cansados de esquivar la propuesta, finalmente la aceptan. Juan
# antes de decirles qué álbumes escuchar, les aclara que no sean nabos, y que hay ciertos albumes que deben escuchar antes que
# otros: imaginate que alguien escuche The Getaway antes que By the Way! Esta información la brinda con el siguiente formato:
# para cada álbum menciona una lista de álbumes que hay que escuchar antes que ese.
# a) Modelar este problema con grafos, indicando claramente qué son los vértices y qué son las aristas. (hecho en cuaderno)
# b) Implementar un algoritmo que le de a los alumnos un orden posible para escuchar los albumes que les recomendó Juan.

def obtener_grados(grafo):
    grados_entrada = {}
    for v in grafo: grados_entrada[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados_entrada[w] += 1
    return grados_entrada

def orden_album(grafo):
    grados_e = obtener_grados(grafo)
    cola = Cola() # considerar como implementado
    for v in grafo:
        if grados_e[v] == 0:
            cola.encolar(v)
    orden = []
    while not cola.esta_vacia(v):
        v = cola.desencolar()
        orden.append(v)
        for w in grafo.adyacentes(v):
            grados_e[w] -= 1
            if grados_e[w] == 0:
                cola.encolar(w)
    return orden