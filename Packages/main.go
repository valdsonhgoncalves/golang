package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Go run go!")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("111111")
	fmt.Println(erro)
}
