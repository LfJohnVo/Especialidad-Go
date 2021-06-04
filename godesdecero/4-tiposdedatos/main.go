package main

import "fmt"

func main(){
	//bool, string, numeric,
	var a bool = true
	var b string = "EDTeam"
	//uint para numeros positivos y int para negativos y positivos pequeños
	var c uint8 = 8
	//byte es un alias de uint8
	//rune alias para int32
	//float 32 y float64 para datos con decimales
	var d float64 = 23.56
	//operaciones de variables con tipos distintos
	var e uint8 = 100
	var f uint16 = 23000
	g := uint16(e) + f
	//el operador blank (identificador en blanco) sirve para especificar que dejare una variable sin usar y que se utilizara despues
	_ = 234
	//el valor cero
	var h uint

	//%T se refiere al tipo, mientras que %v te permite mostrar el valor de la variable
	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)
	fmt.Printf("Tipo: %T, Valor: %v\n", c, c)
	fmt.Printf("Tipo: %T, Valor: %v\n", d, d)
	fmt.Printf("Tipo: %T, Valor: %v\n", g, g)
	fmt.Printf("Tipo: %T, Valor: %v\n", h, h)
}
