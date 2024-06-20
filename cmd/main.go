package main

import (
	"cmd/main.go/internal/application"
	postgresRepository "cmd/main.go/internal/infrastructure/db/repository"
	"cmd/main.go/internal/interfaces/router"
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db_connection_string := os.Getenv("DATASOURCE_URL")
	db, err := sql.Open("postgres", db_connection_string)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//define repositories
	userRepo := postgresRepository.NewPostgresUserRepository(db)
	messageRepo := postgresRepository.NewPostgresMessageRepository(db)

	//define service
	userService := application.NewUserService(userRepo)
	messageServicce := application.NewMessageService(messageRepo, userRepo)

	//define router
	userRouter := router.NewUserRouter(userService)
	messageRouter := router.NewMessageRouter(messageServicce)

	app := fiber.New()

	app.Post("/user", userRouter.CreateUser)
	app.Post("/message", messageRouter.SendMessage)
}
