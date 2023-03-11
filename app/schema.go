package main

import "github.com/sashabaranov/go-openai"

type BaseRs struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type ChatSessionRq struct {
	Sid int    `json:"sid"`
	Uid string `json:"uid"`
}

type ChatSessionRs struct {
	BaseRs
	Sid int
}

type ChatgptRq struct {
	Sid     int                            `json:"sessionId,omitempty"`
	Content string                         `json:"content,omitempty"`
	Model   string                         `json:"model,omitempty"`
	H       []openai.ChatCompletionMessage `json:"h,omitempty"`
}

type ChatgptRs struct {
	BaseRs
	Content string `json:"content,omitempty"`
}
