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
			ResultsReturned:  1,
			ResultsAvailable: 91,
			ResultsStart:     1,
		},
		Events: []models.ConnpassEvent{
			{
				ID:          364,
				Title:       "BPStudy#56",
				Catch:       "株式会社ビープラウドが主催するWeb系技術討論の会",
				Description: "今回は「Python プロフェッショナル　プログラミング」執筆プロジェクトの継続的ビルドについて、お話しして頂きます。",
				URL:         "https://bpstudy.connpass.com/event/364/",
				HashTag:     "bpstudy",
				StartedAt:   time.Date(2012, 4, 17, 18, 30, 0, 0, time.UTC),
				EndedAt:     time.Date(2012, 4, 17, 20, 30, 0, 0, time.UTC),
				Limit:       80,
				EventType:   models.Participation,
				OpenStatus:  models.Open,
				Group: models.Group{
					ID:        1,
					Subdomain: "bpstudy",
					Title:     "BPStudy",
					URL:       "https://bpstudy.connpass.com/",
				},
				Address:          "東京都豊島区東池袋3-1-1",
				Place:            "BPオフィス (サンシャイン60 45階)",
				Lat:              "35.729402000000",
				Lon:              "139.718209000000",
				OwnerID:          8,
				OwnerNickname:    "haru860",
				OwnerDisplayName: "佐藤 治夫",
				Accepted:         80,
				Waiting:          15,
				UpdatedAt:        time.Date(2012, 3, 20, 12, 7, 32, 0, time.UTC),
			},
		},
	}
}
