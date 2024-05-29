package main

import (
	"bufio"
	"log"
	"txtparser/config"
	"txtparser/service"
	"txtparser/utils"
)

func main() {

	inputFilePath := config.INPUT_FILE_PATH
	outputFilePath := config.OUTPUT_FILE_PATH
	recreateOnStart := config.RECREATE_ON_START

	fileExists := utils.CheckIfFileExists(outputFilePath)
	if fileExists && recreateOnStart {
		utils.DeleteFile(outputFilePath)
	} else if fileExists && !recreateOnStart {
		log.Fatalf("The file: %v exists and recreation is set to: %v \n", outputFilePath, recreateOnStart)
	}

	inputFile := utils.OpenFile(inputFilePath)
	defer inputFile.Close()

	outputFile := utils.CreateFile(outputFilePath)
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	inputLines := utils.ScanLinesFromFile(inputFilePath, scanner)
	changedLines := service.ChangeLinesToIndexHumanAndChar(inputLines, "|") // Use the desired function from service package
	utils.WriteLinesToFile(outputFilePath, changedLines, writer)

	utils.EnsureBuffersSuccess(inputFilePath, outputFilePath, scanner, writer)
}
