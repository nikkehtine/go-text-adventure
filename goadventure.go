package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"gitlab.com/nikkehtine/go-text-adventure/logic"
)

func main() {
	fmt.Println(logic.Test())
	fmt.Println(worldgen.Test())

	showHelp()
	for {
		handleInput()
	}
}

func handleInput() {
	fmt.Print(">>>")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "start":
		fmt.Println("Hello world!")
	case "quit":
		print("Goodbye soldier. We will miss you.")
		os.Exit(0)
	case "help":
		showHelp()
	default:
		print("No such command. Did you make a typo?")
	}
}

func print(message string) {
	fmt.Println()
	fmt.Printf("%s", message)
	fmt.Println()
}

func showHelp() {
	commands := map[string]string{
		"start":   "Start a new game",
		"help":    "Print help message",
		"credits": "Print credits for the game",
		"quit":    "Quit the game",
	}

	head := color.New(color.FgGreen).SprintFunc()
	print(head("Available commands:"))

	for command, description := range commands {
		name := color.New(color.FgYellow).SprintFunc()
		print(fmt.Sprintf("%s\n %s", name(command), description))
	}
}
