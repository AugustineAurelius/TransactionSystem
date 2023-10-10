package validator

import (
	"TransactionSystem/internal/connection"
	"TransactionSystem/internal/model/transaction"
	"TransactionSystem/internal/model/user"
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
)

type Validate interface {
	DoValidate()
}

type Validator struct {
	counter int64
}

func NewValidator() *Validator {
	return &Validator{counter: 0}
}

func (v *Validator) DoValidate() error {

	keys := connection.RedisVault.Keys(context.Background(), "*")

	result, err := keys.Result()
	if err != nil {
		return fmt.Errorf("cannot get keys from cache")
	}

	if len(result) != 0 {
		var trans transaction.Transaction

		marsheldTrans := connection.RedisVault.Get(context.Background(), result[0])

		err = json.Unmarshal([]byte(marsheldTrans.Val()), &trans)
		if err != nil {
			return fmt.Errorf("cannot marshal transaction from cache")
		}

		user1 := new(user.User)
		user2 := new(user.User)
		if trans.Type == "FromTo" {
			connection.DB.First(&user1, trans.From)
			connection.DB.First(&user2, trans.To)
			err = v.MoveMoney(user1, user2, trans.Amount)
			if err != nil {
				return fmt.Errorf("cannot move money %w", err)
			}
		} else if trans.Type == "Exchange" {
			connection.DB.First(&user1, trans.From)
			err = v.MoveMoneyToExchanger(user1, trans.Amount)
			if err != nil {
				return fmt.Errorf("cannot move money %w", err)
			}
		} else {

			connection.DB.First(&user2, trans.To)
			err = v.MoveMoneyReceive(user2, trans.Amount)
			if err != nil {
				return fmt.Errorf("cannot move money %w", err)
			}
		}

		connection.RedisVault.Del(context.Background(), result[0])
	}

	return nil
}

func (v *Validator) MoveMoney(user1 *user.User, user2 *user.User, amount float64) error {
	if ok, err := user1.BalanceCheck(amount); ok {
		defer connection.DB.Save(&user1)
		defer connection.DB.Save(&user2)
		user1.Money = user1.Money - amount
		user2.Money = user2.Money + amount
		atomic.AddInt64(&v.counter, 1)
	} else {
		return fmt.Errorf("cannot move money %w", err)
	}
	return nil
}

func (v *Validator) MoveMoneyToExchanger(user *user.User, amount float64) error {
	if ok, err := user.BalanceCheck(amount); ok {
		defer connection.DB.Save(&user)
		user.Money = user.Money - amount
		atomic.AddInt64(&v.counter, 1)
	} else {
		return fmt.Errorf("cannot move money %w", err)
	}
	return nil
}
func (v *Validator) MoveMoneyReceive(user *user.User, amount float64) error {
	defer connection.DB.Save(&user)
	user.Money = user.Money + amount
	atomic.AddInt64(&v.counter, 1)
	return nil
}

func (v *Validator) CounterOfTransactions() int64 {
	return atomic.LoadInt64(&v.counter)
}
