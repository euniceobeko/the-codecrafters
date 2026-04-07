package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*fmt.Println("-----------------------------------------------")
	fmt.Println("----- CodeCrafters — Operation Gopher Protocol")
	fmt.Println("----- Module: File Pipeline")
	fmt.Println("----- Author: Eunice Obeko")
	fmt.Println("----- Squad: The Gophers")
	fmt.Println("-----------------------------------------------")*/

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if input == output {
		fmt.Println("Input and output cannot be the same file")
		return
	}

	inFile, err := os.Open(input)
	if err != nil {
		fmt.Println("File not found:", input)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Cannot create output file")
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)

	linesRead := 0
	linesWritten := 0
	linesRemoved := 0

	writer.WriteString("Gopher's Sentinel Field Report - Processed\n\n")

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		line = strings.TrimSpace(line)

		if line == "" || strings.Trim(line, "-") == "" {
			linesRemoved++
			continue
		}

		line = strings.ReplaceAll(line, "TODO:", "✦ ACTION:")

		if line == strings.ToLower(line) {
			line = strings.ToUpper(line)
		}

		if line == strings.ToUpper(line) {
			words := strings.Fields(strings.ToLower(line))
			for i := 0; i < len(words); i++ {
				words[i] = strings.ToUpper(words[i][:1]) + words[i][1:]
			}
			line = strings.Join(words, " ")
		}

		if strings.Contains(line, "REVERSE") {
			words := strings.Fields(line)
			for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
				words[i], words[j] = words[j], words[i]
			}
			line = strings.Join(words, " ")
		}

		linesWritten++
		writer.WriteString(fmt.Sprintf("%d. %s\n", linesWritten, line))
	}

	writer.Flush()

	fmt.Println("\nSUMMARY")
	fmt.Println("Lines read   :", linesRead)
	fmt.Println("Lines written:", linesWritten)
	fmt.Println("Lines removed:", linesRemoved)
}