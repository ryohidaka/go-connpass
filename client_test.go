package connpass_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/internal/config"
)

func ExampleNewClient() {
	// APIキーを取得
	apiKey := config.GetAPIKey()

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

func TestRequest_Success(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		// モックのレスポンスデータ
		expected := map[string]string{"status": "ok"}
		respBody, _ := json.Marshal(expected)

		mockClient := &http.Client{
			Transport: &mockRoundTripper{
				roundTripFunc: func(req *http.Request) *http.Response {
					return &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBuffer(respBody)),
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

	t.Run("異常系", func(t *testing.T) {
		c := connpass.NewClient("dummy")
		err := c.Request("::://bad-url", nil, &map[string]string{})
		if err == nil {
			t.Error("不正なURLに対してエラーが返されることが期待されましたが、nilが返されました")
		}
	})
}
