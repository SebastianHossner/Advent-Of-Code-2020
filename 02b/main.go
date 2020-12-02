package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func splir(x string) (int, int, byte, string) {
	re := regexp.MustCompile(`^(\d+)-(\d+)\s(.):\s(.*)$`)
	sl := re.FindStringSubmatch(x)
	min, _ := strconv.Atoi(sl[1])
	max, _ := strconv.Atoi(sl[2])
	return min, max, sl[3][0], sl[4]
}

func check(min int, max int, value byte, pv string) bool {
	return (pv[min-1] == value || pv[max-1] == value) && !(pv[min-1] == value && pv[max-1] == value)
}
