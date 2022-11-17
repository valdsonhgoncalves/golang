package main

import "fmt"

func main() {
	//ARITMÉTICOS

	soma := 10 + 5
	subtracao := 10 - 5
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var num1 int16 = 10
	var num2 int16 = 5

	soma2 := num1 + num2

	fmt.Println(soma2)

	// ATRIBUIÇÃO
	var variavel string = "String"
	variavel2 := "String"
	fmt.Println(variavel, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

}
