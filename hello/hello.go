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
	resp, _ := http.Get(site)
	fmt.Println(resp)

}

func devolveNomeEIdade() (string, int) {
	nome := "João"
	idade := 10
	return nome, idade
}

func devolveNomeEIdade2() (nome string, ano int) {
	nome = "João"
	fmt.Scanln(&ano)
	ano = 2023 - ano
	return
}
