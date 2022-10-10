/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init service",
	Long:  "init service",
	Run: func(cmd *cobra.Command, args []string) {
		// command, err := services.GetDiscordInstance().CreateChangeVoiceCommand()
		// if err != nil {
		// 	println("fail")
		// 	fmt.Printf("%+v\n", err)
		// 	return
		// }
		// fmt.Printf("%#v\n", command)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
