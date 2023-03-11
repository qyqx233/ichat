package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"unsafe"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qyqx233/chat-go-api/util"
	"github.com/sashabaranov/go-openai"
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

func wrap[S, R any](t *testing.T, prepare func(*S) (url string, app *fiber.App), checkRes func(*R)) {
	var c = yamlConfig
	newDB(c.Db)
	initApi(c)
	var s = new(S)
	var r = new(R)
	var b bytes.Buffer
	url, app := prepare(s)
	json.NewEncoder(&b).Encode(s)
	req := httptest.NewRequest("POST", url, &b)
	req.Header = map[string][]string{
		"Content-type": {"application/json"},
	}
	rcv, err := app.Test(req)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if rcv.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200; got %d", rcv.StatusCode)
		t.Fail()
	}
	data, err := io.ReadAll(rcv.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	json.Unmarshal(data, r)
	checkRes(r)
}

func TestNewSession(t *testing.T) {
	wrap(t, func(r *ChatSessionRq) (string, *fiber.App) {
		r.Sid = 100
		r.Uid = "jack"
		return "/api/chat/session", newApp()
	}, func(r *ChatSessionRs) {
		t.Log(r)
	})
}

func TestQuery(t *testing.T) {
	newDB(yamlConfig.Db)
	defer client.Close()
	r, err := client.Session.Query().All(context.Background())
	if err != nil {
		t.Log(err)
		t.Fatal()
	}
	t.Log(r)
}

func TestChat(t *testing.T) {
	wrap(t, func(s *ChatgptRq) (string, *fiber.App) {
		s.Sid = 9
		s.Content = "hello"
		s.Model = openai.GPT3Dot5Turbo
		return "/api/chat/qa", newApp()
	}, func(r *ChatgptRs) {
		t.Log(r)
	})
}

func TestDb(t *testing.T) {
	newDB(yamlConfig.Db)
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func init() {
	util.ParseYaml("app.yaml")
	yamlConfig = util.YamlConfig
}

func TestHacker(t *testing.T) {
	cli := openai.NewClient("")
	// t.Log((reflect.ValueOf(cli).Elem().FieldByName("config").Interface()))
	proxyURL, _ := url.Parse(yamlConfig.Param.Proxy.URL)
	(*openai.ClientConfig)(unsafe.Pointer(&cli)).HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
}
