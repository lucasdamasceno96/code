package getbet

import (
	"fmt"
)

// GetBet solicita a aposta do usuário e retorna a aposta inserida
func GetBet() int {
	var bet int
	fmt.Print("Enter your bet: ")

	// Lê a aposta do usuário
	_, err := fmt.Scanln(&bet)

	// Verifica se houve erro ao ler a aposta
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	// Verifica se a aposta é menor que 1
	if bet < 1 {
		fmt.Println("You must enter a bet greater than 0")
		return 0
	}

	// Exibe a aposta e retorna a aposta
	fmt.Printf("Your bet is %d\n", bet)
	return bet
}
