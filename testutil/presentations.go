package testutil

import (
	"time"

	"github.com/ryohidaka/go-connpass/models"
)

// イベント資料一覧取得テスト用のダミークエリを返却する
func DummyGetEventPresentationsQuery() models.GetEventPresentationsQuery {
	return models.GetEventPresentationsQuery{
		Start: 1,
		Count: 10,
	}
}

// イベント資料一覧取得テスト用のダミーレスポンスを返却する
func DummyGetEventPresentationsResponse() *models.GetEventPresentationsResponse {
	return &models.GetEventPresentationsResponse{
		BaseResponse: models.BaseResponse{
			ResultsReturned:  1,
			ResultsAvailable: 91,
			ResultsStart:     1,
		},
		Presentations: []models.Presentation{
			{
				User: &models.User{
					ID:       8,
					Nickname: "haru860",
				},
				URL:  "https://togetter.com/li/294875",
				Name: "Togetterまとめ",
				Presenter: &models.User{
					ID:       8,
					Nickname: "haru860",
				},
				PresentationType: models.PresentationTypeBlog,
				CreatedAt:        time.Date(2012, 4, 29, 19, 44, 3, 0, time.UTC),
			},
		},
	}
}
