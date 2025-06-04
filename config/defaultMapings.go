package config

import "rnnr/classes"

var DefaultColorMap = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",

	"bright_black":   "\033[90m",
	"bright_red":     "\033[91m",
	"bright_green":   "\033[92m",
	"bright_yellow":  "\033[93m",
	"bright_blue":    "\033[94m",
	"bright_magenta": "\033[95m",
	"bright_cyan":    "\033[96m",
	"bright_white":   "\033[97m",
}

var EditorList = []string{
	"nvim",
	"gedit",
	"nano",
	"vi",
}

var TerminalAndCommandList = []classes.TerminalCommand{
	{Name: "ptyxis", Args: []string{"--", "bash", "-c"}},
	{Name: "alacritty", Args: []string{"-e", "bash", "-c"}},
	{Name: "kitty", Args: []string{"bash", "-c"}},
	{Name: "wezterm", Args: []string{"start", "--", "bash", "-c"}},
	{Name: "gnome-terminal", Args: []string{"--", "bash", "-c"}},
	{Name: "xterm", Args: []string{"-e", "bash", "-c"}},
}
