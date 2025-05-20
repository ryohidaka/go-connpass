package connpass

import "github.com/ryohidaka/go-connpass/models"

// 検索クエリの条件に応じたグループ一覧を取得する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%82%B0%E3%83%AB%E3%83%BC%E3%83%97/operation/connpass_group_group_api_v2_views_group_search
func (c *Connpass) GetGroups(query ...models.GetGroupsQuery) (*models.GetGroupsResponse, error) {
	var response models.GetGroupsResponse
	if err := c.Request("groups", query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
