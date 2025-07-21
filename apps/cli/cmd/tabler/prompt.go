package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func confirmDeletion(taskTitle string, reader io.Reader) bool {
	fmt.Printf("Delete task \"%s\"? (y/N): ", taskTitle)
	
	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		return strings.ToLower(input) == "y"
	}
	
	return false
}