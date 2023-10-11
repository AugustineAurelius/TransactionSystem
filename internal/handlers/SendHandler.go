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

func SendHandler(c *fiber.Ctx) error {

	IdUser1 := c.Params("user1")
	IdUser2 := c.Params("user2")
	amount := c.Params("amount")

	float, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("cannot parse float %w", err)
	}

	parseInt1, err := strconv.Atoi(IdUser1)
	if err != nil {
		return fmt.Errorf("cannot parse user1 id %w", err)
	}

	parseInt2, err := strconv.Atoi(IdUser2)
	if err != nil {
		return fmt.Errorf("cannot parse user2 id %w", err)
	}

	trans := transaction.TransactionFactory("FromTo", parseInt1, parseInt2, float)

	marshal, err := json.Marshal(trans)

	if err != nil {
		return fmt.Errorf("cannot marshal transaction")
	}

	connection.RedisVault.Set(context.Background(), fmt.Sprintf("%d", trans.ID), marshal, 5*time.Minute)

	logrus.Info("Successfully cache transaction")

	return c.Send(marshal)
}
