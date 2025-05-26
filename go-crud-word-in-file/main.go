package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName = "word.txt"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter operation (c: create, r: read, u: update, d: delete): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(strings.ToLower(operation))

	switch operation {
	case "c":
		createWord(reader)
	case "r":
		readWord()
	case "u":
		updateExistingWordTo(reader)
	case "d":
		deleteWord()
	default:
		fmt.Println("Invalid operation. Use c, r, u, or d.")
	}
}

func createWord(reader *bufio.Reader) {
	if wordExists() {
		fmt.Println("A word already exists. Use update (u) to modify it.")
		return
	}
	fmt.Print("Enter word to create: ")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	err := os.WriteFile(fileName, []byte(word), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Word created successfully.")
}

func readWord() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No word found. File doesn't exist.")
		} else {
			fmt.Println("Error reading file:", err)
		}
		return
	}
	word := strings.TrimSpace(string(data))
	if word == "" {
		fmt.Println("No word stored.")
	} else {
		fmt.Println("Stored word:", word)
	}
}

func updateExistingWordTo(reader *bufio.Reader) {
	if !wordExists() {
		fmt.Println("No existing word to update. Use create (c) instead.")
		return
	}
	fmt.Print("Enter new word to update: ")
	newWord, _ := reader.ReadString('\n')
	newWord = strings.TrimSpace(newWord)
	err := os.WriteFile(fileName, []byte(newWord), 0644)
	if err != nil {
		fmt.Println("Error updating file:", err)
		return
	}
	fmt.Println("Word updated successfully.")
}

func deleteWord() {
	err := os.WriteFile(fileName, []byte(""), 0644)
	if err != nil {
		fmt.Println("Error deleting word:", err)
		return
	}
	fmt.Println("Word deleted successfully.")
}

// helper function
func wordExists() bool {
	data, err := os.ReadFile(fileName)
	if err != nil || strings.TrimSpace(string(data)) == "" {
		return false
	}
	return true
}
