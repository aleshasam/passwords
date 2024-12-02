package accountManager

import (
	"encoding/json"
	"errors"
	"github.com/fatih/color"
	"math/rand/v2"
	"net/url"
	"time"
)

var lettersRunes = []rune("abcdefABCDEF1234")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) OutputPassword() {
	color.Red(acc.Password)
}

func (acc Account) ToBytes() ([]byte, error) {
	bytes, err := json.MarshalIndent(acc, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("Empty login")
	}

	_, urlParseError := url.ParseRequestURI(urlString)
	if urlParseError != nil {
		return nil, errors.New("Invalid url")
	}

	acc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		acc.generatePassword(12)
	}

	return acc, nil
}

func (acc *Account) generatePassword(n int) {
	password := make([]rune, n)

	for i := range password {
		password[i] = lettersRunes[rand.IntN(len(lettersRunes))]
	}

	acc.Password = string(password)
}
