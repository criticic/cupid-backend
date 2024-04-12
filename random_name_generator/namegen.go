package random_name_generator

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

var ADJECTIVE_FILE = "random_name_generator/adjectives.txt"
var ANIMAL_FILE = "random_name_generator/animals.txt"

// Generate a random name using the adjectives and animals files.
func GenerateName() string {
	adjective := GetRandomLine(ADJECTIVE_FILE)
	animal := GetRandomLine(ANIMAL_FILE)

	username := adjective + "-" + animal
	username = strings.ReplaceAll(username, " ", "-")

	return username
}

// Get a random line from a file.
func GetRandomLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines[rand.Intn(len(lines))]
}
