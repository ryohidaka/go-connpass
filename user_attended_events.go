package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// ユーザー参加イベント一覧
//
// ユーザーが参加したイベント一覧を取得する。
//
// [APIリファレンス]
//
// パラメータ:
//   - nickname: ニックネーム (例: "haru860")
//   - query: ユーザー参加イベント検索用のクエリパラメータ（省略可能）
//
// 戻り値:
//   - ユーザー参加イベント一覧のレスポンス
//   - エラーが発生した場合は error を返す
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC/operation/connpass_account_account_api_v2_views_user_attended_event
func (c *Connpass) GetUserAttendedEvents(nickname string, query ...models.GetUserAttendedEventsQuery) (*models.GetUserAttendedEventsResponse, error) {
	var response models.GetUserAttendedEventsResponse
	endpoint := fmt.Sprintf("users/%s/attended_events", nickname)
	if err := c.Request(endpoint, query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
