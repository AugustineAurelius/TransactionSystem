package main

import (
	"TransactionSystem/internal/config"
	"TransactionSystem/internal/connection"
	"TransactionSystem/internal/handlers"
	"TransactionSystem/internal/model/user"
	"TransactionSystem/internal/model/validator"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

var server = fiber.New()
var cfg *config.Config

func init() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Error in init func :%s", err)
	}
	connection.ConnectionToDatabase(cfg)
	connection.InitializeCache(cfg)
	server.Post("/create_user", handlers.CreateHandler)
	server.Get("/send/:user1/:user2/:amount", handlers.SendHandler)
	server.Get("/exchange/:user/:amount", handlers.ExchangeHandler)
	server.Get("/receive/:user/:amount", handlers.ReceiveHandler)
	logrus.Info("Successfully init main func")
}

func main() {
	err := connection.DB.AutoMigrate(&user.User{})
	if err != nil {
		logrus.Fatal("cannot migrate")
	}

	go func() {
		err = server.Listen(cfg.Port) // инизиализируем сервер
		if err != nil {
			logrus.Fatal("Could not to start listen on port")
		}
	}()

	val := validator.NewValidator()
	for {
		time.Sleep(1 * time.Second)
		go func() {
			err := val.DoValidate()
			if err != nil {
				logrus.Errorf("can't validate because %s", err)
			}
		}()
		fmt.Println(val.CounterOfTransactions())

	}

}
