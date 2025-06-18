package mq

import (
	"context"
	"math"
	"time"

	"github.com/IBM/sarama"
	"github.com/bytedance/sonic"
	"github.com/muammarahlnn/learnyscape-backend/pkg/mq"
	"github.com/muammarahlnn/user-service/internal/constant"
	"github.com/muammarahlnn/user-service/internal/log"
)

type UserCreatedProducer struct {
	producer  sarama.AsyncProducer
	topic     string
	retryChan chan *sarama.ProducerMessage
}

func NewUserCreatedProducer(producer sarama.AsyncProducer) mq.KafkaProducer {
	pr := &UserCreatedProducer{
		producer:  producer,
		topic:     constant.UserCreatedTopic,
		retryChan: make(chan *sarama.ProducerMessage, 100),
	}

	go pr.HandleErrors()
	go pr.Retry()

	return pr
}

func (p *UserCreatedProducer) Send(ctx context.Context, event mq.KafkaEvent) error {
	bytes, err := sonic.Marshal(event)
	if err != nil {
		return err
	}

	select {
	case p.producer.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.StringEncoder(event.ID()),
		Value: sarama.ByteEncoder(bytes),
	}:
		log.Logger.Infof("topic %v sent a message %v", p.Topic(), string(bytes))
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (p *UserCreatedProducer) HandleErrors() {
	for err := range p.producer.Errors() {
		if err != nil {
			log.Logger.Errorf("failed to send message: %s", err)
		}

		p.retryChan <- err.Msg
	}
}

func (p *UserCreatedProducer) Retry() {
	for msg := range p.retryChan {
		metadata := &mq.KafkaMetadata{}
		p.unmarshalMetadata(msg, metadata)

		if metadata.Retry > constant.KafkaProducerRetryLimit {
			log.Logger.Errorf(
				"failed to send mesasge after %d retries [partition-%v]-[offet-%v] %v",
				constant.KafkaProducerRetryLimit,
				msg.Partition,
				msg.Offset,
				msg.Value,
			)
			return
		}

		metadata.Retry++
		metaBytes, _ := sonic.Marshal(metadata)
		msg.Metadata = metaBytes

		backoff := time.Duration(math.Pow(constant.KafkaProducerRetryDelay, float64(metadata.Retry))) * time.Second
		time.Sleep(backoff)

		select {
		case p.producer.Input() <- msg:
			log.Logger.Infof("retrying message (attempt %d) to topic %s", metadata.Retry, msg.Topic)
		default:
			log.Logger.Infof("failed to retry message (attempt %d) to topic %s", metadata.Retry, msg.Topic)
		}
	}
}

func (p *UserCreatedProducer) unmarshalMetadata(msg *sarama.ProducerMessage, metadata *mq.KafkaMetadata) {
	if msg.Metadata == nil {
		return
	}

	metaBytes, ok := msg.Metadata.([]byte)
	if !ok {
		return
	}

	sonic.Unmarshal(metaBytes, metadata)
}

func (p *UserCreatedProducer) Topic() string {
	return p.topic
}
