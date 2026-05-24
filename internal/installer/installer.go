package installer

import (
	"fmt"
	"os/exec"
)

// scoopPackages maps tool names to their scoop package names
var scoopPackages = map[string]string{
	"wezterm":  "wezterm",
	"starship": "starship",
	"zoxide":   "zoxide",
	"fzf":      "fzf",
	"eza":      "eza",
	"bat":      "bat",
	"fnm":      "fnm",
	"pwsh":     "pwsh",
}

func InstallAll(tools []string) error {
	// Check scoop is available first
	if _, err := exec.LookPath("scoop"); err != nil {
		return fmt.Errorf("scoop not found — install it first: https://scoop.sh")
	}

	for _, tool := range tools {
		if err := installOne(tool); err != nil {
			fmt.Printf("⚠  Failed to install %s: %v\n", tool, err)
			continue
		}
		fmt.Printf("✓ Installed: %s\n", tool)
	}
	return nil
}

func installOne(tool string) error {
	pkg, ok := scoopPackages[tool]
	if !ok {
		return fmt.Errorf("no scoop package defined for %s", tool)
	}

	cmd := exec.Command("scoop", "install", pkg)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
