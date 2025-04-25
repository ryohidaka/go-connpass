package connpass

import "github.com/ryohidaka/go-connpass/models"

// グループ一覧
//
// 検索クエリの条件に応じたグループ一覧を取得する。
//
// [APIリファレンス](https://connpass.com/about/api/v2/#tag/%E3%82%B0%E3%83%AB%E3%83%BC%E3%83%97/operation/connpass_group_group_api_v2_views_group_search)
//
// パラメータ:
//   - query: グループ検索用のクエリパラメータ（省略可能）
//
// 戻り値:
//   - グループ一覧のレスポンス
//   - エラーが発生した場合は error を返す
func (c *Connpass) GetGroups(query *models.GetGroupsQuery) (*models.GetGroupsResponse, error) {
	var response models.GetGroupsResponse
	if err := c.Request("groups", query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
