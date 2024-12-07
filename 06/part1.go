package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	INPUT_FILE = "input.txt"
	WIDTH      = 130

	OBSTACLE    = '#'
	CLEAR       = '.'
	VISITED     = 'X'
	GUARD_NORTH = '^'
	GUARD_EAST  = '>'
	GUARD_SOUTH = 'V'
	GUARD_WEST  = '<'
)

func printMap(data []byte) {
	for i, _ := range data {
		switch data[i] {
		case CLEAR:
			fmt.Printf("\033[1;32m%s", ".")
		case VISITED:
			fmt.Printf("\033[1;31m%s", "X")
		case OBSTACLE:
			fmt.Printf("\033[1;30m%s", "#")
		default:
			fmt.Printf("%s", "X")
		}
		if (i+1)%WIDTH == 0 {
			fmt.Println()
		}
	}
}

func moveGuard(data []byte, pos int) int {
	direction := data[pos]

	current := pos
	if direction == GUARD_NORTH {
		for current > 0 {
			if current < 0 || current-WIDTH < 0 {
				return -1
			}
			if data[current-WIDTH] == OBSTACLE {
				data[current] = GUARD_EAST
				return current
			}
			data[current] = VISITED
			current -= WIDTH
		}
	}
	if direction == GUARD_EAST {
		for current%WIDTH < WIDTH {
			if current%WIDTH > WIDTH || current%WIDTH+1 > WIDTH {
				return -1
			}
			if data[current+1] == OBSTACLE {
				data[current] = GUARD_SOUTH
				return current
			}
			data[current] = VISITED
			current++
		}
	}
	if direction == GUARD_SOUTH {
		for current < len(data) {
			if current > len(data) || current+WIDTH > len(data) {
				return -1
			}
			if data[current+WIDTH] == OBSTACLE {
				data[current] = GUARD_WEST
				return current
			}
			data[current] = VISITED
			current += WIDTH
		}
	}
	if direction == GUARD_WEST {
		for current > 0 {
			if current%WIDTH == 0 {
				return -1
			}
			if data[current-1] == OBSTACLE {
				data[current] = GUARD_NORTH
				return current
			}
			data[current] = VISITED
			current -= 1
		}
	}

	return -1
}

func findGuard(data []byte) int {
	for i, b := range data {
		if b == GUARD_NORTH || b == GUARD_EAST || b == GUARD_SOUTH || b == GUARD_WEST {
			return i
		}
	}

	return -1
}

func countVisited(data []byte) int {
	sum := 0
	for _, v := range data {
		if v == VISITED {
			sum++
		}
	}

	return sum
}

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}

	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))

	printMap(data)
	next := findGuard(data)
	for next > 0 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
		printMap(data)
		next = moveGuard(data, next)
	}

	fmt.Println(countVisited(data) + 1)
}
