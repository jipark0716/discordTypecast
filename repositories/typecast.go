package repositories

import (
	"github.com/jipark0716/typecastgo"
)

type Typecast struct {
	typecastgo.Session
	Actors []typecastgo.TypecastActor
}

func NewTypecast() (typecast Typecast) {
	typecast.Session = typecastgo.NewSession()
	typecast.Connect(&typecastgo.LoginRequest{
		Email:             GetConfigInstnace().Get("TYPECAST_EMAIL").(string),
		Password:          GetConfigInstnace().Get("TYPECAST_PASSWORD").(string),
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
