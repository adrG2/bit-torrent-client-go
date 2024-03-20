package cli

import (
	"log"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const inputPath = "input"
const outPath = "out"

// InitDownloadCmd initialize download command
func InitDownloadCmd() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Download torrent",
		Run:   runDownloadFn(),
	}

	downloadCmd.Flags().StringP(inputPath, "i", "", "input file path")
	downloadCmd.Flags().StringP(outPath, "o", "", "out file path")

	return downloadCmd
}

func runDownloadFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString(inputPath)
		out, _ := cmd.Flags().GetString(outPath)

		log.Printf("{input:\"%s\", output:\"%s\"}", input, out)
	}
}
