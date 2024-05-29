package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func CheckIfFileExists(filePath string) bool {
	fmt.Printf("Check if file: %v exists. \n", filePath)
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func DeleteFile(filePath string) {
	fmt.Printf("File needs to be recreated; file: %v \n", filePath)

	err := os.Remove(filePath)
	if err != nil {
		log.Fatalf("Error on deletion; file: %v, with error: %s \n", filePath, err)
	} else {
		fmt.Printf("Deleted file; file: %v \n", filePath)
	}
}

func OpenFile(filePath string) *os.File {
	fmt.Printf("Open file for scanning: %v \n", filePath)
	inputFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open input file: %v, with error: %s \n", filePath, err)
	}
	return inputFile
}

func CreateFile(filePath string) *os.File {
	fmt.Printf("Creating file for output: %v \n", filePath)
	outputFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create the output file: %v, with error: %s \n", filePath, err)
	} else {
		fmt.Printf("File created for output: %v \n", filePath)
	}
	return outputFile
}

func ScanLinesFromFile(filePath string, scanner *bufio.Scanner) []string {
	var lines []string
	fmt.Printf("Started scanning file: %v \n", filePath)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	fmt.Printf("Finished scanning file: %v \n", filePath)
	return lines
}

func WriteLinesToFile(filePath string, lines []string, writer *bufio.Writer) {
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Failed to write to output file: %v, with error: %s, at line: %v \n", filePath, err, line)
		}
	}
}

func EnsureBuffersSuccess(inputFilePath, outputFilePath string, scanner *bufio.Scanner, writer *bufio.Writer) {
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file: %v, with error: %s \n", inputFilePath, err)
	}
	if err := writer.Flush(); err != nil {
		log.Fatalf("Failed to flush output file: %v, with error: %s \n", outputFilePath, err)
	}
	fmt.Printf("Process complete for file: %v \n", outputFilePath)
}
