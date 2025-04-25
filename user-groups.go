package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// ユーザー所属グループ一覧
//
// ユーザーが所属しているグループ一覧を取得する。
//
// [APIリファレンス](https://connpass.com/about/api/v2/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC/operation/connpass_account_account_api_v2_views_user_group)
//
// パラメータ:
//   - nickname: ニックネーム (例: "haru860")
//   - query: ユーザー所属グループ検索用のクエリパラメータ（省略可能）
//
// 戻り値:
//   - ユーザー所属グループ一覧のレスポンス
//   - エラーが発生した場合は error を返す
func (c *Connpass) GetUserGroups(nickname string, query *models.GetUserGroupsQuery) (*models.GetUserGroupsResponse, error) {
	var response models.GetUserGroupsResponse
	endpoint := fmt.Sprintf("users/%s/groups", nickname)
	if err := c.Request(endpoint, query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
