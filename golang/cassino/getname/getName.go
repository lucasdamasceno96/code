package getname

import "fmt"

// GetName solicita o nome do usuário e retorna o nome inserido
func GetName() string {
	var name string
	fmt.Print("Welcome to Script Cassino\nEnter your name: ")

	// Lê o nome do usuário
	_, err := fmt.Scanln(&name)

	// Verifica se houve erro ao ler o nome
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Verifica se o nome está vazio
	if name == "" {
		fmt.Println("You must enter a name")
		return ""
	}

	// Exibe a saudação e retorna o nome
	fmt.Printf("Hello %s, welcome to Script Cassino\n", name)
	return name
}
