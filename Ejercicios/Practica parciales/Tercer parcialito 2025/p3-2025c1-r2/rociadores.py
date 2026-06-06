# En nuestra huerta tenemos unos rociadores de insecticidas automáticos. Cada rociador cuenta con la dosis apropiada para
# cubrir hasta un máximo de 5 plantaciones a su alrededor. Es necesario averiguar si algún rociador tiene más de 5 plantaciones
# alrededor ya que de tener una mayor cantidad la dosis sería insuficiente. Se tiene un grafo en donde los vértices son los
# rociadores y plantas, es no pesado y dirigido (el origen de una arista es el rociador y el destino es una planta en su rango).
# Implementar una función que reciba este grafo y devuelva, en caso que un rociador tenga más de 5 plantaciones alrededor,
# el conjunto de plantaciones alrededor de dicho rociador (si hay más de un rociador que cumpla esta condición, devolver la
# información correspondiente a cualquiera de estos). En caso contrario, devolver None. Indicar y justificar la complejidad de la
# función.

def obtener_rociadores_sobrecargado(grafo) -> dict | None:
    sobrecargados = {}
    for v in grafo:
        adyacentes = grafo.adyacentes(v)
        if len(adyacentes) > 5:
            sobrecargados[v] = []
            for w in adyacentes:
                sobrecargados[v].append(w)
    if len(sobrecargados) == 0:
        return None
    return sobrecargados