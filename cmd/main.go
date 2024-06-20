package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"whatsapp-like/internal/application"
	postgresRepository "whatsapp-like/internal/infrastructure/db/repository"
	"whatsapp-like/internal/interfaces/router"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
