package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func help() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Uppercase\n2. Lowercase\n3. Capitalize Words\n")
	fmt.Println("4. Title Case\n5. Snake Case\n6. Reverse Words\n")
	fmt.Println("7. Count Text\n8. Palindrome Check\n9. Exit\n")
	fmt.Println("----------PROGRAM AUTOMATICALLY CONTINUES----------\n")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var opStr string
		var option int

		fmt.Println("-----------------------------------------------")
		fmt.Println("----- CodeCrafters — Operation Gopher Protocol")
		fmt.Println("----- Module: String Transformer")
		fmt.Println("----- Author: Eunice Obeko")
		fmt.Println("----- Squad: The Gophers")
		fmt.Println("-----------------------------------------------")
		fmt.Println("                                                                         ")

		fmt.Println("-------------------------------------------------------------------------")
		fmt.Println("------------ WELCOME TO YOUR STRING TRANSFORMER ------------")
		fmt.Println("-------------------------------------------------------------------------")

		fmt.Print("Enter Your Choice:\n1. Uppercase\n2. Lowercase\n3. Capitalize Words\n4. Title Case\n5. Snake Case\n6. Reverse Words\n7. Count Text\n8. Palindrome Check\n9. Exit\n")
		fmt.Print("Now, Select An Option: ")

		opStr, _ = reader.ReadString('\n')
		opStr = strings.TrimSpace(opStr)

		option, err := strconv.Atoi(opStr)
		if err != nil || option < 1 || option > 9 {
			fmt.Println("Invalid Option!")
			help()
			continue
		}

		if option == 9 {
			fmt.Println("Goodbye!")
			return
		}

		fmt.Print("Enter Text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "" {
			fmt.Println("Empty input!")
			continue
		}

		var output string

		if option == 1 {
			output = toUpper(text)
		} else if option == 2 {
			output = toLower(text)
		} else if option == 3 {
			output = capFirst(text)
		} else if option == 4 {
			output = titleCase(text)
		} else if option == 5 {
			output = snakeCase(text)
		} else if option == 6 {
			output = reverseCase(text)
		} else if option == 7 {
			output = countText(text)
		} else if option == 8 {
			output = palindromeCheck(text)
		}

		fmt.Println("Result:\n", output)
	}
}

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
