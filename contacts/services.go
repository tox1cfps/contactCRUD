package contacts

import (
	"contactCRUD/database"
	"fmt"
)

func CreateContact() {
	var name, email, phone string

	fmt.Print("\nDigite o nome do contato: ")
	fmt.Scan(&name)

	fmt.Print("\nDigite o telefone do contato: ")
	fmt.Scan(&phone)

	fmt.Print("\nDigite o email do contato: ")
	fmt.Scan(&email)

	_, err := database.DB.Exec("INSERT INTO contacts (nome, telefone, email) VALUES ($1, $2, $3)", name, phone, email)
	if err != nil {
		fmt.Println("Erro ao Cadastrar contato:", err)
		return
	}
	fmt.Println("Contato cadastrado com sucesso!")
}

func ListContacts() {
	rows, err := database.DB.Query("SELECT id, nome, telefone, email FROM contacts")
	if err != nil {
		fmt.Println("Erro ao Listar contatos:", err)
		return
	}

	defer rows.Close()

	fmt.Println("\n### Lista de Contatos ###")

	for rows.Next() {
		var id int
		var nome, telefone, email string

		err := rows.Scan(&id, &nome, &telefone, &email)
		if err != nil {
			fmt.Println("Erro ao ler contatos:", err)
			continue
		}
		fmt.Printf("\nID: %d | Nome: %s | Telefone: %s | Email: %s |", id, nome, telefone, email)

		if err = rows.Err(); err != nil {
			fmt.Println("Erro ao Iterar pelos contatos:", err)
			return
		}
	}
}

func UpdateContact() {
	ListContacts()

	var id int
	fmt.Print("\nDigite o ID do contato que deseja atualizar: ")
	fmt.Scan(&id)

	var novoNome string
	var novoTelefone string
	var novoEmail string

	fmt.Println("Digite o novo nome: ")
	fmt.Scan(&novoNome)

	fmt.Println("Digite o novo telefone: ")
	fmt.Scan(&novoTelefone)

	fmt.Println("Digite o novo email: ")
	fmt.Scan(&novoEmail)

	result, err := database.DB.Exec("UPDATE contacts SET nome=$1, telefone=$2, email=$3 WHERE id=$4", novoNome, novoTelefone, novoEmail, id)
	if err != nil {
		fmt.Println("Erro ao atualizar o contato:", err)
		return
	}

	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		fmt.Println("Contato não encontrado")
		return
	}

	fmt.Println("Contato atualizado com sucesso!")
}

func DeleteContact() {
	ListContacts()

	var id int
	fmt.Print("\nDigite o ID do usuário que deseja deletar: ")
	fmt.Scan(&id)

	result, err := database.DB.Exec("DELETE FROM contacts WHERE id=$1", id)
	if err != nil {
		fmt.Println("Erro ao deletar contato:", err)
		return
	}
	rAffected, _ := result.RowsAffected()
	if rAffected == 0 {
		fmt.Println("Contato não encontrado")
		return
	}
	fmt.Println("Contato deletado com sucesso!")
}
