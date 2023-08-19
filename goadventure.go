package main

import (
	"fmt"

	"github.com/fatih/color"

	"gitlab.com/nikkehtine/go-text-adventure/logic"
)

func main() {
	fmt.Println(logic.Test())
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
}
