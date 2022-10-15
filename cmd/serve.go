package cmd

import (
	"fmt"

	"github.com/jipark0716/discordTypecast/wire"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start service",
	Long:  "start service",
	Run: func(cmd *cobra.Command, args []string) {
		session, err := wire.NewDiscordService()
		err = session.Serve()
		if err != nil {
			println("fail")
			fmt.Printf("%+v\n", err)
		} else {
			println("connect server")
		}
		select {}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
