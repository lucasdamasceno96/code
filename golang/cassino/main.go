package main

import (
	"cassino/getname" // Importando o pacote getname corretamente
	"fmt"
)

func main() {
	name := getname.GetName() // Chama a função GetName do pacote getname
	if name != "" {
		fmt.Println("Let's start to play:", name)
	}
}
