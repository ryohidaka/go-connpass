package connpass_test

import (
	"testing"

	"github.com/ryohidaka/go-connpass"
)

func TestNewClient(t *testing.T) {
	t.Run("Connpassクライアントが正しく初期化されること", func(t *testing.T) {
		apiKey := "test-api-key"
		c := connpass.NewClient(apiKey)

		if c == nil {
			t.Fatal("NewClient が nil を返しました")
		}

		if c.APIKey != apiKey {
			t.Errorf("APIKey = %s; want %s", c.APIKey, apiKey)
		}

		if c.Client == nil {
			t.Error("http.Client が初期化されていません")
		}

		if c.APIVersion != "v2" {
			t.Errorf("APIVersion = %s; want v2", c.APIVersion)
		}

		if c.BaseURL == "" {
			t.Error("BaseURL が空です")
		}
	})
}
