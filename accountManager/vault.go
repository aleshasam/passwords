package accountManager

import (
	"demo/password/files"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	data, err := files.Read("accounts.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red(err.Error())
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account

	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *Vault) DeleteAccountsByUrl(url string) bool {
	var accounts []Account

	for _, account := range vault.Accounts {
		if !strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}

	vault.Accounts = accounts

	vault.save()
	return true
}

func (vault Vault) ToBytes() ([]byte, error) {
	bytes, err := json.MarshalIndent(vault, "", "    ")
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func (vault *Vault) save() bool {

	vault.UpdatedAt = time.Now()
	bytes, err := vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
		return false
	}
	files.Write(bytes, "accounts.json")

	return true
}
