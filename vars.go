package main

import "os"

var (
	openAiToken = os.Getenv("OPENAI_TOKEN")
	openAiUrl   = "https://api.openai.com/v1/chat/completions"
	messages    = []Message{}
	model       = "gpt-3.5-turbo"
)
