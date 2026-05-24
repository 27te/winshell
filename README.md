<div align="center">

# winshell

**Terminal environment manager for Windows**

Set up a fast, beautiful terminal in one command.

[![Release](https://img.shields.io/github/v/release/27te/winshell?color=7aa2f7&labelColor=1a1b26)](https://github.com/27te/winshell/releases)
[![License](https://img.shields.io/github/license/27te/winshell?color=9ece6a&labelColor=1a1b26)](LICENSE)
[![Go](https://img.shields.io/badge/go-1.22+-bb9af7?labelColor=1a1b26)](https://golang.org)

</div>

---

## What is winshell?

Configuring a proper terminal on Windows takes hours — installing tools, hunting for configs on Reddit, breaking your PowerShell profile.

WinShell does it in one command. It installs and configures WezTerm, Starship, Zoxide, PSFzf and more, with switchable themes across everything at once.

## Install

Download `winshell.exe` from [releases](https://github.com/27te/winshell/releases) and add it to your PATH, or:

```powershell
# Coming soon via scoop
scoop install winshell
```

## Usage

```powershell
# Install missing tools automatically
winshell install

# Switch theme across WezTerm + Starship + fzf at once
winshell theme tokyo-night
winshell theme catppuccin-mocha
winshell theme solarized-dark

# List available themes
winshell theme list

# Backup your current config
winshell backup

# Update all tools via scoop
winshell update
```

## Themes

| Theme | Preview |
|-------|---------|
| `tokyo-night` | Blue/purple dark — the most popular dev theme today |
| `catppuccin-mocha` | Soft purple dark — smooth and easy on the eyes |
| `solarized-dark` | Classic yellow/cyan dark — timeless |

## What gets installed

| Tool | Purpose |
|------|---------|
| [WezTerm](https://wezfurlong.org/wezterm/) | GPU-accelerated terminal emulator |
| [Starship](https://starship.rs) | Fast, minimal shell prompt |
| [Zoxide](https://github.com/ajeetdsouza/zoxide) | Smarter `cd` with frecency |
| [fzf](https://github.com/junegunn/fzf) | Fuzzy finder for history and files |
| [eza](https://github.com/eza-community/eza) | Modern `ls` replacement |
| [bat](https://github.com/sharkdp/bat) | `cat` with syntax highlighting |
| [fnm](https://github.com/Schniz/fnm) | Fast Node.js version manager |
| [PSFzf](https://github.com/kelleyma49/PSFzf) | fzf integration for PowerShell |

## Requirements

- Windows 10/11
- [Scoop](https://scoop.sh) package manager
- PowerShell 7+ (`pwsh`)

## Build from source

```powershell
git clone https://github.com/27te/winshell
cd winshell
go build -ldflags="-s -w" -o winshell.exe .
```

## License

MIT — see [LICENSE](LICENSE)
