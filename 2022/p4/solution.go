package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := convFile("./input.txt")
	sections := convSections(input)
	sectionsWithin := countWithin(sections)
	sectionsOverlap := countOverlapping(sections)
	fmt.Println(sectionsWithin)
	fmt.Println(sectionsOverlap)
}

func countWithin(sections [][4]int) int {
	res := 0

	for _, s := range sections {
		if within(s[0], s[1], s[2], s[3]) ||
			within(s[2], s[3], s[0], s[1]) {
			res++
		}
	}

	return res
}

func countOverlapping(sections [][4]int) int {
	res := 0

	for _, s := range sections {
		if overlap(s[0], s[1], s[2], s[3]) {
			res++
		}
	}

	return res
}

func convSections(lines []string) [][4]int {
	res := make([][4]int, len(lines))

	for i, line := range lines {
		sections := strings.Split(line, ",")
		first := strings.Split(sections[0], "-")
		second := strings.Split(sections[1], "-")
		res[i][0], _ = strconv.Atoi(first[0])
		res[i][1], _ = strconv.Atoi(first[1])
		res[i][2], _ = strconv.Atoi(second[0])
		res[i][3], _ = strconv.Atoi(second[1])
	}

	return res
}

func within(s1, e1, s2, e2 int) bool {
	return s2 >= s1 && e2 <= e1
}

func overlap(s1, e1, s2, e2 int) bool {
	return s2 <= e1 && s1 <= e2
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
