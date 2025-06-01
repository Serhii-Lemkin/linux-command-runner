package helpers

import (
	"fmt"
	"rnnr/classes"
	"strings"
)

func LogError(err error) {
	fmt.Println("An Error : ", err)
}

func GetUserConfirmation() bool {
	fmt.Println("Are you sure? y/n")
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

func LogAlias(data string, alias classes.Alias) {
	fmt.Println(data, alias)
}

func Log(args ...any) {
	fmt.Println(args...)
}
