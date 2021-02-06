/*
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	//O go não tem while. então o for sem parametros é "loop infinito"
	for {
		exibeMenu()

		comando := leComando()

		//Exemplo de código com if
		//if comando == 1 {
		//	fmt.Println("Monitorando...")
		//} else if comando == 2 {
		//	fmt.Println("Exibindo Logs...")
		//} else if comando == 0 {
		//	fmt.Println("Saindo do programa...")
		//} else {
		//	fmt.Println("Não conheço este comando")
		//}


		switch comando {
		//O go não precisa do break no switch.
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

//Exemplo de função que devolve 2 variáveis
func devolveNome() (string, int) {
	nome := "André"
	idade := 28
	return nome, idade
}

func exibeIntroducao() {
	nome := "André"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	comandoLido := 0
	//O &(e comercial) aponta para o endereço da variável na memória
	fmt.Scanf("%d", &comandoLido)
	//Checa o endereço da variável
	fmt.Println("O endereço da sua variável é", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://www.alura.com.br",
		"https://random-status-code.herokuapp.com",
		"https://www.globo.com.br",
	}
	//funções podem retornar 2 coisas(por exmeplo o Get)
	//resp, err := http.Get(site)

	//exemplo usando um for convencional
	//for i:=0; i < len(sites); i++{
	//	fmt.Println(sites[i])
	//}

	//For do go usando range no lugar de for "convencional"
	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}
	//Se eu quiser só a resposta por ex:

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas, status", resp.StatusCode)
	}
}

//Exemplo de uso de um slice(uma abstração de um array)
//Se possivel, sempre usar um slice para arrays sem tamanho fixo
//func exibeNomes() {
//	nomes := []string{"André", "Anderson", "Lucas"}
//}
*/