package models

// ユーザー所属グループ一覧のパラメータ
type GetUserGroupsQuery = BaseQuery

// ユーザー所属グループ一覧のレスポンス
type GetUserGroupsResponse struct {
	BaseResponse

	Groups []ConnpassGroup `json:"groups,omitempty"` // グループ一覧
}
