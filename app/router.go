package main

import (
	"context"
	"fmt"
	"strings"

	//  "entgo.io/ent/session"

	"github.com/gofiber/fiber/v2"
	api "github.com/qyqx233/chat-go-api/api"
	"github.com/qyqx233/chat-go-api/ent/dialog"
	sessions "github.com/qyqx233/chat-go-api/ent/session"
)

var QA_SPLIT = "\n"
var QQ_SPLIT = "\n\n"

func newSession(c *fiber.Ctx) (err error) {
	var rq = &ChatSessionRq{}
	var rs = &ChatSessionRs{}
	rs.Code = 500
	fmt.Println(string(c.Body()))
	err = c.BodyParser(rq)
	if err != nil {
		rs.Msg = err.Error()
		return c.JSON(rs)
	}
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
	_ = sessions.Table
	sess, err := client.Session.Query().Where(
		sessions.And(
			sessions.UIDEQ(rq.Uid),
			sessions.IDEQ(rq.Sid),
		)).
		First(context.Background())
	if err != nil {
		rs.Msg = err.Error()
		return c.JSON(rs)
	}
	rs.Code, rs.Sid = 200, sess.ID
	return c.JSON(rs)
}

func getHistory(sid int) (string, error) {
	l, err := client.Dialog.Query().
		Select(dialog.FieldSid).
		Where(dialog.SidEQ(sid)).
		Order().All(context.Background())
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for _, v := range l {
		sb.WriteString(v.Q)
		sb.WriteString(QA_SPLIT)
		sb.WriteString(v.A)
		sb.WriteString(QQ_SPLIT)
	}
	return sb.String(), nil
}

func packPrompt(q string, history string) {

}

func handleChat(c *fiber.Ctx) (err error) {
	var rq = new(ChatgptRq)
	var rs = new(ChatgptRs)
	var ctx context.Context
	snd := &api.OpenAIRq{}
	var rcv *api.OpenAIRs
	rs.Code = 500
	err = c.BodyParser(rq)
	if err != nil {
		rs.Msg = err.Error()
		goto exit
	}
	ctx = context.Background()
	rcv, err = api.Process(ctx, api.OpenAIClient, snd)
	if err != nil {
		rs.Msg = err.Error()
		goto exit
	}
	_, _ = client.Dialog.Create().
		SetSid(rq.Sid).
		SetQ(rq.Content).
		SetA(rcv.Content).
		Save(context.Background())
	rs.Code = 0
	rs.Content = rcv.Content
exit:
	return c.JSON(rs)
}
