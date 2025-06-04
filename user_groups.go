package connpass

import (
	"fmt"

	"github.com/ryohidaka/go-connpass/models"
)

// ユーザーが所属しているグループ一覧を取得する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC/operation/connpass_account_account_api_v2_views_user_group
func (c *Connpass) GetUserGroups(nickname string, query ...models.GetUserGroupsQuery) (*models.GetUserGroupsResponse, error) {
	endpoint := fmt.Sprintf("users/%s/groups", nickname)

	res, err := Request[models.GetUserGroupsResponse](c, endpoint, query)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
