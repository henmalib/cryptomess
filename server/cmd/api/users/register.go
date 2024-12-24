package users

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/henmalib/messenger/cmd/db/repos"
	"github.com/henmalib/messenger/cmd/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	CreateUser(username, hash string) (repos.User, error)
}

type Handler struct {
	store UserStore
}

func CreateHandler(store UserStore) Handler {
	return Handler{
		store,
	}
}

type registerBody struct {
	Username string `validate:"required,min=3,max=64" json:"username"`
	Password string `validate:"required,min=8,max=64" json:"password"`
}

func (handler *Handler) RegisterUser(ctx *gin.Context) {
	var body registerBody
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		utils.WriteError(ctx, err)
		return
	}

	if err = utils.Validate.Struct(body); err != nil {
		utils.WriteError(ctx, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}

	user, err := handler.store.CreateUser(body.Username, string(hash))
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}

	// TODO: generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		utils.WriteError(ctx, err)
		return
	}

	utils.WriteResponse(ctx, http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})

}
