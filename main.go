package main

import (
	"bufio"
	// Fornece ferramentas para leitura e escrita com buffer
	// Usado para ler a entrada do usuário pelo terminal

	"fmt"

	"math/rand/v2"
	// Gera números aleatórios

	"os"
	// Fornece acesso a recursos do sistema operacional
	// Usado para acessar a entrada do terminal (os.Stdin)

	"strconv"
	// Converte strings para outros tipos (int, float, etc)

	"strings"
	// Manipulação de strings
)

func main() {
	fmt.Println("Jogo da Adivinhação!")
	fmt.Println("Um número aleatório entre 0 a 100 será sorteado! Tente acertar")

	// Gera um número aleatório entre 0 e 100 (inclusive)
	x := rand.Int64N(101)

	// Cria um scanner para ler a entrada do usuário pelo terminal
	scanner := bufio.NewScanner(os.Stdin)

	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é seu palpite?")

		// Lê a entrada do usuário
		scanner.Scan()
		chute := scanner.Text()

		// Remove espaços em branco do início e fim da string
		chute = strings.TrimSpace(chute)

		// Converte o chute de string para int64
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Seu chute deve ser um número inteiro")
			return
		}

		switch {
		case chuteInt > x:
			fmt.Println("Chute muito alto! O número é menor que", chuteInt)

		case chuteInt < x:
			fmt.Println("Chute muito baixo! O número é maior que", chuteInt)

		case chuteInt == x:
			fmt.Println("Chute certeiro!")
			fmt.Println("Seus chutes foram:", chutes[:i])
			return
		}

		// Armazena o chute no array
		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Você perdeu! Tente novamente\n"+
			"Você teve 10 tentativas e chutou esses números: %v\n"+
			"O número era: %d",
		chutes,x,
	)
}
