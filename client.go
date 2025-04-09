package connpass

import (
	"fmt"
	"net/http"
	"time"
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
