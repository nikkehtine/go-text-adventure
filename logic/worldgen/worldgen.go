package worldgen

import (
	"errors"
	"fmt"
)

// DISCLAIMER
// I have absolutely no idea what I'm doing here
// What the fuck am I looking at???

func Test() bool {
	return true
}

// TODO: now it returns a string
func GenerateLevel(args ...int) (string, error) {
	var x, y int

	switch len(args) {
	case 0:
		x, y = 9, 5
	case 2:
		x, y = args[0], args[1]
	case 1:
		return "", errors.New("too few arguments (expected 0 or 2)")
	default:
		return "", errors.New("too many arguments (2 max)")
	}

	// Generate at least 5 and at most 16 rooms in a level by default
	// var rooms int = 5 + rand.Intn(11)
	// layout := generateLayout(rooms)
	return fmt.Sprintf("%d %d", x, y), nil
}

// Returns an array containing level layout
func generateLayout(rooms int) [5][9]int {
	var layout = [5][9]int{}
	return layout
}

func genRooms(direction string, x int, y int) {

}
