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
	"github.com/ryohidaka/go-connpass/testutil"
	"github.com/stretchr/testify/assert"
)

func ExampleConnpass_GetUserGroups() {
	// APIキーを取得
	apiKey := config.GetAPIKey()

	// クライアントを初期化
	c := connpass.NewClient(apiKey)

	nickname := "haru860"

	// ユーザー所属グループ取得パラメータを指定
	query := models.GetUserGroupsQuery{
		Start: 1,
		Count: 10,
	}

	// ユーザー所属グループ一覧を取得
	groups, err := c.GetUserGroups(nickname, &query)
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
	// ダミーを生成
	dummyQuery := testutil.DummyGetUserGroupsQuery()
	dummyResponse := testutil.DummyGetUserGroupsResponse()

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

		// ユーザー所属グループ取得
		resp, err := c.GetUserGroups("haru860", &dummyQuery)

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

			// ユーザー所属グループ取得
			_, err := c.GetUserGroups("dummy-nickname", &dummyQuery)

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

			// ユーザー所属グループ取得
			_, err := c.GetUserGroups("dummy-nickname", &dummyQuery)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
