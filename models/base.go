package models

// 共通リクエスト
type BaseQuery struct {
	Start int `url:"start,omitempty"` // 検索の開始位置（1以上、デフォルト:1）
	Count int `url:"count,omitempty"` // 取得件数（1〜100、デフォルト:10）
}

// 共通レスポンス。
type BaseResponse struct {
	ResultsReturned  int `json:"results_returned"`  // 含まれる検索結果の件数
	ResultsAvailable int `json:"results_available"` // 検索結果の総件数
	ResultsStart     int `json:"results_start"`     // 検索結果の開始位置
}
