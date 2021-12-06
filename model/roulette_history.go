package model

import (
	"time"
	"strings"
	"github.com/bwmarrin/discordgo"
)

func CreateRouletteHistory(members []*discordgo.Member, result string, GuildID string) (err error) {
	memberIDs := []string{}
	for _, member := range members {
		memberIDs = append(memberIDs, member.User.ID)
	}
	_, err = Insert(map[string]interface{}{
		"guild_id": GuildID,
		"members": strings.Join(memberIDs, ","),
		"result": result,
		"created_at": time.Now().Format("2006-01-02 15:04:05"),
	}, "roulette_history")
	return
}
