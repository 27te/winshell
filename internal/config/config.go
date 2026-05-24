package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func WezTermPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".wezterm.lua")
}

func StarshipPath() string {
	config, _ := os.UserConfigDir()
	return filepath.Join(config, "starship.toml")
}

func PowerShellProfilePath() string {
	// Ask PowerShell directly — most reliable
	out, err := exec.Command("pwsh", "-NoLogo", "-NonInteractive", "-Command", "$PROFILE").Output()
	if err == nil {
		path := strings.TrimSpace(string(out))
		if path != "" {
			// Check if file exists at reported path
			if _, err := os.Stat(path); err == nil {
				return path
			}
		}
	}

	home, _ := os.UserHomeDir()
	psDir := filepath.Join(home, "Documents", "PowerShell")

	// Try all common profile filenames
	candidates := []string{
		"profile.ps1",
		"Microsoft.PowerShell_profile.ps1",
	}

	for _, name := range candidates {
		p := filepath.Join(psDir, name)
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}

	// Fallback: OneDrive path
	oneDrive := filepath.Join(home, "OneDrive", "Documents", "PowerShell", "profile.ps1")
	if _, err := os.Stat(oneDrive); err == nil {
		return oneDrive
	}

	// Default — return standard path even if it doesn't exist yet
	return filepath.Join(psDir, "Microsoft.PowerShell_profile.ps1")
}

func BackupDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".winshell", "backups")
}

func Backup() (string, error) {
	ts := time.Now().Format("2006-01-02_15-04-05")
	dir := filepath.Join(BackupDir(), ts)

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", fmt.Errorf("could not create backup dir: %w", err)
	}

	files := map[string]string{
		WezTermPath():           "wezterm.lua",
		StarshipPath():          "starship.toml",
		PowerShellProfilePath(): "profile.ps1",
	}

	backed := 0
	for src, name := range files {
		if _, err := os.Stat(src); os.IsNotExist(err) {
			continue
		}
		dst := filepath.Join(dir, name)
		if err := copyFile(src, dst); err != nil {
			fmt.Printf("⚠  Could not backup %s: %v\n", name, err)
			continue
		}
		backed++
	}

	if backed == 0 {
		return "", fmt.Errorf("no config files found to backup")
	}

	return dir, nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0o644)
}
