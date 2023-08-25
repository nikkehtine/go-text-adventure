package worldgen

// DISCLAIMER
// I have absolutely no idea what I'm doing here
// What the fuck am I looking at???

func Test() bool {
	return true
}

func GenerateLevel(args ...int) (string, error) {
	// Generate at least 5 and at most 16 rooms in a level by default
	// var rooms int = 5 + rand.Intn(11)
	// layout := generateLayout(rooms)
}

// Returns an array containing level layout
func generateLayout(rooms int) [5][9]int {
	var layout = [5][9]int{}
	return layout
}

func genRooms(direction string, x int, y int) {

}
