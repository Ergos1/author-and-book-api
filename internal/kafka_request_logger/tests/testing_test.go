//go:build unit
// +build unit

package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_kafka_producer "gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/kafka/mocks"
	kafkarequestlogger "gitlab.ozon.dev/ergossteam/homework-3/internal/kafka_request_logger"
)

type KafkaRequestLoggerFixture struct {
	ctrl          *gomock.Controller
	mock_producer mock_kafka_producer.MockProducerOps
	logger        kafkarequestlogger.KafkaRequestLogger
}

func setUp(t *testing.T) KafkaRequestLoggerFixture {
	ctrl := gomock.NewController(t)
	producer := mock_kafka_producer.NewMockProducerOps(ctrl)
	logger := kafkarequestlogger.NewKafkaRequestLogger(producer, nil)

	return KafkaRequestLoggerFixture{
		ctrl:          ctrl,
		mock_producer: *producer,
		logger:        *logger,
	}
}

func (a *KafkaRequestLoggerFixture) tearDown() {
	a.ctrl.Finish()
}
