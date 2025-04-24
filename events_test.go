package connpass_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/internal/config"
	"github.com/ryohidaka/go-connpass/models"
	"github.com/stretchr/testify/assert"
)

func ExampleConnpass_GetEvents() {
	// APIキーを取得
	apiKey := config.GetAPIKey()

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	// イベント取得パラメータを指定
	query := models.GetEventsQuery{
		Keyword: []string{"Go言語"},
	}

	// イベント一覧を取得
	events, err := c.GetEvents(query)
	if err != nil {
		fmt.Printf("イベント取得に失敗しました: %v", err)
		return
	}

	// 出力
	eventJSON, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		fmt.Printf("JSONマーシャリングに失敗しました: %v", err)
		return
	}
	fmt.Println(string(eventJSON))
}

func TestGetEvents(t *testing.T) {
	dummyQuery := models.GetEventsQuery{
		Keyword: []string{"Go"},
	}

	dummyResponse := &models.GetEventsResponse{
		BaseResponse: models.BaseResponse{
			ResultsStart:     1,
			ResultsReturned:  1,
			ResultsAvailable: 1,
		},
		Events: []models.ConnpassEvent{
			{
				ID:          1,
				Title:       "Event 1",
				Catch:       "Catch 1",
				Description: "Description 1",
				URL:         "https://example.com/event/1",
				HashTag:     "tag1",
				StartedAt:   time.Date(2025, 4, 25, 18, 0, 0, 0, time.UTC),
				EndedAt:     time.Date(2025, 4, 25, 20, 0, 0, 0, time.UTC),
				Limit:       100,
				EventType:   models.Advertisement,
				OpenStatus:  models.Open,
				Group: models.Group{
					ID:        1,
					Subdomain: "",
					Title:     "Series 1",
					URL:       "https://example.com/series/1",
				},
				Address:          "Tokyo",
				Place:            "Somewhere",
				Lat:              "35.6895",
				Lon:              "139.6917",
				OwnerID:          123,
				OwnerNickname:    "owner1",
				OwnerDisplayName: "Owner One",
				Accepted:         50,
				Waiting:          0,
				UpdatedAt:        time.Date(2025, 4, 20, 12, 30, 0, 0, time.UTC),
			},
		},
	}

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

		// イベント取得
		resp, err := c.GetEvents(dummyQuery)

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

			// イベント取得
			_, err := c.GetEvents(dummyQuery)

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

			// イベント取得
			_, err := c.GetEvents(dummyQuery)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
