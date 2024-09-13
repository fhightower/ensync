package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func readKeys(filename string) (map[string]string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read(filename)
	return myEnv, err
}

func readEnvFiles(dirname string) (map[string]string, map[string]string) {
	envKeys, envErr := readKeys(dirname + "/.env")
	envExampleKeys, envExampleErr := readKeys(dirname + "/.env.example")

	if envErr != nil {
		log.Fatalf("Error reading .env file: %s", envErr)
	}

	if envExampleErr != nil {
		log.Fatalf("Error reading .env.example file: %s", envExampleErr)
	}

	return envKeys, envExampleKeys

}

func compareKeys(envKeys map[string]string, envExampleKeys map[string]string) {
	for key := range envKeys {
		if _, ok := envExampleKeys[key]; !ok {
			fmt.Println("Key", key, "is missing in .env.example")
		}
	}

	for key := range envExampleKeys {
		if _, ok := envKeys[key]; !ok {
			fmt.Println("Key", key, "is missing in .env")
		}
	}
}

func compareFiles(dirname string) {
	envKeys, envExampleKeys := readEnvFiles(dirname)
	compareKeys(envKeys, envExampleKeys)
}

func processIfPossible(dirname string) bool {
	fmt.Println("Processing", dirname)
	envExists := fileExists(dirname + "/.env")
	envExampleExists := fileExists(dirname + "/.env.example")
	processed := false

	if envExists && envExampleExists {
		compareFiles(dirname)
		processed = true
	}
	return processed
}

func scanDirs(basepath string) {
	dirs, err := ioutil.ReadDir(basepath)
	if err != nil {
		log.Fatalf("Error reading current directory: %s", err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			processIfPossible(basepath + "/" + dir.Name())
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go [directory]")
	}

	path := os.Args[1]
	processed := processIfPossible(path)

	if !processed {
		scanDirs(path)
	}
}
