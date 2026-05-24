package cmd

import (
	"fmt"

	"github.com/27te/winshell/internal/themes"
	"github.com/27te/winshell/internal/ui"
	"github.com/spf13/cobra"
)

var themeCmd = &cobra.Command{
	Use:   "theme [name]",
	Short: "Switch terminal theme",
	Long: `Switch your terminal theme across WezTerm, Starship and fzf at once.

Available themes:
  tokyo-night       Blue/purple dark
  catppuccin-mocha  Soft purple dark
  solarized-dark    Classic yellow/cyan dark`,
	Args: cobra.ExactArgs(1),
	RunE: runTheme,
}

var themeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available themes",
	Run:   runThemeList,
}

func init() {
	rootCmd.AddCommand(themeCmd)
	themeCmd.AddCommand(themeListCmd)
}

func runTheme(cmd *cobra.Command, args []string) error {
	name := args[0]

	available := themes.List()
	for _, t := range available {
		if t == name {
			fmt.Printf("%s Applying theme %s\n\n", ui.Arrow, ui.ThemeBadge(name))
			err := themes.Apply(name)
			if err != nil {
				return err
			}
			fmt.Printf("\n%s Restart WezTerm to see changes.\n", ui.Muted.Render("hint"))
			return nil
		}
	}

	fmt.Printf("%s Theme %s not found.\n\n", ui.CrossMark, ui.Bold.Render(name))
	runThemeList(cmd, args)
	return nil
}

func runThemeList(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n\n", ui.Title("available themes"))
	current := themes.GetCurrent()
	for _, t := range themes.List() {
		indicator := ui.Dot
		if t == current {
			indicator = ui.CheckMark
		}
		fmt.Printf("  %s %s\n", indicator, ui.ThemeBadge(t))
	}
	fmt.Println()
}
