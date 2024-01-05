package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var request = &Request{}
var spinner = newSpinner()

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
	print("Set message system, ex. \"You are a backend developer.\" (optional): ")
	scanner.Scan()
	if scanner.Text() != "" {
		messages = append(messages, Message{Role: "system", Content: scanner.Text()})
	}
	print("You: ")
	for scanner.Scan() {
		msg := scanner.Text()

		resp := make(chan Message, 1)
		go func() {
			resp <- request.send(msg)
		}()

	SelectLoop:
		for {
			select {
			case respMsg := <-resp:
				print("GPT: ")
				for _, char := range respMsg.Content {
					print(string(char))
					time.Sleep(25 * time.Millisecond)
				}
				fmt.Println("\n---------------------------------")
				print("You: ")
				break SelectLoop
			default:
				fmt.Println(spinner.spin())
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("\033[1A")
			}
		}
	}
}
