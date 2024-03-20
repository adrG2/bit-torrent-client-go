package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "torrent-client",
	Short: "Torrent client",
	Long:  `A fast torrent client using bittorrent protocol`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("First command %s", args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
