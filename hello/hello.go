package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := leComando()
	nome, idade := devolveNomeEIdade()
	fmt.Println("Nome = ", nome, "- Idade = ", idade)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")

	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Não conheço esse comando")
	// }

	switch comando {
	case 1:

		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "João"
	idade := 20
	return nome, idade
}

func exibeIntroducao() {
	nome := "João"
	versao := 1.21
	fmt.Println("Olá, senhor ", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O endereço da minha variável comando é:", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := ("https://www.alura.com.br")
	resp, error := http.Get(site)
	fmt.Println(resp)
	fmt.Println(error)

}
