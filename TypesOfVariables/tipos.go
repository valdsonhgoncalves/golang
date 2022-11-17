package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -100
	fmt.Println(numero)

	//alias
	//int32 = rune
	var numero3 rune = 123556
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 122
	fmt.Println(numero4)

	var texto int16
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Internal error")
	fmt.Println(erro)

}
