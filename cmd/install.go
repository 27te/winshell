package cmd

import (
	"fmt"

	"github.com/27te/winshell/internal/detector"
	"github.com/27te/winshell/internal/installer"
	"github.com/27te/winshell/internal/ui"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and configure your terminal environment",
	RunE:  runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) error {
	fmt.Println(ui.Title("winshell install"))
	fmt.Printf("\n%s Detecting installed tools...\n\n", ui.Arrow)

	installed := detector.GetInstalled()
	missing := detector.GetMissing()

	for _, t := range installed {
		fmt.Printf("  %s %s\n", ui.CheckMark, ui.Muted.Render(t))
	}

	if len(missing) == 0 {
		fmt.Printf("\n%s All tools are already installed.\n", ui.CheckMark)
		return nil
	}

	fmt.Printf("\n%s Missing: ", ui.Warn.Render("!"))
	for _, t := range missing {
		fmt.Printf("%s ", ui.Error.Render(t))
	}
	fmt.Printf("\n\n%s Installing...\n\n", ui.Arrow)

	return installer.InstallAll(missing)
}
