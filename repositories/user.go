package repositories

import (
	"context"

	"github.com/jipark0716/discordTypecast/ent"
)

type UserRepository struct {
	Database *ent.Client
}

func NewUserRepository(database *ent.Client) UserRepository {
	return UserRepository{
		Database: database,
	}
}

func (ur *UserRepository) SaveActor(userId string, actorId string) error {
	return ur.
		Database.
		UserTypecastSetting.
		Create().
		SetActorID(actorId).
		SetUserID(userId).
		OnConflict().
		UpdateActorID().
		Exec(context.Background())
}
