package config

import "rnnr/classes"

func getDefaultConfig() classes.Config {
	cfg := classes.Config{}
	cfg.DefaultSameTerminal = true
	cfg.SearchColor = "red"
	cfg.ColorMap = DefaultColorMap
	cfg.EditorList = EditorList
	cfg.TerminalList = TerminalAndCommandList
	return cfg
}
