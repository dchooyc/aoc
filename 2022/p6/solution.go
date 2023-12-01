package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := convFile("./input.txt")
	packetMarker := packMark(input, 4)
	startMessage := packMark(input, 14)
	fmt.Println(packetMarker)
	fmt.Println(startMessage)
}

func packMark(lines []string, k int) int {
	s := lines[0]
	letters := []byte(s[:k-1])

	for i := k - 1; i < len(s); i++ {
		letters = append(letters, s[i])

		if check(letters) {
			return i + 1
		}

		letters = letters[1:]
	}

	return len(s)
}

func check(letters []byte) bool {
	m := make(map[byte]bool)

	for _, letter := range letters {
		if m[letter] {
			return false
		}
		m[letter] = true
	}

	return true
}

func convFile(src string) []string {
	res := []string{}

	file, err := os.Open(src)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
