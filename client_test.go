package connpass_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/models"
)

func ExampleNewClient() {
	// APIキーを取得
	apiKey := os.Getenv("CONNPASS_API_KEY")

	// クライアントを初期化して返す
	c := connpass.NewClient(apiKey)

	// 出力
	fmt.Println(c.BaseURL)

	// Output:
	// https://connpass.com/api/v2
}

func TestNewClient(t *testing.T) {
	t.Run("Connpassクライアントが正しく初期化されること", func(t *testing.T) {
		apiKey := "test-api-key"
		c := connpass.NewClient(apiKey)

		if c == nil {
			t.Fatal("NewClient が nil を返しました")
		}

		if c.APIKey != apiKey {
			t.Errorf("APIKey = %s; want %s", c.APIKey, apiKey)
		}

		if c.Client == nil {
			t.Error("http.Client が初期化されていません")
		}

		if c.BaseURL == "" {
			t.Error("BaseURL が空です")
		}
	})
}

// モック用の http.RoundTripper を定義
type mockRoundTripper struct {
	roundTripFunc func(req *http.Request) *http.Response
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.roundTripFunc(req), nil
}

func TestRequest(t *testing.T) {
	t.Run("正常系（構造体）", func(t *testing.T) {
		// モックのレスポンスデータ
		expected := map[string]string{"status": "ok"}
		respBody, _ := json.Marshal(expected)

		mockClient := &http.Client{
			Transport: &mockRoundTripper{
				roundTripFunc: func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBuffer(respBody)),
						Header:     make(http.Header),
					}
				},
			},
		}

		c := &connpass.Connpass{
			APIKey:  "dummy",
			Client:  mockClient,
			BaseURL: "https://connpass.com/api/v2",
		}

		// モックレスポンス用の構造体
		var out map[string]string

		err := c.Request("mock", struct{}{}, &out)
		if err != nil {
			t.Fatalf("リクエストに失敗しました: %v", err)
		}
		if out["status"] != "ok" {
			t.Errorf("期待されたステータス 'ok' と異なります: %v", out["status"])
		}
	})

	t.Run("正常系（スライス）", func(t *testing.T) {
		// モックのレスポンスデータ
		expected := map[string]string{"status": "ok"}
		respBody, _ := json.Marshal(expected)

		mockClient := &http.Client{
			Transport: &mockRoundTripper{
				roundTripFunc: func(req *http.Request) *http.Response {
					if !strings.Contains(req.URL.RawQuery, "event_id") {
						t.Error("event_id パラメータがクエリに含まれていません")
					}
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBuffer(respBody)),
						Header:     make(http.Header),
					}
				},
			},
		}

		c := &connpass.Connpass{
			APIKey:  "dummy",
			Client:  mockClient,
			BaseURL: "https://connpass.com/api/v2",
		}

		// スライスとして渡すクエリ
		query := models.GetEventsQuery{
			EventID: []int{123},
			BaseQuery: models.BaseQuery{
				Start: 1,
				Count: 10,
			},
		}

		// モックレスポンス用の構造体
		var out map[string]string
		err := c.Request("mock", []models.GetEventsQuery{query}, &out)
		if err != nil {
			t.Fatalf("スライスでのリクエストに失敗しました: %v", err)
		}
		if out["status"] != "ok" {
			t.Errorf("期待されたステータス 'ok' と異なります: %v", out["status"])
		}
	})

	t.Run("異常系", func(t *testing.T) {
		// 無効な URL に対するテスト
		c := connpass.NewClient("dummy")
		err := c.Request("::://bad-url", nil, &map[string]string{})
		if err == nil {
			t.Error("不正なURLに対してエラーが返されることが期待されましたが、nilが返されました")
		}
	})

	t.Run("HTTP 400 エラー", func(t *testing.T) {
		// 400 エラーの場合
		mockClient := &http.Client{
			Transport: &mockRoundTripper{
				roundTripFunc: func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: 400,
						Body:       io.NopCloser(bytes.NewBufferString("Bad Request")),
						Header:     make(http.Header),
					}
				},
			},
		}

		c := &connpass.Connpass{
			APIKey:  "dummy",
			Client:  mockClient,
			BaseURL: "https://connpass.com/api/v2",
		}

		var out map[string]string
		err := c.Request("mock", struct{}{}, &out)
		if err == nil {
			t.Error("400 エラーに対してエラーが返されることが期待されましたが、nilが返されました")
		}
	})

	t.Run("HTTP 500 エラー", func(t *testing.T) {
		// 500 エラーの場合
		mockClient := &http.Client{
			Transport: &mockRoundTripper{
				roundTripFunc: func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: 500,
						Body:       io.NopCloser(bytes.NewBufferString("Internal Server Error")),
						Header:     make(http.Header),
					}
				},
			},
		}

		c := &connpass.Connpass{
			APIKey:  "dummy",
			Client:  mockClient,
			BaseURL: "https://connpass.com/api/v2",
		}

		var out map[string]string
		err := c.Request("mock", struct{}{}, &out)
		if err == nil {
			t.Error("500 エラーに対してエラーが返されることが期待されましたが、nilが返されました")
		}
	})

	t.Run("JSON デコードエラー", func(t *testing.T) {
		// 無効な JSON レスポンスの場合
		mockClient := &http.Client{
			Transport: &mockRoundTripper{
				roundTripFunc: func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString("Invalid JSON")),
						Header:     make(http.Header),
					}
				},
			},
		}

		c := &connpass.Connpass{
			APIKey:  "dummy",
			Client:  mockClient,
			BaseURL: "https://connpass.com/api/v2",
		}

		var out map[string]string
		err := c.Request("mock", struct{}{}, &out)
		if err == nil {
			t.Error("無効な JSON レスポンスに対してエラーが返されることが期待されましたが、nilが返されました")
		}
	})
}
