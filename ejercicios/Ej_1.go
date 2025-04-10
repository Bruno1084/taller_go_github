/* 
Ejercicio:
Una empresa realizó una encuesta de satisfacción entre sus clientes.
Cada cliente dejó su opinión en forma de puntaje del 1 al 5 (siendo 5 el mejor).

Tareas:
- Guardar en una lista los 10 primeros puntajes recibidos.
- Contar cuántas veces aparece cada puntaje.
- Mostrar por pantalla
	-La cantidad de veces que se votó cada puntaje.
	- Si hay más votos positivos (4 o 5) que negativos (1 o 2), imprimí
	  "¡Buen resultado!", si no, imprimí "Resultado mejorable".
- Subir a un repositorio de ustedes en github.
*/

package main
import "fmt"


func main() {
	var puntajes = [12]int{1, 5, 1, 2, 5, 4, 2, 1, 3, 3, 4, 1}

	fmt.Printf("Puntajes: %v \n", puntajes)
	fmt.Printf("Puntajes recientes: %v \n", puntajesRecientes(puntajes[:]))
	fmt.Printf("Estadisticas por puntaje: %v\n", contarPuntaje(puntajes[:]))
	fmt.Printf("Resultado general:\n %v", resultadoGeneral(puntajes[:]))
}

func puntajesRecientes(puntajes []int) []int {
	var sliceArr = puntajes[len(puntajes) -10:]

	return sliceArr
}

func contarPuntaje(puntajes []int) map[int]int{
	var contPuntajes = map[int]int{}

	for i:=0; i < len(puntajes); i++ {
		var _, keyExist = contPuntajes[puntajes[i]]
		if(!keyExist) {
			contPuntajes[puntajes[i]] += 1
		} else {
			contPuntajes[puntajes[i]] += 1
		}
	}

	return contPuntajes
}

func resultadoGeneral(puntajes []int)string {
	var resultado = ""
	var buenPuntaje int = 0
	var malPuntaje int = 0

	//Tiene que haber una forma más eficiente de hacer esto:
	for i:=0; i < len(puntajes); i++ {
		if(puntajes[i] > 3) {
			buenPuntaje++
		} else {
			malPuntaje++
		}
	}

	if(buenPuntaje > malPuntaje) {
		resultado = "Buen resultado :D"
	} else {
		resultado = "Resultado mejorable :/"
	}

	return resultado
}