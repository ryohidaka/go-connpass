package connpass_test

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/ryohidaka/go-connpass"
	"github.com/ryohidaka/go-connpass/models"
	"github.com/stretchr/testify/assert"
)

//go:embed __fixtures__/groups.json
var groupsJSON []byte

func ExampleConnpass_GetGroups() {
	// APIキーを取得
	apiKey := os.Getenv("CONNPASS_API_KEY")

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	// グループ取得パラメータを指定
	query := models.GetGroupsQuery{
		Subdomain: []string{"bpstudy"},
		BaseQuery: models.BaseQuery{
			Start: 1,
			Count: 10,
		},
	}

	// グループ一覧を取得
	groups, err := c.GetGroups(query)
	if err != nil {
		fmt.Printf("グループ取得に失敗しました: %v\n", err)
		return
	}

	// スロットリング対策
	time.Sleep(1 * time.Second)

	// 各グループのグループIDとグループ名を出力
	for _, group := range groups.Groups {
		fmt.Printf("グループID: %d, グループ名: %s\n", group.ID, group.Title)
	}

	// Output:
	//
}

func TestGetGroups(t *testing.T) {
	// モックサーバーを作成
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("正常系", func(t *testing.T) {
		// モックレスポンスを設定
		httpmock.RegisterResponder("GET", connpass.BaseURL+"/groups",
			httpmock.NewStringResponder(200, string(groupsJSON)))

		// クライアント設定
		c := connpass.NewClient("dummy-api-key")

		query := models.GetGroupsQuery{
			Subdomain: []string{"bpstudy"},
			BaseQuery: models.BaseQuery{
				Start: 1,
				Count: 10,
			},
		}

		// グループ取得
		resp, err := c.GetGroups(query)

		// レスポンスの確認
		assert.NoError(t, err)
		assert.Equal(t, resp.ResultsReturned, 1)
		assert.Equal(t, resp.ResultsAvailable, 1)
		assert.Equal(t, resp.ResultsStart, 1)

		g := resp.Groups[0]
		assert.Equal(t, g.ID, 1)
		assert.Equal(t, g.Title, "BPStudy")
	})

	// 異常系テストケース
	t.Run("異常系", func(t *testing.T) {
		t.Run("APIエラー", func(t *testing.T) {
			// APIエラーを模擬（400 Bad Request）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/groups",
				httpmock.NewStringResponder(400, "Bad Request"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// グループ取得
			_, err := c.GetGroups()

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "APIリクエストに失敗しました: 400 Bad Request", err.Error())
		})

		t.Run("予期しないエラーAPIエラー", func(t *testing.T) {
			// APIエラーを模擬（500 Internal Server Error）
			httpmock.RegisterResponder("GET", connpass.BaseURL+"/groups",
				httpmock.NewStringResponder(500, "Internal Server Error"))

			// クライアント設定
			c := connpass.NewClient("dummy-api-key")

			// グループ取得
			_, err := c.GetGroups()

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
