package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := convFile("./input.txt")
	maxElf := maxCalElf(input)
	elves := convToElves(input)
	topThree := sumTopElves(elves, 3)

	fmt.Println("This is the most calories carried by one elf:")
	fmt.Println(maxElf)
	fmt.Println("This is the total calories of the top three elves:")
	fmt.Println(topThree)
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

func sumTopElves(elves []int, k int) int {
	sort.Ints(elves)

	res, n := 0, len(elves)

	for i := n - 1; i >= 0 && k > 0; i, k = i-1, k-1 {
		res += elves[i]
	}

	return res
}

func convToElves(snacks []string) []int {
	elves, curElf := []int{}, 0

	for _, snack := range snacks {
		if len(snack) == 0 {
			elves = append(elves, curElf)
			curElf = 0
		} else {
			calories, err := strconv.Atoi(snack)

			if err != nil {
				log.Fatal(err)
			}

			curElf += calories
		}
	}

	return elves
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
