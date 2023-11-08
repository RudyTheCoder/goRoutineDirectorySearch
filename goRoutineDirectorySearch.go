package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func readFile(target string, filePath string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group when the function exits
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), target) {
			fmt.Printf("%s (%d): %s\n", filePath, line, scanner.Text())
		}
		line++
	}
}

func main() {
	target := "is"
	folderName := "homework"

	// Change the current working directory
	err := os.Chdir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			wg.Add(1) // Increment the wait group counter
			go readFile(target, f.Name(), &wg)
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines have finished")

	// Sleep to give other threads time to finish (optional)
	time.Sleep(100 * time.Millisecond)
}
