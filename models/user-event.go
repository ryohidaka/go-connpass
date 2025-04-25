package models

// ユーザー参加イベント一覧のパラメータ
type GetUserAttendedEventsQuery = BaseQuery

// ユーザー参加イベント一覧のレスポンス
type GetUserAttendedEventsResponse struct {
	BaseResponse

	Events []ConnpassEvent `json:"events,omitempty"` // イベント一覧
}
