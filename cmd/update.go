package cmd

import (
	"fmt"
	"os/exec"

	"github.com/27te/winshell/internal/ui"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all installed tools via scoop",
	RunE:  runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s Updating all tools via scoop...\n\n", ui.Arrow)

	c := exec.Command("scoop", "update", "*")
	c.Stdout = nil
	c.Stderr = nil

	if err := c.Run(); err != nil {
		return fmt.Errorf("scoop update failed: %w", err)
	}

	fmt.Printf("%s All tools updated.\n", ui.CheckMark)
	return nil
}
