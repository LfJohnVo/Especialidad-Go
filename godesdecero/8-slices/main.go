package main

import "fmt"

func main(){
	//Es un array que almacena datos de manera dinamicamente
	//Son apuntadores a array
	//esto es un array
	set := [7]string{
		"🦁",
		"🐎",
		"🐄",
		"🦋",
		"🐦",
		"🤽",
		"🗃",
	}
	//esto es un slice
	//un slice no posee datos, son apuntadores al array
	animals := set[0:5]
	fly := set[3:7]
	fly[0] = "🦅"
	fly = append(fly, "🍍")
	//len numero de elementos en el slice
	//cap numero de elementos del array donde apunta el slice, a partir del indice de donde se creo el slice
	//append agregar elementos al slice
	//make devuelve un slice de stings
	//fly := make([]string , 0 , 3)

	fmt.Println("Array: ", set)
	fmt.Println("Animales: ",animals)
	fmt.Println("Voladores: ",fly[0])
	fmt.Println("ALL: ", set[:])
	fmt.Println("Tamaño: ", len(fly))
	fmt.Println("Capacidad: ", cap(fly))
	fmt.Println("Tamaño: ", len(fly))
	fmt.Println(fly == nil)
}
