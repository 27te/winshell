package detector

import "os/exec"

var tools = []string{
	"wezterm",
	"starship",
	"zoxide",
	"fzf",
	"eza",
	"bat",
	"fnm",
	"pwsh",
}

func IsInstalled(tool string) bool {
	_, err := exec.LookPath(tool)
	return err == nil
}

func GetMissing() []string {
	missing := []string{}
	for _, tool := range tools {
		if !IsInstalled(tool) {
			missing = append(missing, tool)
		}
	}
	return missing
}

func GetInstalled() []string {
	installed := []string{}
	for _, tool := range tools {
		if IsInstalled(tool) {
			installed = append(installed, tool)
		}
	}
	return installed
}
