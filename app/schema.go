package main

type BaseRs struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type ChatSessionRq struct {
	Sid int    `json:"sid,omitempty"`
	Uid string `json:"uid,omitempty"`
}

type ChatSessionRs struct {
	BaseRs
	Sid int
}

type ChatgptRq struct {
	Sid     int    `json:"sessionId,omitempty"`
	Content string `json:"content,omitempty"`
}

type ChatgptRs struct {
	BaseRs
	Content string `json:"content,omitempty"`
}
