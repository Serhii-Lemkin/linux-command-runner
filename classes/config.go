package classes

type Config struct {
	PreferedTerminal    string `json:"preferred_terminal"`
	PreferedFileManager string `json:"preferred_file_manager"`
	DefaultSameTerminal bool   `json:"default_same_terminal"`
}
