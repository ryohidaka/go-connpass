package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// ユーザー発表イベント一覧
//
// ユーザーが発表したイベント一覧を取得する。
//
// [APIリファレンス](https://connpass.com/about/api/v2/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC/operation/connpass_account_account_api_v2_views_user_presenter_event)
//
// パラメータ:
//   - nickname: ニックネーム (例: "haru860")
//   - query: ユーザー発表イベント検索用のクエリパラメータ（省略可能）
//
// 戻り値:
//   - ユーザー発表イベント一覧のレスポンス
//   - エラーが発生した場合は error を返す
func (c *Connpass) GetUserPresenterEvents(nickname string, query *models.GetUserPresenterEventsQuery) (*models.GetUserPresenterEventsResponse, error) {
	var response models.GetUserPresenterEventsResponse
	endpoint := fmt.Sprintf("users/%s/presenter_events", nickname)
	if err := c.Request(endpoint, query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
