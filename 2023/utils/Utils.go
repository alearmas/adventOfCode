package Utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadTXTFile(filename string) ([]string, error) {
	// Leer el contenido completo del archivo
	contenido, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Convertir el contenido a un slice de strings, dividiendo por líneas
	lineas := strings.Split(string(contenido), "\n")

	// Eliminar la última línea si está vacía (útil si el archivo termina con un salto de línea)
	if len(lineas) > 0 && lineas[len(lineas)-1] == "" {
		lineas = lineas[:len(lineas)-1]
	}

	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return nil, nil
	}

	return lineas, nil
}

func SumSlice(nums []int) int {
	s := 0
	for _, value := range nums {
		s += value
	}
	return s
}
