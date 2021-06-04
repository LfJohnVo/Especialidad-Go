package main

import "fmt"

func main(){
	//es un contenedor asociativo
	animals := make(map[string]string)
	animals["gato"] = "🐈"
	animals["perro"] = "🐕"

	fmt.Println(animals)

	fruits := map[string]string{
		"apple": "🍎",
		"banana": "🍌",

	}

	fmt.Println(fruits)

	//eliminar elementos de map
	delete(fruits, "banana")

	fmt.Println(fruits)

	fmt.Println(animals["gato"])

	//generar validacion de si una llave no existe la crea
	if _, ok := animals["gorilla"]; !ok {
		animals["gorilla"] = "🦍"
	}

	fmt.Println(animals)

}
