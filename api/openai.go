package api

import (
	"context"

	"github.com/rs/zerolog/log"

	openai "github.com/sashabaranov/go-openai"
	// openaiConfig "github.com/sashabaranov/go-openai/config"
)

type OpenAIRq struct {
	Model   string
	Content string
	H       []openai.ChatCompletionMessage
}

type OpenAIRs struct {
	Content string       `json:"content,omitempty"`
	Usage   openai.Usage `json:"usage,omitempty"`
}

type OpenAIType ClientProcessor[*OpenAIRq, *OpenAIRs]

type OpenAI struct {
	c *openai.Client
}

var DefaultChatCompletionMessage = []openai.ChatCompletionMessage{{
	Role:    openai.ChatMessageRoleSystem,
	Content: "You are a helpful assistant.",
}}

func (o *OpenAI) Process(ctx context.Context, rq *OpenAIRq) (rs *OpenAIRs, err error) {
	var messages = make([]openai.ChatCompletionMessage, 4)
	messages = append(messages, rq.H...)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: rq.Content,
	})
	log.Debug().Interface("rq", rq).Msg("print rq")
	resp, err := o.c.CreateChatCompletion(
		// context.Background(),
		ctx,
		openai.ChatCompletionRequest{
			Model:    rq.Model,
			Messages: messages,
		},
	)
	if err != nil {
		return rs, err
	}
	rs.Content = resp.Choices[0].Message.Content
	rs.Usage = resp.Usage
	return rs, nil
}

func newOpenAI(cfg ClientCfg) OpenAIType {
	cli := openai.NewClient(cfg.Ak)
	// cli.config
	return &OpenAI{cli}
}

var OpenAIClient []ClientProcessor[*OpenAIRq, *OpenAIRs]
