package main

import "fmt"

func main(){
	//Operadores aritmeticos (),*,/,%,+,-
	var a = 4 + 2 * 3
	fmt.Println(a)
	//Operadores de asignación: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	fmt.Println(b)
	//declaración post-incremento y post-decremento: ++, --
	//(no son una expresion sino una declaración
	var c = 3
	//Equivalente a c++ o c--
	c++
	fmt.Println(c)
	//Operadores comparación >, <, ==, !=, >= <=
	fmt.Println(4 > 5)
	//Operadores lógicos &&, ||
	var age = 30
	fmt.Println(age >= 18 && age <= 60)
	//Unario !
	fmt.Println(!(4 == 4))
}
