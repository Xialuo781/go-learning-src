package bank

import (
	"errors"
	"fmt"
)

// 接口类
// 在本例中，Account类和CustomAccount类都实现了接口Statement
// 也许可以这样理解：Account和CustomAccount都是Bank的子类
type Bank interface {
	Statement() string
}

// 全局函数，可以传Bank及实现了Statement的“子类”对象
func Statement(b Bank) string {
	return b.Statement()
}

type Customer struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Phone   string `json:"Phone"`
}

type Account struct {
	Customer
	Number  int32   `json:"Number"`
	Balance float64 `json:"Balance"`
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if amount > a.Balance {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to transfer should be greater than the account's balance")
	}

	a.Withdraw(amount)
	dest.Deposit(amount)
	return nil
}
