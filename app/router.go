package main

import (
	"context"
	"reflect"

	//  "entgo.io/ent/session"

	"github.com/gofiber/fiber/v2"
	api "github.com/qyqx233/chat-go-api/api"
	"github.com/qyqx233/chat-go-api/ent"
	"github.com/qyqx233/chat-go-api/ent/dialog"
	sessions "github.com/qyqx233/chat-go-api/ent/session"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	// "github.com/rs/zerolog/log"
)

var QA_SPLIT = "\n"
var QQ_SPLIT = "\n\n"

// var log = zerolog.New(os.Stdout).OutputJSON()

func newSession(c *fiber.Ctx) (err error) {
	var rq = &ChatSessionRq{}
	var rs = &ChatSessionRs{}
	rs.Code = 500
	if err = c.BodyParser(rq); err != nil {
		rs.Msg = err.Error()
		return c.JSON(rs)
	}
	log.Debug().Interface("rq", rq).Msg("haha")
	if rq.Sid == 0 {
		sess, err := client.Session.Create().
			SetUID(rq.Uid).
			Save(context.Background())
		if err != nil {
			rs.Code, rs.Msg = 500, err.Error()
			return c.JSON(rs)
		}
		rs.Code, rs.Sid = 0, sess.ID
		return c.JSON(rs)
	}
	sess, err := client.Session.Query().Where(
		sessions.And(
			sessions.UIDEQ(rq.Uid),
			sessions.IDEQ(rq.Sid),
		)).
		First(context.Background())
	log.Info().Interface("err", reflect.TypeOf(err)).Msg("err")
	if err != nil {
		if ent.IsNotFound(err) {
			rs.Code = 404
			return c.JSON(rs)
		}
		rs.Msg = err.Error()
		return c.JSON(rs)
	}
	rs.Code, rs.Sid = 200, sess.ID
	return c.JSON(rs)
}

func getHistory(rq *ChatgptRq) (r []openai.ChatCompletionMessage, err error) {
	var l []*ent.Dialog
	l, err = client.Dialog.Query().
		Select(dialog.FieldSid).
		Where(dialog.SidEQ(rq.Sid)).
		Order(ent.Asc("created_at")).
		Limit(yamlConfig.Param.Openai.Count).
		All(context.Background())
	if err != nil {
		return
	}
	r = append(r, api.DefaultChatCompletionMessage...)
	r = append(r, rq.H...)
	leftToken := yamlConfig.Param.Openai.Token - len(rq.Content)
	for pos := len(l) - 1; pos >= 0 && leftToken > 0; pos-- {
		d := l[pos]
		token := len(d.User) + len(d.Assistant)
		if token > leftToken {
			break
		}
		r = append(r, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: d.User,
		})
		r = append(r, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: d.Assistant,
		})
		leftToken -= token
	}
	return
}

func packPrompt(q string, history string) {

}

func handleChat(c *fiber.Ctx) (err error) {
	var rq = new(ChatgptRq)
	var rs = new(ChatgptRs)
	var ctx context.Context
	var rcv *api.OpenAIRs
	var snd *api.OpenAIRq
	if err = c.BodyParser(rq); err != nil {
		rs.Msg = err.Error()
		goto exit
	}
	snd = &api.OpenAIRq{
		Model:   rq.Model,
		Content: rq.Content,
	}
	rs.Code = 500
	ctx = context.Background()
	snd.H, err = getHistory(rq)
	if err != nil {
		rs.Msg = "failed"
		goto exit
	}
	if rcv, err = api.Process(ctx, api.OpenAIClient, snd); err != nil {
		rs.Msg = err.Error()
		goto exit
	}
	_, _ = client.Dialog.Create().
		SetSid(rq.Sid).
		SetAssistant(rq.Content).
		SetUser(rcv.Content).
		Save(context.Background())
	rs.Code = 200
	rs.Content = rcv.Content
exit:
	return c.JSON(rs)
}
