package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := convFile("./input.txt")
	prior := genPriority()
	totalPriorityItems := sumPriority(input, prior)
	totalPriorityItemsGroup := sumPriorityGroup(input, prior)
	fmt.Println(totalPriorityItems)
	fmt.Println(totalPriorityItemsGroup)
}

func sumPriority(rucksacks []string, prior map[byte]int) int {
	res := 0

	for _, rucksack := range rucksacks {
		m := len(rucksack) >> 1
		first := rucksack[:m]
		second := rucksack[m:]
		items := findCommon(first, second)

		for item := range items {
			res += prior[item]
		}
	}

	return res
}

func sumPriorityGroup(rucksacks []string, prior map[byte]int) int {
	res := 0

	for i := 0; i < len(rucksacks); i += 3 {
		first := rucksacks[i]
		second := rucksacks[i+1]
		third := rucksacks[i+2]

		m1 := findCommon(first, second)
		m2 := findCommon(second, third)

		var badge byte

		for item := range m1 {
			if m2[item] {
				badge = item
				break
			}
		}

		if badge == 0 {
			continue
		}

		res += prior[badge]
	}

	return res
}

func findCommon(s1, s2 string) map[byte]bool {
	res, m := make(map[byte]bool), make(map[byte]bool)

	for i := 0; i < len(s1); i++ {
		char := s1[i]
		m[char] = true
	}

	for i := 0; i < len(s2); i++ {
		char := s2[i]

		if m[char] {
			res[char] = true
		}
	}

	return res
}

func genPriority() map[byte]int {
	priority := make(map[byte]int)

	for i := 0; i < 26; i++ {
		lower := 'a' + i
		upper := 'A' + i
		priority[byte(lower)] = i + 1
		priority[byte(upper)] = i + 27
	}

	return priority
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
