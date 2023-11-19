package filesystem

import (
	"bufio"
	"log"
	"os"
)

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func DeleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadLine(filename string, line int) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 0

	for scanner.Scan() {
		currentLine++
		if currentLine == line {
			return scanner.Text()
		}
	}

	return ""
}

func WriteFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}
