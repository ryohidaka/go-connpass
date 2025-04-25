package testutil

import (
	"time"

	"github.com/ryohidaka/go-connpass/models"
)

// イベント一覧取得テスト用のダミークエリを返却する
func DummyGetEventsQuery() models.GetEventsQuery {
	return models.GetEventsQuery{
		Keyword: []string{"Go"},
	}
}

// イベント一覧取得テスト用のダミーレスポンスを返却する
func DummyGetEventsResponse() *models.GetEventsResponse {
	return &models.GetEventsResponse{
		BaseResponse: models.BaseResponse{
			ResultsStart:     1,
			ResultsReturned:  1,
			ResultsAvailable: 1,
		},
		Events: []models.ConnpassEvent{
			{
				ID:          1,
				Title:       "Event 1",
				Catch:       "Catch 1",
				Description: "Description 1",
				URL:         "https://example.com/event/1",
				HashTag:     "tag1",
				StartedAt:   time.Date(2025, 4, 25, 18, 0, 0, 0, time.UTC),
				EndedAt:     time.Date(2025, 4, 25, 20, 0, 0, 0, time.UTC),
				Limit:       100,
				EventType:   models.Advertisement,
				OpenStatus:  models.Open,
				Group: models.Group{
					ID:        1,
					Subdomain: "",
					Title:     "Series 1",
					URL:       "https://example.com/series/1",
				},
				Address:          "Tokyo",
				Place:            "Somewhere",
				Lat:              "35.6895",
				Lon:              "139.6917",
				OwnerID:          123,
				OwnerNickname:    "owner1",
				OwnerDisplayName: "Owner One",
				Accepted:         50,
				Waiting:          0,
				UpdatedAt:        time.Date(2025, 4, 20, 12, 30, 0, 0, time.UTC),
			},
		},
	}
}
