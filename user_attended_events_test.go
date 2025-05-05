package connpass_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/internal/config"
	"github.com/ryohidaka/go-connpass/models"
	"github.com/ryohidaka/go-connpass/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleConnpass_GetUserAttendedEvents() {
	// APIキーを取得
	apiKey := config.GetAPIKey()

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	nickname := "haru860"

	// ユーザー参加イベント取得パラメータを指定
	query := models.GetUserAttendedEventsQuery{
		Start: 1,
		Count: 10,
	}

	// ユーザー参加イベント一覧を取得
	events, err := c.GetUserAttendedEvents(nickname, &query)
	if err != nil {
		fmt.Printf("ユーザー参加イベント取得に失敗しました: %v\n", err)
		return
	}

	// スロットリング対策
	time.Sleep(1 * time.Second)

	// 各イベントのイベントIDとイベント名を出力
	if len(events.Events) > 0 {
		fmt.Printf("イベントID: %d, イベント名: %s\n", events.Events[0].ID, events.Events[0].Title)
	}

	// Output:
	//　イベントID: 5, イベント名: 世界のやまちゃんで手羽先を喰らう会
}

func TestGetUserAttendedEvents(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// モックサーバーを作成
	t.Run("正常系", func(t *testing.T) {
		// モックレスポンスを設定
		err := testutil.MockResponseFromFile(connpass.BaseURL+"/users/haru860/attended_events", "user-events")
		assert.NoError(t, err)

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")

		query := models.GetUserAttendedEventsQuery{
			Start: 1,
			Count: 10,
		}

		// ユーザー参加イベント取得
		resp, err := c.GetUserAttendedEvents("haru860", &query)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, resp.ResultsReturned, 10)
		assert.Equal(t, resp.ResultsAvailable, 60)
		assert.Equal(t, resp.ResultsStart, 1)

		e := resp.Events[0]
		assert.Equal(t, e.ID, 353126)
		assert.Equal(t, e.Title, "BPStudy#213〜ビジネスアナリシスとDDD（ドメイン駆動設計）")
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// APIエラーを模擬（400 Bad Request）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/users/dummy-nickname/attended_events",
				httpmock.NewStringResponder(400, "Bad Request"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// ユーザー参加イベント取得
			_, err := c.GetUserAttendedEvents("dummy-nickname", nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// APIエラーを模擬（500 Internal Server Error）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/users/dummy-nickname/attended_events",
				httpmock.NewStringResponder(500, "Internal Server Error"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// ユーザー参加イベント取得
			_, err := c.GetUserAttendedEvents("dummy-nickname", nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
