package producer

import (
	"errors"
	producerErrors "hokusai/internal/component/rabbitmq/errors"
	"hokusai/internal/component/rabbitmq/mocks"
	"hokusai/internal/component/rabbitmq/utils"
	mock_channel "hokusai/mocks/component/rabbitmq/channel"
	"math"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendToDirect_Success(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockIsEnabled := true
	mockExchange := "exchange_direct"
	mockRoutingKey := "routing_key_direct"
	mockDelay := 5000
	mockMessage := []byte("Test message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
		Headers: amqp.Table{
			"x-delay": int64(mockDelay),
		},
	}

	mockConfig.On("IsEnabled").Return(mockIsEnabled)
	mockConfig.On("GetExchangeDirect").Return(mockExchange)
	mockConfig.On("GetRoutingKeyDirect").Return(mockRoutingKey)
	mockConfig.On("GetDelay").Return(mockDelay)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).Return(nil)

	// Call the function being tested
	success, err := producerService.SendToDirect(mockConfig, mockMessage)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, success)
	mockChannel.AssertExpectations(t)
}

func TestSendToDirect_Failed(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockIsEnabled := true
	mockExchange := "exchange_direct"
	mockRoutingKey := "routing_key_direct"
	mockDelay := 5000
	mockMessage := []byte("Test message")
	mockError := errors.New("error sending message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
		Headers: amqp.Table{
			"x-delay": int64(mockDelay),
		},
	}

	mockConfig.On("IsEnabled").Return(mockIsEnabled)
	mockConfig.On("GetExchangeDirect").Return(mockExchange)
	mockConfig.On("GetRoutingKeyDirect").Return(mockRoutingKey)
	mockConfig.On("GetDelay").Return(mockDelay)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).
		Return(mockError)

	// Call the function being tested
	success, err := producerService.SendToDirect(mockConfig, mockMessage)

	// Assertions
	assert.Error(t, err)
	assert.False(t, success)
	assert.EqualError(t, err, "error sending message")
	mockChannel.AssertExpectations(t)
}

func TestSendToDirect_FailedOnDisabledConfig(t *testing.T) {
	var (
		mockCh          = new(mock_channel.WrappedChannelService)
		mockConfig      = new(mocks.MockRabbitMQConfig)
		producerService = NewProducerService(mockCh)

		producerErr *producerErrors.ProducerError
	)

	mockConfig.On("IsEnabled").Return(false)
	mockConfig.On("GetQueueDirect").Return("mock_queue_direct")

	isSuccess, err := producerService.SendToDirect(mockConfig, []byte(``))
	assert.False(t, isSuccess)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &producerErr)
	assert.ErrorContains(t, err, "can't publish a message to direct queue")
}

func TestSendToJunk_Success(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockIsEnabled := true
	mockExchange := "exchange_junk"
	mockRoutingKey := "routing_key_junk"
	mockMessage := []byte("Test message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
	}

	mockConfig.On("IsEnabled").Return(mockIsEnabled)
	mockConfig.On("GetExchangeJunk").Return(mockExchange)
	mockConfig.On("GetRoutingKeyJunk").Return(mockRoutingKey)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).Return(nil)

	// Call the function being tested
	success, err := producerService.SendToJunk(mockConfig, mockMessage)

	// Assertions
	assert.NoError(t, err)
	assert.True(t, success)
	mockChannel.AssertExpectations(t)
}

func TestSendToJunk_Failed(t *testing.T) {
	mockChannel := new(mock_channel.WrappedChannelService)
	mockConfig := new(mocks.MockRabbitMQConfig)
	producerService := NewProducerService(mockChannel)

	// Set up inputs
	mockIsEnabled := true
	mockExchange := "exchange_junk"
	mockRoutingKey := "routing_key_junk"
	mockMessage := []byte("Test message")
	mockError := errors.New("error sending message")

	// Set expectations on the mock
	expectedMsg := amqp.Publishing{
		ContentType:  utils.GetContentType(),
		DeliveryMode: amqp.Persistent,
		Body:         mockMessage,
	}

	mockConfig.On("IsEnabled").Return(mockIsEnabled)
	mockConfig.On("GetExchangeJunk").Return(mockExchange)
	mockConfig.On("GetRoutingKeyJunk").Return(mockRoutingKey)
	mockChannel.On("PublishMessage", mockExchange, mockRoutingKey, expectedMsg).
		Return(mockError)

	// Call the function being tested
	success, err := producerService.SendToJunk(mockConfig, mockMessage)

	// Assertions
	assert.Error(t, err)
	assert.False(t, success)
	assert.EqualError(t, err, "error sending message")
	mockChannel.AssertExpectations(t)
}

