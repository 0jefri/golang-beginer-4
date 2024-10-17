package model

import (
	"errors"
)

type UserCreate interface {
	CreateUser(id string, name string, email string, saldo float64) (Account, error)
}

type Balance struct {
	Amount float64
}

type Account struct {
	Id    string
	Name  string
	Email string
	Saldo Balance
}

func (a Account) CreateUser(id string, name string, email string, saldo float64) (Account, error) {
	if id == "" {
		return Account{}, errors.New("ID tidak boleh kosong")
	}
	if name == "" {
		return Account{}, errors.New("nama tidak boleh kosong")
	}
	if email == "" {
		return Account{}, errors.New("email tidak boleh kosong")
	}
	if saldo < 0 {
		return Account{}, errors.New("saldo tidak boleh negatif")
	}

	return Account{
		Id:    id,
		Name:  name,
		Email: email,
		Saldo: Balance{Amount: saldo},
	}, nil
}

func (a *Account) TambahSaldo(jumlah float64) error {
	if jumlah < 0 {
		return errors.New("jumlah yang ditambahkan tidak boleh negatif")
	}
	a.Saldo.Amount += jumlah
	return nil
}

func (a *Account) KurangiSaldo(jumlah float64) error {
	if jumlah < 0 {
		return errors.New("jumlah yang dikurangi tidak boleh negatif")
	}
	if jumlah > a.Saldo.Amount {
		return errors.New("saldo tidak mencukupi")
	}
	a.Saldo.Amount -= jumlah
	return nil
}
