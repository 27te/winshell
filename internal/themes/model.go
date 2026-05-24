package themes

import "encoding/json"

type Theme struct {
	Name     string        `json:"name"`
	WezTerm  WezTermTheme  `json:"wezterm"`
	Starship StarshipTheme `json:"starship"`
	Fzf      FzfTheme      `json:"fzf"`
}

type WezTermTheme struct {
	ColorScheme string `json:"color_scheme"`
}

type StarshipTheme struct {
	Palette     string `json:"palette"`
	PromptColor string `json:"prompt_color"`
	DirColor    string `json:"dir_color"`
	BranchColor string `json:"branch_color"`
	StatusColor string `json:"status_color"`
	NodeColor   string `json:"node_color"`
	PythonColor string `json:"python_color"`
	GoColor     string `json:"go_color"`
}

type FzfTheme struct {
	Opts string `json:"opts"`
}

func parseTheme(data []byte) (*Theme, error) {
	var t Theme
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
