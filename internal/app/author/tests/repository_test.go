package tests

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	author_service "gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
)

func TestAuthor_GetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id, name FROM authors WHERE id=$1", gomock.Any()).Return(nil)

		// act
		author, err := s.repo.GetById(ctx, int64(id))

		// assert
		require.NoError(t, err)
		assert.Equal(t, int64(0), author.ID)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id, name FROM authors WHERE id=$1", gomock.Any()).Return(pgx.ErrNoRows)

		// act
		author, err := s.repo.GetById(ctx, int64(id))

		// assert
		require.EqualError(t, err, author_service.ErrAuthorNotFound.Error())
		assert.Nil(t, author)
	})

	t.Run("internal error", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), "SELECT id, name FROM authors WHERE id=$1", gomock.Any()).Return(assert.AnError)

		// act
		author, err := s.repo.GetById(ctx, int64(id))

		// assert
		require.EqualError(t, err, "assert.AnError general error for testing")
		assert.Nil(t, author)
	})
}

func TestAuthor_Create(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Create(gomock.Any(), gomock.Any(), "INSERT INTO authors(id, name) VALUES($1, $2) RETURNING id", gomock.Any(), gomock.Any()).Return(nil)

		// act
		id, err := s.repo.Create(ctx, Author().Valid().P())

		// assert
		require.NoError(t, err)
		assert.Equal(t, int64(0), id)
	})

	t.Run("fail duplicate", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()
		s.mockDb.EXPECT().Create(gomock.Any(), gomock.Any(), "INSERT INTO authors(id, name) VALUES($1, $2) RETURNING id", gomock.Any(), gomock.Any()).Return(&pgconn.PgError{Code: "23505"})

		// act
		id, err := s.repo.Create(ctx, Author().Valid().P())

		// assert
		require.Error(t, err, author_service.ErrAuthorDuplicate)
		assert.Equal(t, int64(0), id)
	})
}
