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

func ExampleConnpass_GetUsers() {
	// APIキーを取得
	apiKey := config.GetAPIKey()

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	// ユーザー取得パラメータを指定
	query := models.GetUsersQuery{
		Nickname: []string{"haru860"},
		BaseQuery: models.BaseQuery{
			Start: 1,
			Count: 10,
		},
	}

	// ユーザー一覧を取得
	users, err := c.GetUsers(&query)
	if err != nil {
		fmt.Printf("イベント取得に失敗しました: %v\n", err)
		return
	}

	// スロットリング対策
	time.Sleep(1 * time.Second)

	// 各ユーザーのユーザーIDとニックネームを出力
	for _, user := range users.Users {
		fmt.Printf("ユーザーID: %d, ニックネーム: %s\n", user.ID, user.Nickname)
	}

	// Output:
	// ユーザーID: 8, ニックネーム: haru860
}

func TestGetUsers(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// モックサーバーを作成
	t.Run("正常系", func(t *testing.T) {
		// モックレスポンスを設定
		err := testutil.MockResponseFromFile(connpass.BaseURL+"/users", "users")
		assert.NoError(t, err)

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")

		query := models.GetUsersQuery{
			Nickname: []string{"haru860"},
			BaseQuery: models.BaseQuery{
				Start: 1,
				Count: 10,
			},
		}

		// ユーザー取得
		resp, err := c.GetUsers(&query)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, resp.ResultsReturned, 1)
		assert.Equal(t, resp.ResultsAvailable, 1)
		assert.Equal(t, resp.ResultsStart, 1)

		u := resp.Users[0]
		assert.Equal(t, u.ID, 8)
		assert.Equal(t, u.Nickname, "haru860")
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// APIエラーを模擬（400 Bad Request）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/users",
				httpmock.NewStringResponder(400, "Bad Request"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// ユーザー取得
			_, err := c.GetUsers(nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// APIエラーを模擬（500 Internal Server Error）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/users",
				httpmock.NewStringResponder(500, "Internal Server Error"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// ユーザー取得
			_, err := c.GetUsers(nil)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
