/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/jipark0716/discordTypecast/database"
	"github.com/jipark0716/discordTypecast/wire"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init service",
	Long:  "init service",
	Run: func(cmd *cobra.Command, args []string) {
		discord, err := wire.NewDiscordService()
		command, err := discord.CreateChangeVoiceCommand()
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
		fmt.Printf("%#v\n", command)
		err = database.Migrate()
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
