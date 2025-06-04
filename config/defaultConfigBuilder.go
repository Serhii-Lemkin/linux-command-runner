package config

import "rnnr/classes"

func getDefaultConfig() classes.Config {
	cfg := classes.Config{}
	cfg.DefaultSameTerminal = true
	cfg.SearchColor = "blue"
	cfg.ColorMap = DefaultColorMap
	return cfg
}
