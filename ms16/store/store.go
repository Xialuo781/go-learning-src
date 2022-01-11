package store

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

// NOTE a要设置成一个指针
func (a *Account) ChangeName(newname string) {
	a.FirstName = newname
}

type Employee struct {
	Account
	Credits float64
}

func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("invalid credit amount")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("you can't remove more credits than the account has")
	}
	return 0.0, errors.New("you can't remove negative numbers")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

// 为什么这里的函数声明不需要加*，猜测也许是因为这里是在实现接口，必须按照接口的格式来声明
func (e Employee) String() string {
	return fmt.Sprintf("Name %s %s\nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(FirstName, LastName string, credits float64) (*Employee, error) {
	return &Employee{Account{FirstName, LastName}, credits}, nil
}
