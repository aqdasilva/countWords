package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var File string

func input() {
	fmt.Println("Enter Text File: ")

	//takes input from user
	fmt.Scanln(&File)

	//prints file contents
	data, err := os.Open(File)

	if err != nil {
		log.Panicf("failed, you suck try again: %s", err)
	}

	reader := bufio.NewReader(data)
	counter := make(map[string]int)
	for {
		line, _ := reader.ReadString('\n')

		fields := strings.Fields(line)

		for _, word := range fields {
			word = strings.ToLower(word)
			counter[word]++
		}
		if line == "" {
			break
		}
	}

	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return counter[words[i]] > counter[words[j]]
	})

	for _, word := range words {
		fmt.Printf("%v %v\n", word, counter[word])
	}
}

func main() {
	input()
}
