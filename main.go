package main

import (
	"demo/password/accountManager"
	"fmt"
	"github.com/fatih/color"
)

func main() {

Menu:
	for {
		showMenu()

		prompt := getUserPrompt("Выберите действие: ")

		switch prompt {
		case "1":
			createAccount()
		case "2":
			findAccount()
		case "3":
			deleteAccount()
		case "4":
			break Menu
		default:
			fmt.Println("Такого пункта нет -_-")
		}
	}
}

func deleteAccount() {
	url := getUserPrompt("Введите url для поиска: ")

	vault := accountManager.NewVault()
	isDeleted := vault.DeleteAccountsByUrl(url)

	if isDeleted {
		fmt.Println("Аккаунт удален")
	} else {
		color.Red("Ошибка удаления аккаунта")
	}

}

func findAccount() {
	url := getUserPrompt("Введите url для поиска: ")
	vault := accountManager.NewVault()
	result := vault.FindAccountsByUrl(url)
	if len(result) > 0 {
		for _, acc := range result {
			fmt.Printf("Сайт: %s\nЛогин: %s\nПароль: %s\n\n",
				acc.Url, acc.Login, acc.Password)
		}
	} else {
		color.Yellow("Ничего не найдено.")
	}
}

func showMenu() {
	fmt.Print("1. Создать аккаунт\n" +
		"2. Найти аккаунт\n" +
		"3. Удалить акканут\n" +
		"4. Выход\n\n")
}

func createAccount() {
	login := getUserPrompt("Введите логин: ")
	password := getUserPrompt("Введите пароль: ")
	urlString := getUserPrompt("Введите url: ")

	account, newAccountError := accountManager.NewAccount(login, password, urlString)
	if newAccountError != nil {
		fmt.Println(newAccountError)
	}

	vault := accountManager.NewVault()
	vault.AddAccount(*account)
}

func getUserPrompt(message string) string {
	fmt.Print(message)
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println("Произошла ошибка: ", err)
	}

	return input
}
