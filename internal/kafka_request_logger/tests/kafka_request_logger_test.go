//go:build unit
// +build unit

package tests

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Log(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		s := setUp(t)
		defer s.tearDown()
		s.mock_producer.EXPECT().SendSyncMessage(gomock.Any()).Return(int32(0), int64(0), nil)

		err := s.logger.Log("GET", "example/1", nil)

		require.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		s := setUp(t)
		defer s.tearDown()
		s.mock_producer.EXPECT().SendSyncMessage(gomock.Any()).Return(int32(0), int64(0), errors.New("some error from kafka"))

		err := s.logger.Log("GET", "example/1", nil)

		require.Error(t, err, errors.New("some error from kafka"))
	})
}
