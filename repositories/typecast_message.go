package repositories

import (
	"context"

	"github.com/jipark0716/discordTypecast/ent"
)

type TypecastMessageRepository struct {
	Database *ent.Client
}

func NewTypecastMessageRepository(database *ent.Client) TypecastMessageRepository {
	return TypecastMessageRepository{
		Database: database,
	}
}

func (tr *TypecastMessageRepository) Create(setting *ent.UserTypecastSetting, guildID string, channelID string, text string) (*ent.TypecastMessage, error) {
	create := tr.
		Database.
		TypecastMessage.
		Create().
		SetUserID(setting.UserID).
		SetText(text).
		SetGuildID(guildID).
		SetChannelID(channelID).
		SetActorID(setting.ActorID).
		SetLang(setting.Lang).
		SetMaxSeconds(setting.MaxSeconds).
		SetNaturalness(setting.Naturalness).
		SetSpeedX(setting.SpeedX).
		SetGid(setting.Gid).
		SetStyleIdx(setting.StyleIdx).
		SetPitch(setting.Pitch).
		SetStyleLabel(setting.StyleLabel).
		SetStyleLabelVersion(setting.StyleLabelVersion).
		SetTempo(setting.Tempo)

	if setting.LastPitch != nil {
		create.SetLastPitch(*setting.LastPitch)
	}

	if setting.Mode != nil {
		create.SetMode(*setting.Mode)
	}

	return create.Save(context.Background())
}
