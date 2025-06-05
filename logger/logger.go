package logger

import (
	"fmt"
	"rnnr/classes"
	"strings"
)

func LogError(err error) {
	fmt.Println("An Error : ", err)
}

func GetUserConfirmation() bool {
	fmt.Print("Are you sure? y/n ")
	var input string

	for {
		fmt.Scanln(&input)
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("incorrect input, use y/n")
			continue
		}
	}
}

func LogAlias(message string, alias classes.Alias, name string) {
	fmt.Println(message, name, alias)
}

func Log(args ...any) {
	fmt.Println(args...)
}
