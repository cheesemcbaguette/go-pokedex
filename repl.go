package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// Get the input text
	reader := bufio.NewScanner(os.Stdin)

	cfg := &config{}

	// REPL loop
	for {
		fmt.Print("Pokedex > ")

		// Wait for input
		reader.Scan()

		// Get the input text and clean it
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		// get the first word as a command
		commandName := words[0]

		// check if the command exist
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	// convertis en lowercase
	output := strings.ToLower(text)
	// split le texte en mots
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}
