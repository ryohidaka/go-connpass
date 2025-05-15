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

func ExampleConnpass_GetUserGroups() {
	// APIキーを取得
	apiKey := os.Getenv("CONNPASS_API_KEY")

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	nickname := "haru860"

	// ユーザー所属グループ取得パラメータを指定
	query := models.GetUserGroupsQuery{
		Start: 1,
		Count: 10,
	}

	// ユーザー所属グループ一覧を取得
	groups, err := c.GetUserGroups(nickname, query)
	if err != nil {
		fmt.Printf("ユーザー所属グループ取得に失敗しました: %v\n", err)
		return
	}

	// スロットリング対策
	time.Sleep(1 * time.Second)

	// 各グループのグループIDとグループ名を出力
	if len(groups.Groups) > 0 {
		fmt.Printf("グループID: %d, グループ名: %s\n", groups.Groups[0].ID, groups.Groups[0].Title)
	}

	// Output:
	//　グループID: 1, グループ名: BPStudy
}

func TestGetUserGroups(t *testing.T) {
	// モックのHTTPサーバーを有効化
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// モックサーバーを作成
	t.Run("正常系", func(t *testing.T) {
		// モックレスポンスを設定
		err := testutil.MockResponseFromFile(connpass.BaseURL+"/users/haru860/groups", "user-groups")
		assert.NoError(t, err)

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")

		query := models.GetUserGroupsQuery{
			Start: 1,
			Count: 10,
		}

		// ユーザー所属グループ取得
		resp, err := c.GetUserGroups("haru860", query)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, resp.ResultsReturned, 10)
		assert.Equal(t, resp.ResultsAvailable, 93)
		assert.Equal(t, resp.ResultsStart, 1)

		g := resp.Groups[0]
		assert.Equal(t, g.ID, 1)
		assert.Equal(t, g.Title, "BPStudy")
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// APIエラーを模擬（400 Bad Request）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/users/dummy-nickname/groups",
				httpmock.NewStringResponder(400, "Bad Request"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// ユーザー所属グループ取得
			_, err := c.GetUserGroups("dummy-nickname")

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// APIエラーを模擬（500 Internal Server Error）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/users/dummy-nickname/groups",
				httpmock.NewStringResponder(500, "Internal Server Error"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// ユーザー所属グループ取得
			_, err := c.GetUserGroups("dummy-nickname")

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
