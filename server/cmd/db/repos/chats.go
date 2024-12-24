package repos

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type ChatRepo struct {
	DB *sqlx.DB
}

type Chat struct {
	Id    int
	Title string

	CreatedAt time.Time `db:"created_at"`
}

func CreateChatRepo(db *sqlx.DB) ChatRepo {
	return ChatRepo{
		DB: db,
	}
}

func (repo ChatRepo) GetChatsByUserId(user_id int) ([]Chat, error) {
	chats := []Chat{}
	err := repo.DB.Select(&chats, `SELECT chats.* from chats
		JOIN chat_users ON chat_users.chat_id = chats.id
		WHERE chat_users.user_id = $1;`, user_id)

	return chats, err
}
