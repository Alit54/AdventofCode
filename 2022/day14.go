package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	Alit "github.com/Alit54/General/gofunc"
)

type Waterfall struct {
	minY int
	maxY int
	minX int
	maxX int
	draw map[Alit.Position]byte
}

func main() {
	fmt.Println(Day14())
}

func Day14() (int, int) {
	return Part14A(), Part14B()
}

func Part14A() int {
	var counter int
	waterfall := Waterfall{0, 0, 500, 500, make(map[Alit.Position]byte)}
	scanner := bufio.NewScanner(Alit.File(os.Args[1]))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " -> ")
		for i := 0; i < len(row)-1; i++ {
			waterfall.createRock(row[i], row[i+1])
		}
	}
	sand := Alit.Position{500, 0}
a:
	for {
		for {
			// Check: first Sud, then Sud-West, then Sud-East
			actual := Alit.Position{sand.Row, sand.Column + 1}
			if waterfall.draw[actual] != 'R' && waterfall.draw[actual] != 'S' && sand.Column < waterfall.maxY {
				sand = actual
				continue
			}
			actual = Alit.Position{sand.Row - 1, sand.Column + 1}
			if waterfall.draw[actual] != 'R' && waterfall.draw[actual] != 'S' && sand.Column < waterfall.maxY {
				sand = actual
				continue
			}
			actual = Alit.Position{sand.Row + 1, sand.Column + 1}
			if waterfall.draw[actual] != 'R' && waterfall.draw[actual] != 'S' && sand.Column < waterfall.maxY {
				sand = actual
				continue
			}
			if sand.Column >= waterfall.maxY || sand.Column <= waterfall.minY || sand.Row <= waterfall.minX || sand.Row >= waterfall.maxX {
				break a
			}
			waterfall.draw[sand] = 'S'
			sand = Alit.Position{500, 0}
			break
		}
		counter++
	}
	return counter
}

func Part14B() int {
	var counter int
	waterfall := Waterfall{0, 0, 500, 500, make(map[Alit.Position]byte)}
	scanner := bufio.NewScanner(Alit.File(os.Args[1]))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " -> ")
		for i := 0; i < len(row)-1; i++ {
			waterfall.createRock(row[i], row[i+1])
		}
	}
	sand := Alit.Position{500, 0}
	sw := Alit.Position{499, 1} // Sud-West of Sand
	s := Alit.Position{500, 1}  // Sud of Sand
	se := Alit.Position{501, 1} // // Sud-East of Sand
a:
	for {
		for {
			// Check if sand is blocked
			if (waterfall.draw[sw] == 'R' || waterfall.draw[sw] == 'S') && (waterfall.draw[s] == 'R' || waterfall.draw[s] == 'S') && (waterfall.draw[se] == 'R' || waterfall.draw[se] == 'S') {
				counter++
				break a
			}
			actual := Alit.Position{sand.Row, sand.Column + 1}
			if waterfall.draw[actual] != 'R' && waterfall.draw[actual] != 'S' {
				if sand.Column > waterfall.maxY {
					waterfall.draw[sand] = 'S'
					sand = Alit.Position{500, 0}
					break
				}
				sand = actual
				continue
			}
			actual = Alit.Position{sand.Row - 1, sand.Column + 1}
			if waterfall.draw[actual] != 'R' && waterfall.draw[actual] != 'S' {
				if sand.Row < waterfall.minX {
					waterfall.minX = sand.Row
				}
				if sand.Column > waterfall.maxY {
					waterfall.draw[sand] = 'S'
					sand = Alit.Position{500, 0}
					break
				}
				sand = actual
				continue
			}
			actual = Alit.Position{sand.Row + 1, sand.Column + 1}
			if waterfall.draw[actual] != 'R' && waterfall.draw[actual] != 'S' {
				if sand.Row > waterfall.maxX {
					waterfall.maxX = sand.Row
				}
				if sand.Column > waterfall.maxY {
					waterfall.draw[sand] = 'S'
					sand = Alit.Position{500, 0}
					break
				}
				sand = actual
				continue
			}
			if sand.Column > waterfall.maxY {
				waterfall.draw[sand] = 'S'
				sand = Alit.Position{500, 0}
				break
			}

			waterfall.draw[sand] = 'S'
			sand = Alit.Position{500, 0}
			break
		}
		counter++

	}
	return counter
}

// Creates a line of rocks into the waterfall.
func (waterfall *Waterfall) createRock(a, b string) {
	// Data
	coordinate1 := strings.Split(a, ",")
	coordinate2 := strings.Split(b, ",")
	x1, _ := strconv.Atoi(coordinate1[0])
	y1, _ := strconv.Atoi(coordinate1[1])
	x2, _ := strconv.Atoi(coordinate2[0])
	y2, _ := strconv.Atoi(coordinate2[1])
	// Updating
	waterfall.update(x1, x2, y1, y2)
	// Create line of rocks
	if x1 == x2 {
		for i := y1; i <= y2; i++ {
			waterfall.draw[Alit.Position{x1, i}] = 'R'
		}

		for i := y1; i >= y2; i-- {
			waterfall.draw[Alit.Position{x1, i}] = 'R'
		}
	} else {

		for i := x1; i <= x2; i++ {
			waterfall.draw[Alit.Position{i, y1}] = 'R'
		}

		for i := x1; i >= x2; i-- {
			waterfall.draw[Alit.Position{i, y1}] = 'R'
		}
	}
}

// Updates the minimun and maximun row and column of the waterfall.
func (waterfall *Waterfall) update(x1, x2, y1, y2 int) {

	if x1 < waterfall.minX {
		waterfall.maxX = x1
	}
	if x2 < waterfall.minX {
		waterfall.minX = x2
	}
	if x1 > waterfall.maxX {
		waterfall.maxX = x1
	}
	if x2 > waterfall.maxX {
		waterfall.maxX = x2
	}
	if y1 < waterfall.minY {
		waterfall.minY = y1
	}
	if y2 < waterfall.minY {
		waterfall.minY = y2
	}
	if y1 > waterfall.maxY {
		waterfall.maxY = y1
	}
	if y2 > waterfall.maxY {
		waterfall.maxY = y2
	}
}
