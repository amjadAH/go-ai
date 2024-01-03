package main

import "os"

var (
	openAiToken = os.Getenv("OPENAI_TOKEN")
	openAiUrl   = "https://api.openai.com/v1/chat/completions"
	messages    = []Message{{Role: "system", Content: "You are a backend developer, specializing in PHP, and willing to learn new technologies"}}
	model       = "gpt-3.5-turbo"
)
