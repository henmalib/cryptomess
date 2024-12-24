package main

import (
	"log"

	"github.com/henmalib/messenger/cmd/api"
	"github.com/henmalib/messenger/cmd/api/chats"
	"github.com/henmalib/messenger/cmd/api/users"
	"github.com/henmalib/messenger/cmd/db"
	"github.com/henmalib/messenger/cmd/db/repos"
)

type User struct {
	Id        int
	Username  string
	CreatedAt int
}

func main() {
	db, err := db.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	server := api.InitServer()

	chatRepo := repos.CreateChatRepo(db)

	chatHandler := chats.CreateHandler(chatRepo)
	server.GET("/chats", chatHandler.GetChats)

	userRepo := repos.CreateUserRepo(db)
	usersHandler := users.CreateHandler(userRepo)
	server.POST("/register", usersHandler.RegisterUser)

	if err = server.Run(); err != nil {
		log.Fatalln(err)
	}
}
