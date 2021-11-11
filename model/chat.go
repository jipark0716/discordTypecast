package model

import (
	"fmt"
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

type ChatMemberCount struct {
	MemberID string `db:"member_id"`
	Count int `db:"cnt"`
}

func GuildsID() (result []string, err error) {
	err = Get(&result, "SELECT DISTINCT `guild_id` FROM `typecast_chats`")
	return
}

func ChatCount(guildID string) (result []ChatMemberCount, err error) {
	sql := fmt.Sprintf("SELECT COUNT(*) cnt, `member_id` FROM `typecast_chats` WHERE `guild_id` = '%s' GROUP BY `member_id` ORDER BY cnt desc", guildID)
	err = Get(&result, sql)
	return
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
