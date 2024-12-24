package chats

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henmalib/messenger/cmd/db/repos"
	"github.com/henmalib/messenger/cmd/utils"
)

type ChatStore interface {
	GetChatsByUserId(user_id int) ([]repos.Chat, error)
}

type Handler struct {
	store ChatStore
}

func CreateHandler(store ChatStore) Handler {
	return Handler{
		store,
	}
}

func (handler *Handler) GetChats(ctx *gin.Context) {
	// TODO: get id from jwt token
	chats, err := handler.store.GetChatsByUserId(1)

	if err != nil {
		utils.WriteError(ctx, err)
		return
	}

	utils.WriteResponse(ctx, http.StatusOK, gin.H{
		"chats": chats,
	})
}
