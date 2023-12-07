package main

import (
	Utils "adventOfCode/2023/utils"
	"fmt"
	"unicode"
)

const (
	filename = "inputDay03.txt"
)

func main() {
	contenidos, err := Utils.ReadTXTFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	for _, s := range contenidos {
		numbers := findSurroundedNumbers(s)
		fmt.Printf("String: %s, Números encontrados: %v\n", s, numbers)
	}

	//fmt.Println(contenidos)
}

func findSurroundedNumbers(s string) []string {
	var result []string
	var currentNumber string

	for _, char := range s {
		if unicode.IsDigit(char) {
			currentNumber += string(char)
		} else if char == '.' {
			// Continuar añadiendo dígitos al número actual si el punto está presente
			currentNumber += string(char)
		} else {
			// Se encontró un carácter no permitido, reiniciar el número actual
			currentNumber = ""
		}

		// Si encontramos un número rodeado únicamente por puntos, lo añadimos a los resultados
		if len(currentNumber) > 0 && string(char) != "." {
			result = append(result, currentNumber)
			currentNumber = ""
		}
	}

	return result
}
