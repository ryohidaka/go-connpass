package testutil

import (
	"github.com/ryohidaka/go-connpass/models"
)

// グループ一覧取得テスト用のダミークエリを返却する
func DummyGetGroupsQuery() models.GetGroupsQuery {
	return models.GetGroupsQuery{
		Subdomain: []string{"bpstudy"},
		BaseQuery: models.BaseQuery{
			Start: 1,
			Count: 10,
		},
	}
}

// グループ一覧取得テスト用のダミーレスポンスを返却する
func DummyGetGroupsResponse() *models.GetGroupsResponse {
	return &models.GetGroupsResponse{
		BaseResponse: models.BaseResponse{
			ResultsReturned:  1,
			ResultsAvailable: 91,
			ResultsStart:     1,
		},
		Groups: []models.ConnpassGroup{
			{
				ID:               1,
				Subdomain:        "bpstudy",
				Title:            "BPStudy",
				SubTitle:         "株式会社ビープラウドが主催するIT勉強会",
				URL:              "https://bpstudy.connpass.com/",
				Description:      "string",
				OwnerText:        "株式会社ビープラウド",
				ImageURL:         "string",
				WebsiteURL:       "http://www.beproud.jp/",
				WebsiteName:      "株式会社ビープラウド",
				TwitterUsername:  "bpstudy",
				FacebookURL:      "https://www.facebook.com/beproud.inc",
				MemberUsersCount: 5743,
			},
		},
	}
}
