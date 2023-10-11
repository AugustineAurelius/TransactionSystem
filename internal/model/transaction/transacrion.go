package transaction

import (
	"math/rand"
)

type Transaction struct {
	ID     int     `json:"ID"`
	From   int     `json:"From"`
	To     int     `json:"To"`
	Amount float64 `json:"Amount"`
	Type   string  `json:"Type"`
}

func TransactionFactory(Type string, user1ID int, user2ID int, amount float64) *Transaction {
	if Type == "FromTo" {
		return newBasicTransaction(user1ID, user2ID, amount)
	} else if Type == "Exchange" {
		return newExchangeTransaction(user1ID, amount)
	} else {
		return newReceiveTransaction(user2ID, amount)
	}

}
func newBasicTransaction(user1ID int, user2ID int, amount float64) *Transaction {
	return &Transaction{
		ID:     rand.Intn(100000000),
		From:   user1ID,
		To:     user2ID,
		Amount: amount,
		Type:   "FromTo",
	}
}
func newExchangeTransaction(userID int, amount float64) *Transaction {

	return &Transaction{
		ID:     rand.Intn(100000000),
		From:   userID,
		Amount: amount,
		Type:   "Exchange",
	}
}
func newReceiveTransaction(userID int, amount float64) *Transaction {

	return &Transaction{
		ID:     rand.Intn(100000000),
		To:     userID,
		Amount: amount,
		Type:   "Receive",
	}
}
