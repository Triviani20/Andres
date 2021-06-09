package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	//fmt.Printf("El contenido del archivo es: %s", contenido)

	//fileCoin := "moneda.txt"
	var fileCoin string
	var fileParams string

	printCoin(readCoin(fileCoin))
	printParams(readParams(fileParams))

}

func readCoin(filename string) []byte {

	fileName := "moneda.txt"
	readBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	return readBytes
}

func printCoin(p []byte) {
	var j Coin

	err1 := json.Unmarshal(p, &j)

	if err1 != nil {

		fmt.Printf("Error decodificando: %v\n", err1)
	} else {
		fmt.Printf("TypeCoin : %s\n", j.TypeCoin)
		fmt.Printf("Value: %d\n", j.Value)

	}

}

func readParams(filename string) []byte {

	fileName := "parametros.txt"
	readBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %v", err)
	}
	return readBytes
}

func printParams(p []byte) {

	var j Parametros

	err1 := json.Unmarshal(p, &j)

	if err1 != nil {

		fmt.Printf("Error decodificando: %v\n", err1)
	} else {
		fmt.Printf("Number_slots_left : %d\n", j.Number_slots_left)
		fmt.Printf("Number_slots_right: %d\n", j.Number_slots_right)
		fmt.Printf("Strike_gap: %d\n", j.Strike_gap)
		fmt.Printf("Interval_frame: %d\n", j.Interval_frame)
		fmt.Printf("Asserts: %v\n", j.Asserts)
	}
}
