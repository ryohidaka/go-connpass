package models

import "time"

// イベント一覧のパラメータ
type GetEventsQuery struct {
	BaseQuery

	EventID       []int        `url:"event_id,omitempty"`       // イベントID（複数指定可）
	Keyword       []string     `url:"keyword,omitempty"`        // キーワード(AND)（複数指定可）
	KeywordOr     []string     `url:"keyword_or,omitempty"`     // キーワード(OR)（複数指定可）
	Ym            []string     `url:"ym,omitempty"`             // イベント開催年月（例: 201204）
	Ymd           []string     `url:"ymd,omitempty"`            // イベント開催年月日（例: 20120406）
	Nickname      []string     `url:"nickname,omitempty"`       // 参加者のニックネーム
	OwnerNickname []string     `url:"owner_nickname,omitempty"` // 管理者のニックネーム
	GroupID       []int        `url:"group_id,omitempty"`       // グループID
	Subdomain     []string     `url:"subdomain,omitempty"`      // サブドメイン
	Prefecture    []Prefecture `url:"prefecture,omitempty"`     // 都道府県
	Order         EventOrder   `url:"order,omitempty"`          // 表示順（1: 更新日時順, 2: 開催日時順, 3: 新着順）
}

// イベント一覧のレスポンス
type GetEventsResponse struct {
	BaseResponse
	Events []ConnpassEvent `json:"events"` // イベント一覧
}

// イベント
type ConnpassEvent struct {
	ID               int        `json:"id"`                 // イベントID
	Title            string     `json:"title"`              // イベント名
	Catch            string     `json:"catch"`              // キャッチ
	Description      string     `json:"description"`        // 概要
	URL              string     `json:"event_url"`          // connpass.com上のURL
	ImageURL         string     `json:"image_url"`          // イベント画像URL
	HashTag          string     `json:"hash_tag"`           // X(Twitter)のハッシュタグ
	StartedAt        time.Time  `json:"started_at"`         // イベント開催日時
	EndedAt          time.Time  `json:"ended_at"`           // イベント終了日時
	Limit            int        `json:"limit"`              // 定員
	EventType        EventType  `json:"event_type"`         // イベント参加タイプ
	OpenStatus       OpenStatus `json:"open_status"`        // イベントの開催状態
	Group            Group      `json:"group"`              // グループ情報
	Address          string     `json:"address"`            // 開催場所
	Place            string     `json:"place"`              // 開催会場
	Lat              string     `json:"lat"`                // 開催会場の緯度
	Lon              string     `json:"lon"`                // 開催会場の経度
	OwnerID          int        `json:"owner_id"`           // 管理者のID
	OwnerNickname    string     `json:"owner_nickname"`     // 管理者のニックネーム
	OwnerDisplayName string     `json:"owner_display_name"` // 管理者の表示名
	Accepted         int        `json:"accepted"`           // 参加者数
	Waiting          int        `json:"waiting"`            // 補欠者数
	UpdatedAt        time.Time  `json:"updated_at"`         // 更新日時
}

// 検索結果の表示順
type EventOrder int

const (
	UpdatedAt EventOrder = 1 // 更新日時順
	StartedAt EventOrder = 2 // 開催日時順
	CreatedAt EventOrder = 3 // 新着順
)

// イベント参加タイプ
type EventType string

const (
	Participation EventType = "participation" // connpassで参加受付あり
	Advertisement EventType = "advertisement" // 告知のみ
)

// イベントの開催状態
type OpenStatus string

const (
	PreOpen   OpenStatus = "preopen"   // 開催前
	Open      OpenStatus = "open"      // 開催中
	Close     OpenStatus = "close"     // 終了
	Cancelled OpenStatus = "cancelled" // 中止
)

// グループ
type Group struct {
	ID        int    `json:"id"`        // グループID
	Subdomain string `json:"subdomain"` // サブドメイン
	Title     string `json:"title"`     // グループ名
	URL       string `json:"url"`       // グループのconnpass.com上のURL
}
