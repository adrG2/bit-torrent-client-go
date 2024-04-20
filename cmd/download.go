package cli

import (
	"log"

	"github.com/spf13/cobra"
)

type Cobra func(cmd *cobra.Command, args []string)

const in = "input"
const out = "out"

func InitDownloadCmd() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Download torrent file",
		Run:   runDownload(),
	}

	downloadCmd.Flags().StringP(in, "i", "", "input file path")
	downloadCmd.Flags().StringP(out, "o", "", "output file path")

	return downloadCmd
}

func runDownload() Cobra {
	return func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString(in)
		out, _ := cmd.Flags().GetString(out)

		log.Printf("{input:\"%s\", output:\"%s\"}", input, out)
	}
}
