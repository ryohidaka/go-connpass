package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// イベントに投稿された資料一覧を取得する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%82%A4%E3%83%99%E3%83%B3%E3%83%88/operation/connpass_event_event_api_v2_views_event_presentation
func (c *Connpass) GetEventPresentations(id int, query ...models.GetEventPresentationsQuery) (*models.GetEventPresentationsResponse, error) {
	endpoint := fmt.Sprintf("events/%d/presentations", id)

	res, err := Request[models.GetEventPresentationsResponse](c, endpoint, query)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
