package services

import (
	"github.com/jipark0716/typecastgo"
)

type Typecast struct {
	*typecastgo.Session
	Actors []typecastgo.TypecastActor
}

var typecastInstance *Typecast

func GetTypecastInstance() *Typecast {
	if typecastInstance == nil {
		session := typecastgo.NewSession()
		typecastInstance = &Typecast{
			Session: &session,
		}
		typecastInstance.Connect(&typecastgo.LoginRequest{
			Email:             GetConfigInstance().Get("TYPECAST_EMAIL").(string),
			Password:          GetConfigInstance().Get("TYPECAST_PASSWORD").(string),
			ReturnSecureToken: true,
		})
	}

	return typecastInstance
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
