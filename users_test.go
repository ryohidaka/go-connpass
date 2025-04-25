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
	// ダミーを生成
	dummyQuery := testutil.DummyGetUsersQuery()
	dummyResponse := testutil.DummyGetUsersResponse()

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

		// ユーザー取得
		resp, err := c.GetUsers(&dummyQuery)

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

			// ユーザー取得
			_, err := c.GetUsers(&dummyQuery)

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

			// ユーザー取得
			_, err := c.GetUsers(&dummyQuery)

			// エラーチェック
			assert.Error(t, err)
			assert.Equal(t, "予期しないエラーが発生しました: 500 Internal Server Error", err.Error())
		})
	})
}
