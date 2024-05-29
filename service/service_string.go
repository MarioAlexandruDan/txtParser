package service

import (
	"fmt"
	"log"
	"strings"
)

func ChangeLinesForSingleAppearance(lines []string) []string {

	set := make(map[string]bool)
	for _, line := range lines {
		set[line] = true
	}

	changedLines := make([]string, 0, len(set))
	for line := range set {
		changedLines = append(changedLines, line)
	}

	return changedLines
}

func ChangeLinesForCharApperance(lines []string, char string) []string {

	if len(char) != 1 {
		log.Fatalf("ChangeLinesForCharApperance failed, the provided string is not a char: %v \n", char)
	}

	var changedLines []string
	for _, line := range lines {
		fields := strings.Split(line, char)
		changedLines = append(changedLines, fields[0])
	}

	return changedLines
}

func ChangeLinesToIndexAndChar(lines []string, char string) []string {

	if len(char) != 1 {
		log.Fatalf("ChangeLinesToIndexAndChar failed, the provided string is not a char: %v \n", char)
	}

	var changedLines []string
	for index, line := range lines {
		changedLine := fmt.Sprintf("%v%v%v", index+1, char, line)
		changedLines = append(changedLines, changedLine)
	}

	return changedLines
}

func ChangeLinesToIndexHumanAndChar(lines []string, char string) []string {

	if len(char) != 1 {
		log.Fatalf("ChangeLinesToIndexAndChar failed, the provided string is not a char: %v \n", char)
	}

	var changedLines []string
	for index, line := range lines {
		changedLine := fmt.Sprintf("%v%v%v", index+1, char, line)
		changedLines = append(changedLines, changedLine)
	}

	return changedLines
}
