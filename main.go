package main

import (
	cli "github.com/adrg2/torrent-client/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "torrent-cli"}
	rootCmd.AddCommand(cli.InitDownloadCmd())
	rootCmd.Execute()
}
