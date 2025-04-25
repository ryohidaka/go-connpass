package models

import "time"

// ユーザー一覧のパラメータ
type GetUsersQuery struct {
	BaseQuery

	Nickname []string `url:"nickname,omitempty"` // ニックネーム（例: ["haru860", "ian"]）
}

// ユーザー一覧のレスポンス
type GetUsersResponse struct {
	BaseResponse

	Users []User `json:"users"` // ユーザー一覧
}

// ユーザー
type User struct {
	ID                  int       `json:"id,omitempty"`                    // ユーザーID
	Nickname            string    `json:"nickname,omitempty"`              // ニックネーム
	DisplayName         string    `json:"display_name,omitempty"`          // 表示名
	Description         string    `json:"description,omitempty"`           // 自己紹介文
	URL                 string    `json:"url,omitempty"`                   // connpass上のURL
	ImageURL            string    `json:"image_url,omitempty"`             // ユーザーのサムネイル画像URL
	CreatedAt           time.Time `json:"created_at,omitempty"`            // 利用開始日時
	AttendedEventCount  int       `json:"attended_event_count,omitempty"`  // 参加イベント数
	OrganizeEventCount  int       `json:"organize_event_count,omitempty"`  // 管理イベント数
	PresenterEventCount int       `json:"presenter_event_count,omitempty"` // 発表イベント数
	BookmarkEventCount  int       `json:"bookmark_event_count,omitempty"`  // ブックマークイベント数
}
