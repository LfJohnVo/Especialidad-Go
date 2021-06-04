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
	fmt.Println("Array: ", set)
	fmt.Println("Animales: ",animals)
	fmt.Println("Voladores: ",fly[0])
	fmt.Println("ALL: ", set[:])
}