func TestSendToJunk_FailedOnDisabledConfig(t *testing.T) {
	var (
		mockCh          = new(mock_channel.WrappedChannelService)
		mockConfig      = new(mocks.MockRabbitMQConfig)
		producerService = NewProducerService(mockCh)

		producerErr *producerErrors.ProducerError
	)

	mockConfig.On("IsEnabled").Return(false)
	mockConfig.On("GetQueueDirect").Return("mock_queue_direct")

	isSuccess, err := producerService.SendToJunk(mockConfig, []byte(``))
	assert.False(t, isSuccess)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &producerErr)
	assert.ErrorContains(t, err, "can't publish a message to junk queue")
}

func TestSendToDirectJsonMarshaled_Failed(t *testing.T) {
	var (
		mockCh          = new(mock_channel.WrappedChannelService)
		mockConfig      = new(mocks.MockRabbitMQConfig)
		producerService = NewProducerService(mockCh)

		producerErr *producerErrors.ProducerError
	)

	// fail because fail to marshal json
	isSuccess, err := producerService.SendToDirectJsonMarshaled(mockConfig, math.Inf(1))
	assert.False(t, isSuccess)
	assert.Error(t, err)

	// fail because can't publish a message (disabled config)
	mockConfig.On("IsEnabled").Return(false)
	mockConfig.On("GetQueueDirect").Return("mock_queue_direct")

	isSuccess, err = producerService.SendToDirectJsonMarshaled(mockConfig, []byte(``))
	assert.False(t, isSuccess)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &producerErr)
	assert.ErrorContains(t, err, "can't publish a message to direct queue")
}

func TestSendToDirectJsonWithIncrementDelay(t *testing.T) {
	var (
		mockCh          = new(mock_channel.WrappedChannelService)
		mockConfig      = new(mocks.MockRabbitMQConfig)
		producerService = NewProducerService(mockCh)

		mockExchange   = "mock_exchange"
		mockQueue      = "mock_queue_direct"
		mockRoutingKey = "mock_routing_key"
		mockDelay      = 1

		producerErr *producerErrors.ProducerError
	)

	mockConfig.
		On("GetQueueDirect").Return(mockQueue).
		On("GetExchangeDirect").Return(mockExchange).
		On("GetRoutingKeyDirect").Return(mockRoutingKey).
		On("GetDelay").Return(mockDelay)

	// fail because the config is disabled
	mockConfig.On("IsEnabled").Return(false).Twice()
	isSuccess, err := producerService.SendToDirectJsonWithIncrementDelay(mockConfig, []byte(``), 0)
	assert.False(t, isSuccess)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &producerErr)
	assert.ErrorContains(t, err, "can't publish a delayed message to direct queue")

	// fail to marshal the payload to json
	mockConfig.On("IsEnabled").Return(true)
	isSuccess, err = producerService.SendToDirectJsonWithIncrementDelay(mockConfig, math.Inf(1), 0)
	assert.False(t, isSuccess)
	assert.Error(t, err)

	// fail to publish message to the channel
	mockCh.
		On("PublishMessage", mockExchange, mockRoutingKey, mock.Anything).
		Return(errors.New("an error")).
		Once()
	isSuccess, err = producerService.SendToDirectJsonWithIncrementDelay(mockConfig, []byte(``), 0)
	assert.False(t, isSuccess)
	assert.Error(t, err)

	// succeed
	mockCh.
		On("PublishMessage", mockExchange, mockRoutingKey, mock.Anything).
		Return(nil)
	isSuccess, err = producerService.SendToDirectJsonWithIncrementDelay(mockConfig, []byte(``), 0)
	assert.True(t, isSuccess)
	assert.NoError(t, err)
}
