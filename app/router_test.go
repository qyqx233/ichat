package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func Test2(t *testing.T) {
	rq := &ChatSessionRq{
		Uid: "jack",
	}
	// var b = new(bytes.Buffer)
	// json.NewEncoder(b).Encode(rq)
	data, err := json.Marshal(rq)
	if err != nil {
		panic(err)
	}
	// t.Log(b.String())
	// 创建一个 HTTP 请求
	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/chat/session", bytes.NewReader(data))
	// c := &http.Client{}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Log(err)
		t.Fatal()
	}
	data, err = io.ReadAll(res.Body)
	t.Log(string(data))
}

func TestGetUser(t *testing.T) {
	// 创建一个新的 Fiber 应用程序
	app := fiber.New()
	var c = parseYaml("app.yml")
	t.Log(c)
	newDB(c.Db)
	// 注册一个路由
	app.Post("/api/chat/session", newSession)
	rq := &ChatSessionRq{
		Uid: "jack",
	}
	var b = new(bytes.Buffer)
	json.NewEncoder(b).Encode(rq)
	// t.Log(b.String())
	// 创建一个 HTTP 请求
	req := httptest.NewRequest("POST", "/api/chat/session", b)

	// 使用应用程序处理请求
	rcv, err := app.Test(req)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// 检查响应的状态码
	if rcv.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200; got %d", rcv.StatusCode)
	}
	data, err := io.ReadAll(rcv.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(string(data))

	// 检查响应的主体内容

}
