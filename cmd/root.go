package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prcheckctl",
	Short: "prcheckctl is a CLI to check for open PRs in a organization",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		pool, _ := cmd.Flags().GetInt16("pool")

		getAllPRs(username, int(pool))

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "github username or organization")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.PersistentFlags().Int16P("pool", "p", 30, "time to pool for new PRs")
}
