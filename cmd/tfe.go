package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "tfe",
}

// Run runs command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("install", "i", "", "install terraform specified version")
	rootCmd.Flags().StringP("list-remote", "r", "", "display versions of all available terraform versions")
	rootCmd.Flags().StringP("list", "l", "", "display versions of installed terraform versions")
	rootCmd.Flags().StringP("use", "u", "", "use specified terraform version on local")
}
