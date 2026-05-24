package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorGreen  = lipgloss.Color("#9ece6a")
	colorRed    = lipgloss.Color("#f7768e")
	colorBlue   = lipgloss.Color("#7aa2f7")
	colorYellow = lipgloss.Color("#e0af68")
	colorPurple = lipgloss.Color("#bb9af7")
	colorMuted  = lipgloss.Color("#565f89")
	colorWhite  = lipgloss.Color("#c0caf5")

	Success = lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	Error   = lipgloss.NewStyle().Foreground(colorRed).Bold(true)
	Info    = lipgloss.NewStyle().Foreground(colorBlue)
	Warn    = lipgloss.NewStyle().Foreground(colorYellow)
	Accent  = lipgloss.NewStyle().Foreground(colorPurple).Bold(true)
	Muted   = lipgloss.NewStyle().Foreground(colorMuted)
	Bold    = lipgloss.NewStyle().Foreground(colorWhite).Bold(true)

	CheckMark = Success.Render("✓")
	CrossMark = Error.Render("✗")
	Arrow     = Accent.Render("❯")
	Dot       = Muted.Render("·")
)

func Title(s string) string {
	return lipgloss.NewStyle().
		Foreground(colorPurple).
		Bold(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(colorMuted).
		PaddingBottom(1).
		Width(40).
		Render(s)
}

func Badge(s string, color lipgloss.Color) string {
	return lipgloss.NewStyle().
		Foreground(color).
		Padding(0, 1).
		Bold(true).
		Render(s)
}

func ThemeBadge(name string) string {
	colors := map[string]lipgloss.Color{
		"tokyo-night":      "#7aa2f7",
		"catppuccin-mocha": "#cba6f7",
		"solarized-dark":   "#268bd2",
	}
	c, ok := colors[name]
	if !ok {
		c = colorMuted
	}
	return Badge(name, c)
}
