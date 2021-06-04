package main

import "fmt"

func main(){
	fruit := "🍎"
	//Para declarar un punteros
	var p *string
	p = &fruit

	fmt.Printf("Tipo: %T, Valor: %s, Direccion de memoria: %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s\n ", p,p, *p)
}