package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jipark0716/discordTypecast/ent/migrate"
)

func Migrate() (err error) {
	client, err := GetConnection()
	if err != nil {
		return
	}

	defer client.Close()
	err = client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
	return
}
