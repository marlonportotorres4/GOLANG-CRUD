package main

import (
	"errors"
	"fmt"
)

type motorista struct {
	porcentagemKm float64
	funcionario
}

type vendedor struct {
	porcentagemVenda float64
	funcionario
}

type funcionario struct {
	nome     string
	idade    uint8
	salario  float64
	cpf      string
	endereco endereco
}

type endereco struct {
	numCasa    uint16
	cep        string
	
}

func main() {
	var opcao uint8
	
	
	fmt.Println("------------ Bem vindo ao CRUD de funcionários ------------")

	fmt.Println("\n1 -> Criar Funcionário")
	fmt.Println("2 -> Visualizar Funcionários")
	fmt.Println("3 -> Atualizar lista de Funcionários")
	fmt.Println("4 -> Deletar Funcionário")
	fmt.Println("5 -> Sair do CRUD")
	
	fmt.Println("\nDigite o número da operação desejada:")
	
	_, err := fmt.Scan(&opcao) 
	
	if err != nil {
		msgErro := errors.New("Entrada inválida, por favor, digite um número inteiro.")
		fmt.Println("Erro ao ler a opção:", msgErro)
	} else {
	
	switch opcao {
	case 1:
		fmt.Println("Você escolheu: Criar Funcionário")
		criarFunc()
		return
	case 2:
		fmt.Println("Você escolheu: Visualizar Funcionários")
		visualizarFunc()
	case 3:
		fmt.Println("Você escolheu: Atualizar lista de Funcionários")
		atualizarFunc()
	case 4:
		fmt.Println("Você escolheu: Deletar Funcionários")
		deletarFunc()
	case 5:
		fmt.Println("Saindo do CRUD...")
		return
	default:
		fmt.Println("Opção inválida: Digite um numero válido para continuar")
		
	}
}

}




func criarFunc()  {       //POSSO USAR UM SLICE COM APPEND
	fmt.Println("Lógica para criar um funcionário")
	
	




}

func visualizarFunc()  {  //SLICE PARA VISUALIZAR
	fmt.Println("Lógica para criar um funcionário")
	
}

func atualizarFunc()  {  //ATUALIZAR COM SLICE
	fmt.Println("Lógica para criar um funcionário")
	
}

func deletarFunc()  {  //EXCLUIR DENTRO DE UM SLICE
	fmt.Println("Lógica para criar um funcionário")
	
}