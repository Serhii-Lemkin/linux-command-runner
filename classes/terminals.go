package classes

type TerminalCommand struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}
