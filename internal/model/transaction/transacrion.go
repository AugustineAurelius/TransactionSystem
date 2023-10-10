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

func TransactionFactory(user1ID int, user2ID int, amount float64) *Transaction {

	return &Transaction{
		ID:     rand.Intn(100000000),
		From:   user1ID,
		To:     user2ID,
		Amount: amount,
		Type:   "FromTo",
	}
}
func TransactionFactoryExchange(userID int, amount float64) *Transaction {

	return &Transaction{
		ID:     rand.Intn(100000000),
		From:   userID,
		Amount: amount,
		Type:   "Exchange",
	}
}
func TransactionFactoryReceive(userID int, amount float64) *Transaction {

	return &Transaction{
		ID:     rand.Intn(100000000),
		To:     userID,
		Amount: amount,
		Type:   "Receive",
	}
}
