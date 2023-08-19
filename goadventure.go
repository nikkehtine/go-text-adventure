package main

import (
	"fmt"

	"github.com/fatih/color"

	"gitlab.com/nikkehtine/go-text-adventure/logic"
	"gitlab.com/nikkehtine/go-text-adventure/worldgen"
)

func main() {
	fmt.Println(logic.Test())
	fmt.Println(worldgen.Test())
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red("error"))
}
