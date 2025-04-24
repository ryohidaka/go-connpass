package models

import "time"

// イベント資料一覧のパラメータ
type GetEventPresentationsQuery = BaseQuery

// イベント資料一覧のレスポンス
type GetEventPresentationsResponse struct {
	BaseResponse

	Presentations []Presentation `json:"presentations,omitempty"` // 資料一覧
}

// 資料
type Presentation struct {
	User             *User            `json:"user,omitempty"`              // 投稿者（資料を投稿したユーザー）
	URL              string           `json:"url,omitempty"`               // 資料URL
	Name             string           `json:"name,omitempty"`              // 資料タイトル
	Presenter        *User            `json:"presenter,omitempty"`         // 資料を発表したユーザー
	PresentationType PresentationType `json:"presentation_type,omitempty"` // 資料タイプ
	CreatedAt        time.Time        `json:"created_at,omitempty"`        // 投稿日時 (ISO-8601形式)
}

// 資料タイプ
type PresentationType string

const (
	PresentationTypeSlide PresentationType = "slide" // スライド
	PresentationTypeMovie PresentationType = "movie" // 動画
	PresentationTypeBlog  PresentationType = "blog"  // ブログなど
)

// 投稿者
type User struct {
	ID       int    `json:"id,omitempty"`       // ユーザーID
	Nickname string `json:"nickname,omitempty"` // ニックネーム
}
