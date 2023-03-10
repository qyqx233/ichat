package api

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIRq struct {
	Model   string
	Content string
}

type OpenAIRs struct {
	Content string
}

type OpenAIType ClientProcessor[*OpenAIRq, *OpenAIRs]

type OpenAI struct {
	c *openai.Client
}

func (o *OpenAI) Process(ctx context.Context, rq *OpenAIRq) (rs *OpenAIRs, err error) {
	resp, err := o.c.CreateChatCompletion(
		// context.Background(),
		ctx,
		openai.ChatCompletionRequest{
			Model: rq.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: rq.Content,
				},
			},
		},
	)
	if err != nil {
		return rs, err
	}
	rs.Content = resp.Choices[0].Message.Content
	return rs, nil
}

func newOpenAI(cfg ClientCfg) OpenAIType {
	cli := openai.NewClient(cfg.Ak)
	return &OpenAI{cli}
}

var OpenAIClient []ClientProcessor[*OpenAIRq, *OpenAIRs]
