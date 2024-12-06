package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	INPUT_FILE = "test.txt"
	WIDTH      = 10

	OBSTACLE    = '#'
	CLEAR       = '.'
	GUARD_NORTH = '^'
	GUARD_EAST  = '>'
	GUARD_SOUTH = 'V'
	GUARD_WEST  = '<'
)

func printMap(data []byte) {
	for i, v := range data {
		fmt.Printf("%s", string(v))
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
			if data[current-WIDTH] == OBSTACLE {
				data[current] = GUARD_EAST
				return current
			}
			data[current] = 'X'
			current -= WIDTH
		}
	}
	if direction == GUARD_EAST {
		for current%WIDTH < WIDTH {
			if data[current+1] == OBSTACLE {
				data[current] = GUARD_SOUTH
				return current
			}
			data[current] = 'X'
			current++
		}
	}
	if direction == GUARD_SOUTH {
		for current < len(data) {
			if data[current+WIDTH] == OBSTACLE {
				data[current] = GUARD_WEST
				return current
			}
			data[current] = 'X'
			current += WIDTH
		}
	}
	if direction == GUARD_WEST {
		for current > 0 {
			if data[current-1] == OBSTACLE {
				data[current] = GUARD_NORTH
				return current
			}
			data[current] = 'X'
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

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}

	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))

	fmt.Print("\033[H\033[2J")
	printMap(data)
	startingPos := findGuard(data)
	next := moveGuard(data, startingPos)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)
	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)
	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)
	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

	next = moveGuard(data, next)

	time.Sleep(1 * time.Second)
	fmt.Print("\033[H\033[2J")
	printMap(data)

}
