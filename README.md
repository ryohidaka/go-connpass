# go-connpass

[![Go Reference](https://pkg.go.dev/badge/github.com/ryohidaka/go-connpass.svg)](https://pkg.go.dev/github.com/ryohidaka/go-connpass)
![GitHub Release](https://img.shields.io/github/v/release/ryohidaka/go-connpass)
[![codecov](https://codecov.io/gh/ryohidaka/go-connpass/graph/badge.svg?token=BSFAcheBNm)](https://codecov.io/gh/ryohidaka/go-connpass)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryohidaka/go-connpass)](https://goreportcard.com/report/github.com/ryohidaka/go-connpass)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go 用 connpass API v2 クライアント

## インストール

```bash
go get github.com/ryohidaka/go-connpass
```

## 使用例

> [!IMPORTANT]
> すべての API エンドポイントでは、API キーによる認証が必須です。
>
> API キーの発行には[ヘルプページ](https://help.connpass.com/api/)での利用申請が必要です。

```go
import "github.com/ryohidaka/go-connpass"


func main() {
    // APIキーを取得
    apiKey := "<YOUR_API_KEY>"

    // クライアントを初期化
    c := connpass.NewClient(apiKey)

    // イベント一覧を取得
    events, err := c.GetEvents()

    // イベント資料一覧を取得
    presentations, err := c.GetEventPresentations(364)

    // グループ一覧を取得
    groups, err := c.GetGroups()

    // ユーザー一覧を取得
    users, err := c.GetUsers()

    // ユーザー所属グループ一覧を取得
    userGroups, err := c.GetUserGroups("haru860")

    // ユーザー参加イベント一覧を取得
    userAttendedEvents, err := c.GetUserAttendedEvents("haru860")

    // ユーザー発表イベント一覧を取得
    userPresenterEvents, err := c.GetUserPresenterEvents("haru860")
}
```

## リンク

- [API リファレンス](https://connpass.com/about/api/v2/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
