package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const INPUT_FILE = "test1.txt"

func mul(exp string) int {
	s := strings.Split(exp, ",")

	a, _ := strconv.Atoi(s[0][4:])
	b, _ := strconv.Atoi(s[1][:len(s[1])-1])

	return a * b
}

func main() {
	data, err := os.ReadFile(INPUT_FILE)
	if err != nil {
		log.Println(err)
	}

	sum := 0
	for i := 0; i < len(data)-4; i++ {
		if string(data[i:i+4]) == ("mul(") {
			j := i + 4
			for ; data[j] != ')'; j++ {
			}

			token := data[i : j+1]
			re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
			found := re.FindAll(token, -1)

			if len(found) > 0 {
				exp := string(found[0])
				sum += mul(exp)
				i = j
			}
		}
	}

	fmt.Println(sum)
}
