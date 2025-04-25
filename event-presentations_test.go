package connpass_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/internal/config"
	"github.com/ryohidaka/go-connpass/models"
	"github.com/ryohidaka/go-connpass/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleConnpass_GetEventPresentations() {
	// APIキーを取得
	apiKey := config.GetAPIKey()

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	eventId := 364

	// イベント取得パラメータを指定
	query := models.GetEventPresentationsQuery{
		Start: 1,
		Count: 10,
	}

	// イベント資料一覧を取得
	presentations, err := c.GetEventPresentations(eventId, &query)
	if err != nil {
		fmt.Printf("イベント資料取得に失敗しました: %v\n", err)
		return
	}

	// 各資料のタイトルのみを出力
	for _, presentation := range presentations.Presentations {
		fmt.Printf("Name: %s\n", presentation.Name)
	}

	// Output:
	// タイトル: Togetterまとめ
}

func TestGetEventPresentations(t *testing.T) {
	// ダミーを生成
	dummyQuery := testutil.DummyGetEventPresentationsQuery()
	dummyResponse := testutil.DummyGetEventPresentationsResponse()

	// モックサーバーを作成
	t.Run("正常系", func(t *testing.T) {
		// モックのHTTPサーバーを作成
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// レスポンスを模擬する
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(dummyResponse)
		}))
		defer mockServer.Close()

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")
		c.BaseURL = mockServer.URL // モックサーバーのURLを設定

		// イベント資料取得
		resp, err := c.GetEventPresentations(364, &dummyQuery)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, dummyResponse, resp)
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// モックのHTTPサーバーを作成
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// エラーレスポンスを模擬
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bad Request"))
			}))
			defer mockServer.Close()

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")
			c.BaseURL = mockServer.URL // モックサーバーのURLを設定

			// イベント資料取得
			_, err := c.GetEventPresentations(0, &dummyQuery)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// モックのHTTPサーバーを作成
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// 予期しないエラーレスポンスを模擬
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
			}))
			defer mockServer.Close()

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")
			c.BaseURL = mockServer.URL // モックサーバーのURLを設定

			// イベント資料取得
			_, err := c.GetEventPresentations(0, &dummyQuery)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
