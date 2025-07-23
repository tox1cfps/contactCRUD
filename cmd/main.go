package main

import (
	"contactCRUD/contacts"
	"contactCRUD/database"
	"fmt"
)

func main() {
	database.Conect()
	defer database.DB.Close()

	for {
		var opcao int

		fmt.Println("\n--- Gerenciador de contatos ---")
		fmt.Print("\n1. Criar Contato")
		fmt.Print("\n2. Ver lista de contatos")
		fmt.Print("\n3. Atualizar Contatos")
		fmt.Print("\n4. Deletar Contato")
		fmt.Print("\n0. Sair")
		fmt.Print("\nSelecione a opção desejada: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 1:
			contacts.CreateContact()
		case 2:
			contacts.ListContacts()
		case 3:
			contacts.UpdateContact()
		case 4:
			contacts.DeleteContact()
		case 0:
			fmt.Println("Encerrando programa...")
			return
		}
	}
}
