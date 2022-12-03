package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := convFile("./input.txt")
	maxElf := maxCalElf(input)
	fmt.Println(maxElf)
}

func maxCalElf(snacks []string) int {
	maxElf, curElf := 0, 0

	for _, snack := range snacks {
		if len(snack) == 0 {
			if curElf > maxElf {
				maxElf = curElf
			}

			curElf = 0
		} else {
			calories, err := strconv.Atoi(snack)

			if err != nil {
				log.Fatal(err)
			}

			curElf += calories
		}
	}

	return maxElf
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
