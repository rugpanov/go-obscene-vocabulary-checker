package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := scanLine()

	tabooWords := readTabooWords(fileName)

	for true {
		nextSentence := scanLine()
		if nextSentence == "exit" {
			break
		}
		allWords := strings.Split(nextSentence, " .")
		for i, word := range allWords {
			allWords[i] = applyCensorship(&tabooWords, word)
		}
		fmt.Println(strings.Join(allWords[:], " "))
	}

	fmt.Println("Bye!")
}

func scanLine() (scanLine string) {
	/*
		alternative:
		fmt.Scan(&scanLine)
		return scanLine
	*/
	reader := bufio.NewReader(os.Stdin)
	scanLine, _ = reader.ReadString('\n')
	scanLine = strings.TrimSpace(scanLine)
	return scanLine
}

func readTabooWords(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, strings.ToLower(scanner.Text()))
	}

	return result
}

func applyCensorship(taboos *[]string, word string) string {
	theWordLowercase := strings.ToLower(word)

	shouldCensor := false
	for _, tabooWord := range *taboos {
		if tabooWord == theWordLowercase {
			shouldCensor = true
		}
	}

	if shouldCensor {
		strLen := len([]rune(word))
		return strings.Repeat("*", strLen)
	} else {
		return word
	}
}
