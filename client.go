package connpass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

// Connpass クライアント構造体
type Connpass struct {
	APIKey     string
	Client     *http.Client
	APIVersion string
	BaseURL    string
}

// 指定された API キーで Connpass クライアントを生成する。
//
// https://connpass.com/about/api/v2/#section/%E6%A6%82%E8%A6%81/%E8%AA%8D%E8%A8%BC
//
// 例:
//
//	c := NewClient("YOUR_API_KEY")
func NewClient(apiKey string) *Connpass {
	return &Connpass{
		APIKey:     apiKey,
		Client:     &http.Client{Timeout: 10 * time.Second},
		APIVersion: APIVersion,
		BaseURL:    fmt.Sprintf("%s/%s", BaseURL, APIVersion),
	}
}

// 指定された Connpass API のエンドポイントに GET リクエストを送り、
// その JSON レスポンスを指定された型 T にデコードして返却する。
//
// パラメータ:
//   - endpoint: API の相対パス (例: "events")
//   - query: クエリパラメータのマップ (キーに対して複数の値を指定可能)
//
// 戻り値:
//   - レスポンスをデコードした型 T の値
//   - エラーが発生した場合は error を返す
func (c *Connpass) Request(endpoint string, queryStruct any, out any) error {
	reqURL, err := url.Parse(fmt.Sprintf("%s/%s", c.BaseURL, endpoint))
	if err != nil {
		return fmt.Errorf("URLのパースに失敗しました: %w", err)
	}

	// クエリを構造体で受け取った場合のみエンコード
	if queryStruct != nil && reflect.TypeOf(queryStruct).Kind() == reflect.Struct {
		values, err := query.Values(queryStruct)
		if err != nil {
			return fmt.Errorf("クエリのエンコードに失敗しました: %w", err)
		}
		reqURL.RawQuery = values.Encode()
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("HTTPリクエストの作成に失敗しました: %w", err)
	}
	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("APIリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("APIリクエストに失敗しました: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	if resp.StatusCode >= 500 {
		return fmt.Errorf("予期しないエラーが発生しました: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("JSONデコードに失敗しました: %w", err)
	}

	return nil
}
