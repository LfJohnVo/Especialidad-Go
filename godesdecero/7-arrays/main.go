package main

import "fmt"

func main(){
	//array de 3 posiciones y tipo de datos en string
	//en un array es de tamaΓ±o fijo
	/*var food [3]string
	food[0] = "π―"
	food[1] = "π"
	food[2] = "π­"*/
	//Otra manera de asignar el array
	food := [3]string{
		"π―",
		"π",
		"π­",
	}

	fmt.Println(food)
}
