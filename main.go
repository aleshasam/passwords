package main

import (
	"demo/password/accountManager"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"strings"
)

var menu = map[string]func(*accountManager.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить акканут",
	"5. Выход",
	"Выберите действие",
}

func main() {
	color.Yellow("*** Менеджер паролей ***")
	errEnv := godotenv.Load()
	if errEnv != nil {
		output.PrintError(errEnv.Error())
	}
	vault := accountManager.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())

Menu:
	for {
		prompt := getUserPrompt(menuVariants...)
		menuFunc := menu[prompt]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)

	}
}

func deleteAccount(vault *accountManager.VaultWithDb) {
	url := getUserPrompt("Введите url для поиска")

	isDeleted := vault.DeleteAccountsByUrl(url)

	if isDeleted {
		fmt.Println("Аккаунт удален")
	} else {
		output.PrintError("Ошибка удаления аккаунта")
	}

}

func findAccountByUrl(vault *accountManager.VaultWithDb) {
	url := getUserPrompt("Введите url для поиска")
	result := vault.FindAccounts(url, func(account accountManager.Account, s string) bool {
		return strings.Contains(account.Url, url)
	})
	if len(result) > 0 {
		for _, acc := range result {
			fmt.Printf("Сайт: %s\nЛогин: %s\nПароль: %s\n\n",
				acc.Url, acc.Login, acc.Password)
		}
	} else {
		color.Yellow("Ничего не найдено.")
	}
}

func findAccountByLogin(vault *accountManager.VaultWithDb) {
	login := getUserPrompt("Введите логин для поиска")
	result := vault.FindAccounts(login, func(account accountManager.Account, s string) bool {
		return strings.Contains(account.Login, login)
	})
	if len(result) > 0 {
		for _, acc := range result {
			fmt.Printf("Сайт: %s\nЛогин: %s\nПароль: %s\n\n",
				acc.Url, acc.Login, acc.Password)
		}
	} else {
		color.Yellow("Ничего не найдено.")
	}
}

func createAccount(vault *accountManager.VaultWithDb) {
	login := getUserPrompt("Введите логин")
	password := getUserPrompt("Введите пароль")
	urlString := getUserPrompt("Введите url")

	account, newAccountError := accountManager.NewAccount(login, password, urlString)
	if newAccountError != nil {
		output.PrintError(newAccountError.Error())
	} else {
		vault.AddAccount(*account)
	}
}

func getUserPrompt(messages ...string) string {

	for i := 0; i < len(messages); i++ {
		fmt.Print(messages[i])
		if len(messages) == i+1 {
			fmt.Print(": ")
		} else {
			fmt.Print("\n")
		}
	}

	var input string
	_, err := fmt.Scanln(&input)

	if err != nil && err.Error() != "unexpected newline" {
		output.PrintError(err.Error())
	}

	return input
}
