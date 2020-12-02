package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	i := 0
	for _, x := range text {
		if check(splir(x)) {
			i++
		}
	}
	fmt.Println(i)
}

func splir(x string) (int, int, string, string) {
	re := regexp.MustCompile(`^(\d+)-(\d+)\s(.):\s(.*)$`)
	// re := regexp.MustCompile(`(\d+-\d+)`)
	sl := re.FindStringSubmatch(x)
	min, _ := strconv.Atoi(sl[1])
	max, _ := strconv.Atoi(sl[2])

	return min, max, sl[3], sl[4]
}

func check(min int, max int, value string, pv string) bool {
	ant := strings.Count(pv, value)
	return ant >= min && ant <= max
}
