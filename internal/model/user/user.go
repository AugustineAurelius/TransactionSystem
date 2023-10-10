package user

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          int `gorm:"primaryKey; autoIncrement"`
	Money       float64
	LockedMoney float64
}

func (user *User) TopUpYourAccount(amountOfCurrency float64) {
	user.Money = user.Money + amountOfCurrency
}

func (user *User) WithdrawCurrency(amountOfCurrency float64) error {
	if ok, err := user.BalanceCheck(amountOfCurrency); !ok {
		return fmt.Errorf("you cannot witchdraw more than you have %w", err)
	}

	user.Money = user.Money - amountOfCurrency
	return nil
}
func (user *User) BalanceCheck(amountOfCurrency float64) (bool, error) {
	if user.Money < amountOfCurrency {
		return false, fmt.Errorf("insufficient funds")
	}
	return true, nil
}
