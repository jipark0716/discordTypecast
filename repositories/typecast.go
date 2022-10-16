package repositories

import (
	"github.com/jipark0716/discordTypecast/config"
	"github.com/jipark0716/discordTypecast/ent"
	"github.com/jipark0716/typecastgo"
)

type Typecast struct {
	typecastgo.Session
	Actors []typecastgo.TypecastActor
}

func NewTypecast() (typecast Typecast) {
	typecast.Session = typecastgo.NewSession()
	typecast.Connect(&typecastgo.LoginRequest{
		Email:             config.Get("TYPECAST_EMAIL").(string),
		Password:          config.Get("TYPECAST_PASSWORD").(string),
		ReturnSecureToken: true,
	})
	return
}

func (t *Typecast) GetActors() (actors []typecastgo.TypecastActor, err error) {
	if t.Actors == nil {
		actors, err = t.Session.GetActors()
		if err != nil {
			return
		}
		t.Actors = actors
	}
	actors = t.Actors
	return
}

func (t *Typecast) Do(setting *ent.UserTypecastSetting, text string) ([]byte, error) {
	request := &typecastgo.TypecastExecuteRequest{
		ActorId:           setting.ActorID,
		Text:              text,
		Lang:              setting.Lang,
		MaxSeconds:        setting.MaxSeconds,
		Naturalness:       setting.Naturalness,
		SpeedX:            setting.SpeedX,
		Gid:               setting.Gid,
		StyleIdx:          setting.StyleIdx,
		LastPitch:         setting.LastPitch,
		Pitch:             setting.Pitch,
		StyleLabel:        setting.StyleLabel,
		StyleLabelVersion: setting.StyleLabelVersion,
		Tempo:             setting.Tempo,
	}

	if setting.Mode != nil {
		request.Mode = *setting.Mode
	}

	return t.Session.Do([]*typecastgo.TypecastExecuteRequest{request})
}
