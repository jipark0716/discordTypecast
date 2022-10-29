package repositories

import (
	"context"

	"github.com/jipark0716/discordTypecast/ent"
	"github.com/jipark0716/discordTypecast/ent/usertypecastsetting"
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

func (ur *UserRepository) FindByUserId(userId string) (*ent.UserTypecastSetting, error) {
	return ur.
		Database.
		UserTypecastSetting.
		Query().
		Where(
			usertypecastsetting.
				UserID(userId),
		).
		Only(context.Background())
}
