package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	DIAG_DOWN_RIGHT = iota
	DIAG_DOWN_LEFT

	INPUT_FILE        = "input.txt"
	KEY_WORD          = "MAS"
	KEY_WORD_REVERSED = "SAM"

	INPUT_COL_WIDTH = 140
)

func occurances(data []byte) int {
	s := string(data)
	sum := 0
	if s == KEY_WORD {
		sum++
	}
	if s == KEY_WORD_REVERSED {
		sum++
	}

	return sum
}

func bytesInDirection(data []byte, i int, direction int) []byte {
	b := make([]byte, len(KEY_WORD))
	b[0] = data[i]

	nextIndex := 0
	for j := 1; j < len(KEY_WORD); j++ {
		switch direction {
		case DIAG_DOWN_RIGHT:
			nextIndex = i + j + INPUT_COL_WIDTH*j
		case DIAG_DOWN_LEFT:
			nextIndex = i - j + INPUT_COL_WIDTH*j
		}

		b[j] = data[nextIndex]
	}

	return b
}

func findDiagR(data []byte, i int) int {
	if i%INPUT_COL_WIDTH > INPUT_COL_WIDTH-len(KEY_WORD) {
		return 0
	}

	if i+2+(INPUT_COL_WIDTH*2) > len(data)-1 {
		return 0
	}

	return occurances(bytesInDirection(data, i, DIAG_DOWN_RIGHT))
}

func findDiagL(data []byte, i int) int {
	if i%INPUT_COL_WIDTH < len(KEY_WORD)-1 {
		return 0
	}

	if i > len(data)-(INPUT_COL_WIDTH*2) {
		return 0
	}

	return occurances(bytesInDirection(data, i, DIAG_DOWN_LEFT))
}

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}

	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))

	sum := 0
	for i := 0; i < len(data); i++ {
		r := findDiagR(data, i)
		l := findDiagL(data, i+2)
		if r > 0 && l > 0 {
			sum++
		}
	}

	fmt.Println(sum)
}
