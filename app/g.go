package main

import (
	"fmt"

	"github.com/qyqx233/chat-go-api/ent"
	"github.com/qyqx233/chat-go-api/util"
)

var client *ent.Client

func newDB(c *util.DbConfig) {
	fmt.Println(c)
	var err error
	client, err = ent.Open(c.Driver, c.Source)
	if err != nil {
		panic(err)
	}
}
