package logic

import (
	"fmt"

	"gitlab.com/nikkehtine/go-text-adventure/worldgen"
)

func Test() string {
	var wgstat string
	if worldgen.Test() {
		wgstat = "working"
	} else {
		wgstat = "NOT working!"
	}

	return fmt.Sprintf("Server status: working\nWorldgen status: %s", wgstat)
}
