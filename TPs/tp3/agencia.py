from grafo import Grafo
import biblioteca


class Agencia:
    def __init__(self, ruta_archivo: str) -> None:
        self.grafo = Grafo()
        self.coordenadas = {}
        self._cargar_grafo_y_coordenadas(ruta_archivo)

    def ir(self, inicio, fin, archivo: str):
        if inicio not in self.grafo or fin not in self.grafo:
            return "No se encontro recorrido", ""
        padre, distancia = biblioteca.dijkstra(self.grafo, inicio, fin)
        if distancia[fin] == float("inf"):
            return "No se encontro recorrido", ""
        camino = biblioteca.reconstruir_camino(padre, fin)
        resultado = " -> ".join(camino)
        resultado_tiempo = f"Tiempo total: {distancia[fin]}"
        self._crear_kml(camino, archivo, self.coordenadas)
        return resultado, resultado_tiempo

    def itinerario(self, archivo: str) -> str:
        grafo_dirigido = Grafo(
            es_dirigido=True, vertices_init=self.grafo.obtener_vertices()
        )
        with open(archivo, "r") as f:
            for linea in f:
                ciudades = linea.strip().split(",")
                ciudad_1 = ciudades[0].strip()
                ciudad_2 = ciudades[1].strip()
                grafo_dirigido.agregar_arista(ciudad_1, ciudad_2, 1)
        orden = biblioteca.orden_topologico(grafo_dirigido)
        if len(orden) != len(grafo_dirigido):
            return "No se encontro recorrido"
        else:
            return " -> ".join(orden)

    def viaje(self, origen, archivo: str):
        camino = biblioteca.hierzoler(self.grafo, origen)
        if camino is None:
            return "No se encontro recorrido", ""
        tiempo_total = 0
        for i in range(len(camino) - 1):
            v = camino[i]
            w = camino[i + 1]
            tiempo_total += self.grafo.peso_arista(v, w)
        resultado = " -> ".join(camino)
        resultado_tiempo = f"Tiempo total: {tiempo_total}"
        self._crear_kml(camino, archivo, self.coordenadas)
        return resultado, resultado_tiempo

    def reducir_caminos(self, archivo: str) -> str:
        arbol = biblioteca.mst_prim(self.grafo)
        self._crear_pajek(arbol, archivo)
        peso = 0
        visitados = set()
        for v in arbol:
            for w in arbol.adyacentes(v):
                if w in visitados:
                    continue
                peso += arbol.peso_arista(v, w)
            visitados.add(v)
        return f"Peso total: {peso}"

    def _crear_pajek(self, arbol: Grafo, archivo: str) -> None:
        visitados = set()
        with open(archivo, "w") as f:
            f.write(f"{len(arbol)}\n")
            for v in arbol:
                lat, long = self.coordenadas[v]
                f.write(f"{v},{lat},{long}\n")
            f.write(f"{len(arbol)-1}\n")
            for v in arbol:
                for w in arbol.adyacentes(v):
                    if w in visitados:
                        continue
                    f.write(f"{v},{w},{arbol.peso_arista(v, w)}\n")
                visitados.add(v)

    def _cargar_grafo_y_coordenadas(self, ruta_archivo: str) -> None:
        with open(ruta_archivo, "r") as archivo:
            cantidad_ciudades = archivo.readline().strip()
            for _ in range(int(cantidad_ciudades)):
                ciudad = archivo.readline().strip()
                partes = ciudad.split(",")
                nombre = partes[0].strip()
                lat = float(partes[1].strip())
                lon = float(partes[2].strip())
                self.coordenadas[nombre] = (lat, lon)
                self.grafo.agregar_vertice(nombre)
            cantidad_rutas = archivo.readline().strip()
            for _ in range(int(cantidad_rutas)):
                ruta = archivo.readline().strip()
                partes = ruta.split(",")
                ciudad1 = partes[0].strip()
                ciudad2 = partes[1].strip()
                tiempo = int(partes[2].strip())
                self.grafo.agregar_arista(ciudad1, ciudad2, tiempo)

    def _crear_kml(self, camino, archivo: str, coordenadas: dict) -> None:
        ciudades = set()
        with open(archivo, "w") as f:
            f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
            f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
            f.write("\t<Document>\n")
            f.write("\n")
            f.write(
                f"\t\t<name>Camino desde {camino[0]} hacia {camino[len(camino) - 1]}</name>\n"
            )
            for ciudad in camino:
                if ciudad in ciudades:
                    continue
                ciudades.add(ciudad)
                lat, long = coordenadas[ciudad]
                f.write("\t\t<Placemark>\n")
                f.write(f"\t\t\t<name>{ciudad}</name>\n")
                f.write("\t\t\t<Point>\n")
                f.write(f"\t\t\t\t<coordinates>{long}, {lat}</coordinates>\n")
                f.write("\t\t\t</Point>\n")
                f.write("\t\t</Placemark>\n")
            f.write("\n")
            for i in range(len(camino) - 1):
                ciudad_1 = camino[i]
                ciudad_2 = camino[i + 1]
                lat1, long1 = coordenadas[ciudad_1]
                lat2, long2 = coordenadas[ciudad_2]
                f.write("\t\t<Placemark>\n")
                f.write("\t\t\t<LineString>\n")
                f.write(
                    f"\t\t\t\t<coordinates>{long1}, {lat1} {long2}, {lat2}</coordinates>\n"
                )
                f.write("\t\t\t</LineString>\n")
                f.write("\t\t</Placemark>\n")
            f.write("\t</Document>\n")
            f.write("</kml>\n")
