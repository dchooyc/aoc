package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvFile(src string) []string {
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

func Stia(s string) []int {
	vals := strings.Split(s, " ")
	res := make([]int, len(vals))

	for i, val := range vals {
		num, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println("Error converting string to int")
			return []int{}
		}

		res[i] = num
	}

	return res
}

func Stia64(s string) []int64 {
	vals := strings.Split(s, " ")
	res := make([]int64, len(vals))

	for i, val := range vals {
		num, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println("Error converting string to int64")
			return []int64{}
		}

		res[i] = int64(num)
	}

	return res
}
