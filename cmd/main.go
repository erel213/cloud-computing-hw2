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
	groupRepo := postgresRepository.NewPostgresGroupRepostiroy(db)

	//define service
	userService := application.NewUserService(userRepo)
	messageServicce := application.NewMessageService(messageRepo, userRepo, groupRepo)
	groupService := application.NewGroupService(groupRepo, userRepo)

	//define router
	userRouter := router.NewUserRouter(userService)
	messageRouter := router.NewMessageRouter(messageServicce)
	groupRouter := router.NewGroupRouter(groupService)

	app := fiber.New()
	//User controller
	app.Post("/user", userRouter.CreateUser)
	app.Post("/user/block", userRouter.BlockUser)
	app.Get("/user/:userId", userRouter.GetUserById)

	//Message controller
	app.Post("/message", messageRouter.SendMessage)
	app.Get("/message/user/:userId", messageRouter.GetMessagesForUser)

	//Group controller
	app.Post("/group", groupRouter.CreateGroup)
	app.Post("/group/user", groupRouter.AddUserToGroup)
	app.Delete("/group/:groupID/user/:userID", groupRouter.RemoveUserFromGroup)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
