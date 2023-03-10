package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/qyqx233/chat-go-api/ent"

	_ "github.com/mattn/go-sqlite3"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.Session, error) {
	log.Println("haha")
	u, err := client.Session.Create().
		SetUID("100").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("youyo")
	// var session ent.Session
	r, _ := client.Session.Query().Select().
		Where().All(ctx)
	log.Println("user was created: ", r)
	return u, nil
}

func Test(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	t.Log(CreateUser(context.Background(), client))
}
