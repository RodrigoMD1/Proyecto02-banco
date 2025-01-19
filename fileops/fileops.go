package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


//* para exportar este modulo se tiene que poner la primera letra en mayuscula de las funciones y despues en el archivo llamarlo en la parte de import como si fuera otro package 
// /////////////////////////////////////////////////////////////////////////
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	} // esto es para manejar errores que si hay algun error con la cifra inicial vuelva a un valor predeterminado

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil

	//* ( _ ) agregar ese guion bajo es como decirle al programa que se use eso para manejar errores temporalmente y que despues vas a poner alguna funcion o algo para manejarlo

}

/////////////////////////////////////////////////////////////////////////////////

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
} //* esto permite cuando agrego un valor a la cuenta le sumo 200 por ejemplo esa cantidad que queda del total se guarda en un archivo txt