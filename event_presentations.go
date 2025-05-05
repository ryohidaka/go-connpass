package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// イベントに投稿された資料一覧
//
// イベントに投稿された資料一覧を取得する。
//
// [APIリファレンス](https://connpass.com/about/api/v2/#tag/%E3%82%A4%E3%83%99%E3%83%B3%E3%83%88/operation/connpass_event_event_api_v2_views_event_presentation)
//
// パラメータ:
//   - id: イベントID (例: "364")
//   - query: 資料検索用のクエリパラメータ（省略可能）
//
// 戻り値:
//   - 資料一覧のレスポンス
//   - エラーが発生した場合は error を返す
func (c *Connpass) GetEventPresentations(id int, query *models.GetEventPresentationsQuery) (*models.GetEventPresentationsResponse, error) {
	var response models.GetEventPresentationsResponse
	endpoint := fmt.Sprintf("events/%d/presentations", id)
	if err := c.Request(endpoint, query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
