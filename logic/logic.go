package logic

import (
	"fmt"

	"gitlab.com/nikkehtine/go-text-adventure/logic/worldgen"
)

func Test() string {
	var worldgenStatus string
	if worldgen.Test() {
		worldgenStatus = "working"
	} else {
		worldgenStatus = "NOT working!"
	}

	return fmt.Sprintf("Server status: working\nWorldgen status: %s", worldgenStatus)
}

func GenerateLevel() {
	level, err := worldgen.GenerateLevel()
	if err != nil {
		panic(err)
	}
	fmt.Println(level)
}
