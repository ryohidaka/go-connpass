package connpass

import "github.com/ryohidaka/go-connpass/models"

// 検索クエリの条件に応じたイベント一覧を取得する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%82%A4%E3%83%99%E3%83%B3%E3%83%88/operation/connpass_event_event_api_v2_views_event_search
func (c *Connpass) GetEvents(query ...models.GetEventsQuery) (*models.GetEventsResponse, error) {
	var response models.GetEventsResponse
	if err := c.Request("events", query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
