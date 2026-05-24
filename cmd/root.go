package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "winshell",
	Short: "Terminal environment manager for Windows",
	Long: `WinShell — setup and manage your Windows terminal environment.
Installs and configures WezTerm, Starship, Zoxide, PSFzf and more in one command.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
