package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	sl := []int{}
	for _, x := range text {
		m, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("failed to convert")
		}
		sl = append(sl, m)
	}
	fmt.Println(sl)
	for _, x := range sl {
		for _, y := range sl {
			for _, z := range sl {
				if x+y+z == 2020 {
					fmt.Println(x * y * z)
				}
			}
		}
	}

}
