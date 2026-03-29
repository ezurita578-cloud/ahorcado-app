package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var palabras = []string{
	"AGUACATE", "MARIPOSA", "CASTILLO", "ELEFANTE",
	"GUITARRA", "HOSPITAL", "MONTANA", "VOLCAN",
	"DELFIN", "MOCHILA", "PARAGUAS", "SEMAFORO",
}

var etapas = []string{
	`
  +---+
  |   |
      |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
      |
      |
=========`,
	`
  +---+
  |   |
  O   |
  |   |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|   |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
      |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
=========`,
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
=========`,
}

func letraUsada(usadas []string, letra string) bool {
	for _, l := range usadas {
		if l == letra {
			return true
		}
	}
	return false
}

func mostrarPalabra(palabra string, usadas []string) string {
	resultado := ""
	for _, ch := range palabra {
		letra := string(ch)
		if letraUsada(usadas, letra) {
			resultado += letra + " "
		} else {
			resultado += "_ "
		}
	}
	return resultado
}

func gano(palabra string, usadas []string) bool {
	for _, ch := range palabra {
		if !letraUsada(usadas, string(ch)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== JUEGO DEL AHORCADO ===")
	fmt.Println("Adivina la palabra letra por letra.")
	fmt.Println("Tienes 6 intentos.")
	fmt.Println()

	for {
		indice := rand.Intn(len(palabras))
		palabra := palabras[indice]
		errores := 0
		usadas := []string{}

		for errores < 6 {
			fmt.Println(etapas[errores])
			fmt.Println("Palabra:", mostrarPalabra(palabra, usadas))
			fmt.Println("Letras usadas:", strings.Join(usadas, ", "))
			fmt.Printf("Errores: %d/6\n", errores)

			if gano(palabra, usadas) {
				fmt.Println("¡GANASTE! La palabra era:", palabra)
				break
			}

			fmt.Print("Ingresa una letra: ")
			var letra string
			fmt.Scan(&letra)
			letra = strings.ToUpper(letra)

			if letraUsada(usadas, letra) {
				fmt.Println("Ya usaste esa letra, intenta otra.")
				continue
			}

			usadas = append(usadas, letra)

			if strings.Contains(palabra, letra) {
				fmt.Println("¡Bien! La letra", letra, "está en la palabra.")
			} else {
				fmt.Println("Incorrecto.")
				errores++
			}
		}

		if errores == 6 {
			fmt.Println(etapas[6])
			fmt.Println("¡AHORCADO! La palabra era:", palabra)
		}

		fmt.Print("¿Otra partida? (s/n): ")
		var resp string
		fmt.Scan(&resp)
		if resp != "s" {
			break
		}
	}

	fmt.Println("¡Gracias por jugar!")
}
