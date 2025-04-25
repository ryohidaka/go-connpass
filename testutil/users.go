package testutil

import (
	"time"

	"github.com/ryohidaka/go-connpass/models"
)

// ユーザー一覧取得テスト用のダミークエリを返却する
func DummyGetUsersQuery() models.GetUsersQuery {
	return models.GetUsersQuery{
		Nickname: []string{"haru860"},
		BaseQuery: models.BaseQuery{
			Start: 1,
			Count: 10,
		},
	}
}

// ユーザー一覧取得テスト用のダミーレスポンスを返却する
func DummyGetUsersResponse() *models.GetUsersResponse {
	return &models.GetUsersResponse{
		BaseResponse: models.BaseResponse{
			ResultsReturned:  1,
			ResultsAvailable: 91,
			ResultsStart:     1,
		},
		Users: []models.User{
			{
				ID:                  8,
				Nickname:            "haru860",
				DisplayName:         "佐藤 治夫",
				Description:         "株式会社ビープラウド代表取締役。connpass企画・開発・運営。\nhttp://twitter.com/haru860\nhttp://shacho.beproud.jp/",
				URL:                 "https://connpass.com/user/haru860/",
				ImageURL:            "string",
				CreatedAt:           time.Date(2011, 10, 20, 18, 23, 3, 0, time.UTC),
				AttendedEventCount:  261,
				OrganizeEventCount:  231,
				PresenterEventCount: 34,
				BookmarkEventCount:  57,
			},
		},
	}
}
