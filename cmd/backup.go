package cmd

import (
	"fmt"

	"github.com/27te/winshell/internal/config"
	"github.com/27te/winshell/internal/ui"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your current terminal config",
	RunE:  runBackup,
}

func init() {
	rootCmd.AddCommand(backupCmd)
}

func runBackup(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s Backing up current config...\n\n", ui.Arrow)
	path, err := config.Backup()
	if err != nil {
		return err
	}
	fmt.Printf("  %s wezterm.lua\n", ui.CheckMark)
	fmt.Printf("  %s starship.toml\n", ui.CheckMark)
	fmt.Printf("  %s profile.ps1\n", ui.CheckMark)
	fmt.Printf("\n%s Saved to: %s\n", ui.Muted.Render("path"), ui.Info.Render(path))
	return nil
}
