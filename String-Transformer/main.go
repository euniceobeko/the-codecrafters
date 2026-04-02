package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const maxHistory = 5

type HistoryEntry struct {
	Command string
	Input   string
	Output  string
}

func main() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("----- CodeCrafters — Operation Gopher Protocol")
	fmt.Println("----- Module: String Transformer")
	fmt.Println("----- Author: Eunice Obeko")
	fmt.Println("----- Squad: The Gophers")
	fmt.Println("-----------------------------------------------")

	scanner := bufio.NewScanner(os.Stdin)
	history := []HistoryEntry{}

	validCommands := []string{"upper", "lower", "cap", "title", "snake", "reverse", "count", "palindrome", "history", "exit"}

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			fmt.Println("\nShutting down String Transformer. Goodbye.")
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		fields := strings.Fields(input)
		command := strings.ToLower(fields[0])
		text := strings.TrimSpace(strings.Join(fields[1:], " "))

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		if !contains(validCommands, command) {
			fmt.Printf("✗ Unknown command: \"%s\"\n", command)
			fmt.Print("  Valid commands:")
			for _, c := range validCommands[:len(validCommands)-1] {
				fmt.Print(" ", c+",")
			}
			fmt.Println(" exit")
			continue
		}

		needsText := map[string]bool{
			"upper": true, "lower": true, "cap": true,
			"title": true, "snake": true, "reverse": true,
			"count": true, "palindrome": true,
		}

		if needsText[command] && text == "" {
			fmt.Printf("✗ No text provided. Usage: %s <text>\n", command)
			continue
		}

		var output string

		switch command {
		case "upper":
			output = toUpper(text)
		case "lower":
			output = toLower(text)
		case "cap":
			output = capFirst(text)
		case "title":
			output = titleCase(text)
		case "snake":
			output = snakeCase(text)
		case "reverse":
			output = reverseCase(text)
		case "count":
			output = countText(text)
		case "palindrome":
			output = palindromeCheck(text)
		case "history":
			if len(history) == 0 {
				fmt.Println("No history yet.")
			} else {
				for i, h := range history {
					fmt.Printf("%d. [%s] \"%s\" → %s\n", i+1, h.Command, h.Input, h.Output)
				}
			}
			continue
		}

		if command != "history" {
			fmt.Println("→", output)
			// Add to history
			history = append(history, HistoryEntry{Command: command, Input: text, Output: output})
			if len(history) > maxHistory {
				history = history[1:]
			}
		}
	}
}

// -------------------- Utility --------------------

func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

// -------------------- String Transform Functions --------------------

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func capFirst(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func titleCase(s string) string {
	words := strings.Fields(strings.ToLower(s))
	small := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true, "or": true,
		"for": true, "nor": true, "on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true, "is": true, "it": true,
	}

	for i, w := range words {
		if i == 0 || !small[w] {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

func reverseCase(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		reversed := ""
		for _, ch := range w {
			reversed = string(ch) + reversed
		}
		words[i] = reversed
	}
	return strings.Join(words, " ")
}

func snakeCase(s string) string {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	for i, w := range words {
		clean := ""
		for _, ch := range w {
			if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_' {
				clean += string(ch)
			}
		}
		words[i] = clean
	}
	return strings.Join(words, "_")
}

// -------------------- Bonus Commands --------------------

func countText(s string) string {
	totalChars := len(s)
	totalWords := len(strings.Fields(s))
	totalLetters := 0
	totalSpaces := 0
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			totalLetters++
		} else if unicode.IsSpace(ch) {
			totalSpaces++
		}
	}
	return fmt.Sprintf("Total characters: %d\nTotal letters: %d\nTotal words: %d\nTotal spaces: %d",
		totalChars, totalLetters, totalWords, totalSpaces)
}

func palindromeCheck(s string) string {
	clean := ""
	for _, ch := range s {
		if !unicode.IsSpace(ch) {
			clean += strings.ToLower(string(ch))
		}
	}

	length := len(clean)
	for i := 0; i < length/2; i++ {
		if clean[i] != clean[length-1-i] {
			return fmt.Sprintf("✗ \"%s\" is not a palindrome.", s)
		}
	}
	return fmt.Sprintf("✓ \"%s\" is a palindrome!", s)
}
