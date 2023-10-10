package handlers

import (
	"TransactionSystem/internal/connection"
	"TransactionSystem/internal/model/transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func ReceiveHandler(c *fiber.Ctx) error {
	IdUser := c.Params("user")
	amount := c.Params("amount")

	float, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("cannot parse float %w", err)
	}

	parseInt, err := strconv.Atoi(IdUser)
	if err != nil {
		return fmt.Errorf("cannot parse user1 id %w", err)
	}

	trans := transaction.TransactionFactoryReceive(parseInt, float)

	marshal, err := json.Marshal(trans)

	if err != nil {
		return fmt.Errorf("cannot marshal transaction")
	}

	connection.RedisVault.Set(context.Background(), fmt.Sprintf("%d", trans.ID), marshal, 5*time.Minute)

	logrus.Info("Successfully cache transaction")

	return c.Send(marshal)
}
