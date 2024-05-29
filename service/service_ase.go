package service

import (
	"fmt"
	"log"
	"strings"
)

func ChangeLinesASESpecialities(lines []string) []string {
	var changedLines []string
	for _, line := range lines {
		// 1|Comert, Banca, Asigurari|Licenta|3|IF (la zi)|Bucuresti|2|1919|1926
		fields := strings.Split(line, "|")
		if len(fields) >= 3 {
			changedLine := fmt.Sprintf("%v|%v", fields[1], fields[2])
			changedLines = append(changedLines, changedLine)
		} else {
			log.Fatalf("invalid line: %v", line)
		}
	}
	return changedLines
}

func ChangeLinesASEProgram(lines []string, program string) []string {
	var changedLines []string
	for _, line := range lines {
		// Planificare si Cibernetica Economica|Licenta
		if strings.Contains(line, program) {
			changedLines = append(changedLines, line)
		}
	}
	return changedLines
}
