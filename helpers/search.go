package helpers

import (
	"fmt"
	"os"
	"regexp"
	"rnnr/classes"
	"rnnr/config"
	"rnnr/logger"
	"strings"
)

func FuzzyFindAliasByName() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		logger.Log("the usage is < rnnr fuzzy 'yourSearch' optional(-a/-c)>")
		return
	}

	aliases, err := LoadAliases()
	if err != nil {
		logger.LogError(err)
		return
	}

	if len(aliases) == 0 {
		logger.Log("No aliases saved.")
		return
	}

	foundAliases := make(map[string]classes.Alias)
	searchText := os.Args[2]
	searchInCommands := true
	searchInAliases := true
	if len(os.Args) > 3 {
		option := os.Args[3]
		searchInAliases = option == "-alias" || option == "-a"
		searchInCommands = option == "-command" || option == "-c"
	}

	for name, alias := range aliases {
		foundInName := false
		if searchInAliases {
			foundInName = strings.Contains(name, searchText)
		}

		foundInCommand := false
		if searchInCommands {
			for _, command := range alias.Commands {
				if strings.Contains(command, searchText) {
					foundInCommand = true
					break
				}
			}
		}

		if foundInName || foundInCommand {
			foundAliases[name] = alias
		}
	}

	if len(foundAliases) > 0 {
		logger.Log("Found aliases:")
		for name, a := range foundAliases {
			if searchInAliases {
				fmt.Printf("  %s:\n", highlightMatches(name, searchText))
			} else {
				fmt.Printf("  %s:\n", name)
			}
			for i, c := range a.Commands {
				if searchInCommands {
					fmt.Printf("    %d. %s\n", i+1, highlightMatches(c, searchText))
				} else {
					fmt.Printf("    %d. %s\n", i+1, c)
				}
			}
		}
	} else {
		logger.Log("nothing found")
	}
}

func highlightMatches(text string, search string) string {
	configItem, err := config.GetConfig()
	if err != nil {
		return text
	}

	escapedSearch := regexp.QuoteMeta(search)
	re := regexp.MustCompile(`(?i)` + escapedSearch)

	return re.ReplaceAllStringFunc(text, func(m string) string {
		return configItem.ColorMap[configItem.SearchColor] + m + "\033[0m"
	})
}
