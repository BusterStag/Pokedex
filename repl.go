package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startReplace() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	var set []string
	trim_text := strings.TrimSpace(text)
	lower_text := strings.ToLower(trim_text)
	set = strings.Fields(lower_text)
	return set
}
