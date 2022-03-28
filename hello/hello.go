package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func monitorando() {
	fmt.Println("Monitorando... ")
	// site := "https://www.alura.com.br"https://random-status-code.herokuapp.com
	// sites := []string{"", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	for volta := 0; volta < monitoramentos; volta++ {
		for i, site := range sites {
			fmt.Println("Volta: ", volta, "Testando site: ", i, site)

			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println()
	}

	fmt.Println()
}

func mostraMenu() {
	fmt.Println("0 - Sair")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println()
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func main() {
	// leSitesDoArquivo()
	// registraLog("site-falso", false)
	// imprimeLogs()

	for {
		mostraMenu()

		comando := leComando()

		switch comando {
		case 0:
			os.Exit(0)
		case 1:
			monitorando()
		case 2:
			fmt.Println("Exibindo logs")
			imprimeLogs()
		default:
			fmt.Println("*-* Escolha uma opção válida! *-*")

		} //END SWITCH
		fmt.Println()
	}

	// fmt.Println("Programa encerrado!")

}

func exibeNomes() { //TRABALHANDO COM SLICES
	nomes := []string{"jose", "Manoel", "Antonio", "Maria", "Joaquina"}
	nomes = append(nomes, "Novo Nome")
	fmt.Println(nomes, "tem ", len(nomes), "itens e ", cap(nomes), "capacidade!")
}

func testaSite(site string) {

	st, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	if st.StatusCode == 200 {
		fmt.Println("Site: ", site, " foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "está com problema. Status Code: ", st.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("hello/sites.txt")
	// arquivo, err := ioutil.ReadFile(("hello/sites.txt"))
	// arquivo, err := ioutil.ReadFile(("hello/sites.txt"))

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF { //ERRO DE END OF FILE, OU SEJA, ARQUIVO ACABOU
			// fmt.Println("Erro: ", err)
			break
		}

	}
	arquivo.Close()
	// fmt.Println(sites)

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("hello/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - Nome do site: " + site + " Status: " + strconv.FormatBool(status))

	fmt.Fprintln(arquivo)

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("hello/log.txt")

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(string(arquivo))

}
