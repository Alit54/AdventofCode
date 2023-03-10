package main

import (
	"bufio"
	"fmt"
	"os"

	Alit "github.com/Alit54/General/gofunc"
)

func main() {
	fmt.Println(Day15())
}

func Day15() (int, int) {
	return Part15A(), 0
}

type Tunnel struct {
	minX      int
	maxX      int
	notBeacon map[int]bool
}

func Part15A() int {
	var sensor, beacon Alit.Position
	file := Alit.File(os.Args[1])
	scanner := bufio.NewScanner(file)
	var tunnel Tunnel
	tunnel.minX = Alit.Min()
	tunnel.maxX = Alit.Max()
	tunnel.notBeacon = map[int]bool{}
	var counter int
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.Row, &sensor.Column, &beacon.Row, &beacon.Column)
		distance := Alit.Manhattan(sensor, beacon)
		signal := Alit.Absolute(sensor.Column - 2000000)
		// Calculate how much the signal impact on the 2M° row
		for i := sensor.Row - (distance - signal); i <= sensor.Row+(distance-signal); i++ {
			tunnel.notBeacon[i] = true
		}
		if sensor.Row-(distance-signal) < tunnel.minX {
			tunnel.minX = sensor.Row - (distance - signal)
		}
		if sensor.Row+(distance-signal) > tunnel.maxX {
			tunnel.maxX = sensor.Row + (distance - signal)
		}
	}
	// count how many position are true
	for i := tunnel.minX; i < tunnel.maxX; i++ {
		if tunnel.notBeacon[i] {
			counter++
		}
	}
	return counter
}

func Part15B() int {
	var sensor, beacon Alit.Position
	file := Alit.File(os.Args[1])
	scanner := bufio.NewScanner(file)
	var tunnel Tunnel
	var counter int
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.Row, &sensor.Column, &beacon.Row, &beacon.Column)
		distance := Alit.Manhattan(sensor, beacon)
		signal := Alit.Absolute(sensor.Column - 200000)
		// Calculate how much the signal impact on the 200k° row
		if distance > signal {
			for i := sensor.Row - (distance - signal); i <= sensor.Row+(distance-signal); i++ {
				tunnel.notBeacon[i] = true
			}
			if sensor.Row-signal < tunnel.minX {
				tunnel.minX = sensor.Row - (distance - signal)
			}
			if sensor.Row+signal > tunnel.maxX {
				tunnel.maxX = sensor.Row + (distance - signal)
			}
		}
	}
	// count how many position are true
	for i := tunnel.minX; i <= tunnel.maxX; i++ {
		if tunnel.notBeacon[i] {
			counter++
		}
	}
	return counter
}
