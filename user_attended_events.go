package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// ユーザーが参加したイベント一覧を取得する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC/operation/connpass_account_account_api_v2_views_user_attended_event
func (c *Connpass) GetUserAttendedEvents(nickname string, query ...models.GetUserAttendedEventsQuery) (*models.GetUserAttendedEventsResponse, error) {
	endpoint := fmt.Sprintf("users/%s/attended_events", nickname)

	res, err := Request[models.GetUserAttendedEventsResponse](c, endpoint, query)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
