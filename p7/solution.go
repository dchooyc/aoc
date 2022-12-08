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
	m := genMap(input)
	smallSum := findSmall(m)
	targetDelete := findDelete(m)
	fmt.Println(smallSum)
	fmt.Println(targetDelete)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findDelete(m map[string]int) int {
	target := m["/"] - 40000000
	res := 70000000

	for _, val := range m {
		if val >= target {
			res = min(res, val)
		}
	}

	return res
}

func findSmall(m map[string]int) int {
	res := 0

	for _, val := range m {
		if val <= 100000 {
			res += val
		}
	}

	return res
}

func genMap(input []string) map[string]int {
	path, m := []string{}, make(map[string]int)

	for _, line := range input {
		a := strings.Split(line, " ")

		if a[0] == "$" && a[1] == "cd" {
			if a[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, a[2])
			}
		} else if a[0] != "$" && a[0] != "dir" {
			size, _ := strconv.Atoi(a[0])

			for i := len(path); i > 0; i-- {
				curPath := strings.Join(path[:i], "/")
				m[curPath] += size
			}
		}
	}

	return m
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
