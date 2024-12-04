package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const INPUT_FILE = "input.txt"

func mul(exp string) int {
	s := strings.Split(exp, ",")

	a, _ := strconv.Atoi(s[0][4:])
	b, _ := strconv.Atoi(s[1][:len(s[1])-1])

	return a * b
}

func findMul(data string) int {
	sum := 0
	for i := 0; i < len(data)-4; i++ {
		if data[i:i+4] == "mul(" {
			j := i + 4
			for ; data[j] != ')'; j++ {
			}

			token := data[i : j+1]
			re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
			found := re.FindAll([]byte(token), -1)

			if len(found) > 0 {
				exp := string(found[0])
				sum += mul(exp)
				i = j
			}
		}
	}

	return sum
}

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}

	sum := 0

	// Get batch of mul() before first do()
	for i := 0; i < len(data)-4; i++ {
		if string(data[i:i+4]) == "do()" {
			firstBatch := string(data[0:i])
			sum += findMul(firstBatch)
			break
		}
	}

	// Parse rest
	for i := 0; i < len(data)-4; i++ {
		if string(data[i:i+4]) == "do()" {
			j := i + 4
			for ; j < len(data)-7; j++ {
				if string(data[j:j+7]) == "don't()" {
					doString := string(data[i : j+7])
					sum += findMul(doString)
					break
				}
			}
			i = j + 7
		}
	}

	fmt.Println(sum)
}
