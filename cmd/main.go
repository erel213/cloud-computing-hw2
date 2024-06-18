package main

import (
	"cmd/main.go/internal/application"
	postgres_user_repository "cmd/main.go/internal/infrastructure/db/repository"
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
	userRepo := postgres_user_repository.NewPostgresUserRepository(db)

	//define service
	userService := application.NewUserService(userRepo)

	//define router
	userRouter := router.NewUserRouter(userService)

	app := fiber.New()

	app.Post("/user", userRouter.CreateUser)
}
