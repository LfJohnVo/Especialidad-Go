package main

import "fmt"

func main(){
	//array de 3 posiciones y tipo de datos en string
	//en un array es de tamaño fijo
	/*var food [3]string
	food[0] = "🌯"
	food[1] = "🍕"
	food[2] = "🌭"*/
	//Otra manera de asignar el array
	food := [3]string{
		"🌯",
		"🍕",
		"🌭",
	}

	fmt.Println(food)
}
