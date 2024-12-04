package accountManager

import (
	"demo/password/encrypter"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	data = enc.Decrypt(data)

	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red(err.Error())
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		if checker(account, str) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDb) DeleteAccountsByUrl(url string) bool {
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
	bytes, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return bytes, err
}

func (vault *VaultWithDb) save() bool {
	vault.UpdatedAt = time.Now()
	bytes, err := vault.Vault.ToBytes()
	if err != nil {
		color.Red(err.Error())
		return false
	}
	bytes = vault.enc.Encrypt(bytes)
	vault.db.Write(bytes)
	return true
}
