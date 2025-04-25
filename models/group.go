package models

// グループ一覧のパラメータ
type GetGroupsQuery struct {
	BaseQuery

	Subdomain []string `json:"subdomain,omitempty"` // サブドメイン（例: ["bpstudy", "beproud"]）
}

// グループ一覧のレスポンス
type GetGroupsResponse struct {
	BaseResponse

	Groups []ConnpassGroup `json:"groups"` // グループ一覧
}

// グループ
type ConnpassGroup struct {
	ID               int    `json:"id"`                 // グループID
	Subdomain        string `json:"subdomain"`          // サブドメイン
	Title            string `json:"title"`              // グループ名
	SubTitle         string `json:"sub_title"`          // サブタイトル
	URL              string `json:"url"`                // グループのconnpass.com上のURL
	Description      string `json:"description"`        // 概要
	OwnerText        string `json:"owner_text"`         // 主催者
	ImageURL         string `json:"image_url"`          // グループ画像URL
	WebsiteURL       string `json:"website_url"`        // 公式サイトURL
	WebsiteName      string `json:"website_name"`       // 公式サイト名
	TwitterUsername  string `json:"twitter_username"`   // Twitterアカウント名
	FacebookURL      string `json:"facebook_url"`       // FacebookアカウントURL
	MemberUsersCount int    `json:"member_users_count"` // グループの全メンバー数
}
