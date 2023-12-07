package main

import (
	"adventOfCode/2023/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var palabrasNumericas = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var temporary = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "4",
	"five":  "5e",
	"six":   "6s",
	"seven": "7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	filename := "inputDay01.txt"
	contenidos, err := Utils.ReadTXTFile(filename)
	var draftSlice []string
	var sl []int

	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	for _, linea := range contenidos {
		draftSlice = append(draftSlice, reemplazarNumerosPorPalabras(linea))
	}

	//fmt.Println(draftSlice)

	for _, line := range draftSlice {
		num := getNumbers(line)
		o, _ := getFirstAndLastNumbers(num)
		sl = append(sl, o)
	}

	//fmt.Println(sumSliceWithPrint(sl))
	fmt.Println("Suma total:", Utils.SumSlice(sl))
}

func reemplazarNumerosPorPalabras(cadena string) string {
	nuevaCadena := cadena

	for palabra, numero := range temporary {
		nuevaCadena = replaceAll(nuevaCadena, palabra, numero)
	}

	return nuevaCadena
}

// Función auxiliar para reemplazar todas las ocurrencias de una subcadena
func replaceAll(s, old, new string) string {
	for strings.Contains(s, old) {
		s = strings.Replace(s, old, new, 1)
	}
	return s
}

func getNumbers(s string) []int {
	var numeros []int

	for _, caracter := range s {
		if unicode.IsDigit(caracter) {
			numero, err := strconv.Atoi(string(caracter))
			if err == nil {
				numeros = append(numeros, numero)
			}
		}
	}
	return numeros
}

func getFirstAndLastNumbers(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("Slice vacío")
	}

	primer := nums[0]
	ultimo := nums[len(nums)-1]

	// Condición para verificar si hay un solo número en el slice
	if len(nums) == 1 {
		resultado, err := strconv.Atoi(strconv.Itoa(primer) + strconv.Itoa(primer))
		if err != nil {
			return 0, fmt.Errorf("Error al convertir a entero: %v", err)
		}
		return resultado, nil
	}

	// Concatenar el primer y último número
	resultado, err := strconv.Atoi(strconv.Itoa(primer) + strconv.Itoa(ultimo))
	if err != nil {
		return 0, fmt.Errorf("Error al convertir a entero: %v", err)
	}

	return resultado, nil
}
