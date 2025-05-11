package connpass_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/models"
	"github.com/ryohidaka/go-connpass/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleConnpass_GetEvents() {
	// APIキーを取得
	apiKey := os.Getenv("CONNPASS_API_KEY")

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	// イベント取得パラメータを指定
	query := models.GetEventsQuery{
		EventID: []int{364},
		BaseQuery: models.BaseQuery{
			Start: 1,
			Count: 10,
		},
	}

	// イベント一覧を取得
	events, err := c.GetEvents(&query)
	if err != nil {
		fmt.Printf("イベント取得に失敗しました: %v\n", err)
		return
	}

	// スロットリング対策
	time.Sleep(1 * time.Second)

	// 各イベントのイベントIDとイベント名を出力
	for _, event := range events.Events {
		fmt.Printf("イベントID: %d, イベント名: %s\n", event.ID, event.Title)
	}

	// Output:
	// イベントID: 364, イベント名: BPStudy#56
}

func TestGetEvents(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("正常系", func(t *testing.T) {
		// モックレスポンスを設定
		err := testutil.MockResponseFromFile(connpass.BaseURL+"/events", "events")
		assert.NoError(t, err)

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")

		query := models.GetEventsQuery{
			EventID: []int{364},
			BaseQuery: models.BaseQuery{
				Start: 1,
				Count: 10,
			},
		}

		// イベント取得
		resp, err := c.GetEvents(&query)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, resp.ResultsReturned, 1)
		assert.Equal(t, resp.ResultsAvailable, 1)
		assert.Equal(t, resp.ResultsStart, 1)

		e := resp.Events[0]
		assert.Equal(t, e.ID, 364)
		assert.Equal(t, e.Title, "BPStudy#56")
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// APIエラーを模擬（400 Bad Request）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/events",
				httpmock.NewStringResponder(400, "Bad Request"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// イベント取得
			_, err := c.GetEvents(nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// APIエラーを模擬（500 Internal Server Error）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/events",
				httpmock.NewStringResponder(500, "Internal Server Error"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// イベント取得
			_, err := c.GetEvents(nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
