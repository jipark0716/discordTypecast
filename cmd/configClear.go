/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jipark0716/discordTypecast/services"
	"github.com/spf13/cobra"
)

// configClearCmd represents the configClear command
var configClearCmd = &cobra.Command{
	Use:   "configClear",
	Short: "clear config cache clear",
	Long:  "clear config cache clear",
	Run: func(cmd *cobra.Command, args []string) {
		services.GetConfigInstance()
	},
}

func init() {
	rootCmd.AddCommand(configClearCmd)
}
