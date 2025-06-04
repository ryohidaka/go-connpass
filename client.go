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
	APIKey  string
	Client  *http.Client
	BaseURL string
}

// 指定された API キーで Connpass クライアントを生成する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#section/%E6%A6%82%E8%A6%81/%E8%AA%8D%E8%A8%BC
func NewClient(apiKey string) *Connpass {
	return &Connpass{
		APIKey:  apiKey,
		Client:  &http.Client{Timeout: 10 * time.Second},
		BaseURL: BaseURL,
	}
}

// 指定された Connpass API のエンドポイントに GET リクエストを送り、
// その JSON レスポンスを指定された型 T にデコードして返却する。
func Request[T any](c *Connpass, endpoint string, queryStruct any) (T, error) {
	var zero T

	reqURL, err := url.Parse(fmt.Sprintf("%s/%s", c.BaseURL, endpoint))
	if err != nil {
		return zero, fmt.Errorf("URLのパースに失敗しました: %w", err)
	}

	// クエリ構造体が指定されている場合、URL パラメータとしてエンコード
	if queryStruct != nil {
		v := reflect.ValueOf(queryStruct)

		if v.Kind() == reflect.Slice && v.Len() > 0 {
			v = v.Index(0)
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.Kind() == reflect.Struct {
			values, err := query.Values(v.Interface())
			if err != nil {
				return zero, fmt.Errorf("クエリのエンコードに失敗しました: %w", err)
			}
			reqURL.RawQuery = values.Encode()
		}
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return zero, fmt.Errorf("HTTPリクエストの作成に失敗しました: %w", err)
	}
	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return zero, fmt.Errorf("APIリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return zero, fmt.Errorf("APIリクエストに失敗しました: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	if resp.StatusCode >= 500 {
		return zero, fmt.Errorf("予期しないエラーが発生しました: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var out T
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return zero, fmt.Errorf("JSONデコードに失敗しました: %w", err)
	}

	return out, nil
}
