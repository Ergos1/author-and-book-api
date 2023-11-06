package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	mock_database "gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql/mocks"
)

type authorRepoFixture struct {
	ctrl   *gomock.Controller
	repo   *author.AuthorRepoPsql
	mockDb *mock_database.MockDBops
}

func setUp(t *testing.T) authorRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDBops(ctrl)
	repo := author.NewAuthorRepoPsql(mockDb)

	return authorRepoFixture{
		ctrl:   ctrl,
		repo:   repo,
		mockDb: mockDb,
	}
}

func (a *authorRepoFixture) tearDown() {
	a.ctrl.Finish()
}
