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

func (t *Typecast) Do(message *ent.TypecastMessage) ([]byte, error) {
	request := &typecastgo.TypecastExecuteRequest{
		ActorId:           message.ActorID,
		Text:              message.Text,
		Lang:              message.Lang,
		MaxSeconds:        message.MaxSeconds,
		Naturalness:       message.Naturalness,
		SpeedX:            message.SpeedX,
		Gid:               message.Gid,
		StyleIdx:          message.StyleIdx,
		LastPitch:         message.LastPitch,
		Pitch:             message.Pitch,
		StyleLabel:        message.StyleLabel,
		StyleLabelVersion: message.StyleLabelVersion,
		Tempo:             message.Tempo,
	}

	if message.Mode != nil {
		request.Mode = *message.Mode
	}

	return t.Session.Do([]*typecastgo.TypecastExecuteRequest{request})
}
