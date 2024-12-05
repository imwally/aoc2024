package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const INPUT_FILE = "input.txt"
const KEY_WORD = "XMAS"
const KEY_WORD_REVERSED = "SAMX"

const INPUT_COL_WIDTH = 140

func occurances(s string) int {
	sum := 0
	if s == KEY_WORD {
		sum++
	}
	if s == KEY_WORD_REVERSED {
		sum++
	}

	return sum
}

func findHorizontal(data []byte, i int) int {
	if i%INPUT_COL_WIDTH > INPUT_COL_WIDTH-len(KEY_WORD) {
		return 0
	}

	if i+len(KEY_WORD) > len(data) {
		return 0
	}

	horizontal := string(data[i : i+len(KEY_WORD)])

	return occurances(horizontal)
}

func findVertical(data []byte, i int) int {
	if i+INPUT_COL_WIDTH*3 > len(data)-1 {
		return 0
	}

	vertical := make([]byte, len(KEY_WORD))
	vertical[0] = data[i]
	vertical[1] = data[i+INPUT_COL_WIDTH]
	vertical[2] = data[i+INPUT_COL_WIDTH*2]
	vertical[3] = data[i+INPUT_COL_WIDTH*3]

	verticalString := string(vertical)

	return occurances(verticalString)
}

func findDiagR(data []byte, i int) int {
	if i%INPUT_COL_WIDTH > INPUT_COL_WIDTH-len(KEY_WORD) {
		return 0
	}

	if i+3+(INPUT_COL_WIDTH*3) > len(data)-1 {
		return 0
	}

	diag := make([]byte, len(KEY_WORD))
	diag[0] = data[i]
	diag[1] = data[i+1+INPUT_COL_WIDTH]
	diag[2] = data[i+2+(INPUT_COL_WIDTH*2)]
	diag[3] = data[i+3+(INPUT_COL_WIDTH*3)]

	diagString := string(diag)

	return occurances(diagString)
}

func findDiagL(data []byte, i int) int {
	if i%INPUT_COL_WIDTH < len(KEY_WORD)-1 {
		return 0
	}

	if i > len(data)-(INPUT_COL_WIDTH*3) {
		return 0
	}

	diag := make([]byte, len(KEY_WORD))
	diag[0] = data[i]
	diag[1] = data[i-1+INPUT_COL_WIDTH]
	diag[2] = data[i-2+(INPUT_COL_WIDTH*2)]
	diag[3] = data[i-3+(INPUT_COL_WIDTH*3)]

	diagString := string(diag)

	return occurances(diagString)
}

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}

	data = bytes.ReplaceAll(data, []byte("\n"), []byte(""))

	sum := 0
	for i := 0; i < len(data); i++ {
		sum += findHorizontal(data, i)
		sum += findVertical(data, i)
		sum += findDiagR(data, i)
		sum += findDiagL(data, i)
	}

	fmt.Println(sum)
}
