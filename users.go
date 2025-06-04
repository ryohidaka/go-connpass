package connpass

import "github.com/ryohidaka/go-connpass/models"

// 検索クエリの条件に応じたユーザー一覧を取得する。
//
// [APIリファレンス]
//
// [APIリファレンス]: https://connpass.com/about/api/v2/#tag/%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC/operation/connpass_account_account_api_v2_views_user_search
func (c *Connpass) GetUsers(query ...models.GetUsersQuery) (*models.GetUsersResponse, error) {
	res, err := Request[models.GetUsersResponse](c, "users", query)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
