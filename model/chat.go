package model

import (
	// "fmt"
	"time"
)

type Chat struct {
	ID int64 `db:"id"`
	MemberID string `db:"member_id"`
	GuildID string `db:"guild_id"`
	ChannelID string `db:"channel_id"`
	Content string `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

func NewChat() Chat {
	return Chat{
		CreatedAt: time.Now(),
	}
}

func (this *Chat) Save() (err error) {
	result, err := Insert(map[string]interface{}{
		"member_id": this.MemberID,
		"guild_id": this.GuildID,
		"channel_id": this.ChannelID,
		"content": this.Content,
		"created_at": this.CreatedAt.Format("2006-01-02 15:04:05"),
	}, "typecast_chats")

	this.ID, _ = result.LastInsertId()

	return
}
