package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Payload Payload
}

func (r *Request) send(msg string) string {
	r.appendMessage(Message{Role: "user", Content: msg})
	return r.performPostRequest()
}

func (r *Request) performPostRequest() string {
	r.Payload = Payload{Model: model, Messages: messages}
	payloadBytes, _ := json.Marshal(r.Payload)

	req, err := http.NewRequest("POST", openAiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Error with status code: " + resp.Status)
	}
	msg := r.extractMessage(resp.Body)
	r.appendMessage(msg)

	return msg.Content
}

func (r *Request) extractMessage(body io.Reader) Message {
	var resp ChatResponse
	err := json.NewDecoder(body).Decode(&resp)
	if err != nil {
		panic(err)
	}
	return resp.Choices[0].Message
}

func (r *Request) appendMessage(msg Message) {
	messages = append(messages, msg)
}
