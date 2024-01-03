package main

import (
	"bufio"
	"fmt"
	"os"
)

var request = &Request{}

func main() {
	if openAiToken == "" {
		fmt.Println("OPENAI_TOKEN environment variable is not set")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("You: ")
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Println("GPT: " + request.send(msg))
		fmt.Println("---------------------------------")
		fmt.Print("You: ")
	}

}
