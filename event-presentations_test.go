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

	// スロットリング対策
	time.Sleep(1 * time.Second)

	// 各資料のタイトルのみを出力
	for _, presentation := range presentations.Presentations {
		fmt.Printf("Name: %s\n", presentation.Name)
	}

	// Output:
	//
}

func TestGetEventPresentations(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// モックサーバーを作成
	t.Run("正常系", func(t *testing.T) {
		// モックレスポンスを設定
		err := testutil.MockResponseFromFile(connpass.BaseURL+"/events/364/presentations", "event-presentations")
		assert.NoError(t, err)

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")

		query := models.GetEventPresentationsQuery{
			Start: 1,
			Count: 10,
		}

		// イベント資料取得
		resp, err := c.GetEventPresentations(364, &query)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, resp.ResultsReturned, 1)
		assert.Equal(t, resp.ResultsAvailable, 91)
		assert.Equal(t, resp.ResultsStart, 1)

		p := resp.Presentations[0]
		assert.Equal(t, p.Name, "Togetterまとめ")
		assert.Equal(t, p.User.ID, 8)
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// APIエラーを模擬（400 Bad Request）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/events/0/presentations",
				httpmock.NewStringResponder(400, "Bad Request"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// イベント資料取得
			_, err := c.GetEventPresentations(0, nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// APIエラーを模擬（500 Internal Server Error）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/events/0/presentations",
				httpmock.NewStringResponder(500, "Internal Server Error"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// イベント資料取得
			_, err := c.GetEventPresentations(0, nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
