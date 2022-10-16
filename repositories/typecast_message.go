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

func (tr *TypecastMessageRepository) Create(setting ent.UserTypecastSetting, channelID string, text string) (*ent.TypecastMessage, error) {
	return tr.
		Database.
		TypecastMessage.
		Create().
		SetUserID(setting.UserID).
		SetText(text).
		SetChannelID(channelID).
		SetActorID(setting.ActorID).
		SetLang(setting.Lang).
		SetMaxSeconds(setting.MaxSeconds).
		SetNaturalness(setting.Naturalness).
		SetSpeedX(setting.SpeedX).
		SetGid(setting.Gid).
		SetStyleIdx(setting.StyleIdx).
		SetLastPitch(*setting.LastPitch).
		SetMode(*setting.Mode).
		SetPitch(setting.Pitch).
		SetStyleLabel(setting.StyleLabel).
		SetStyleLabelVersion(setting.StyleLabelVersion).
		SetTempo(setting.Tempo).
		Save(context.Background())
}
