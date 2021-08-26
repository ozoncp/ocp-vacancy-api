package producer

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type Producer interface {
	Send(message Message) error
}

type producer struct {
	saramaProd sarama.SyncProducer
	topic      string
}

func (p *producer) Send(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "failed to marshal message")
	}

	_, _, err = p.saramaProd.SendMessage(
		&sarama.ProducerMessage{
			Topic:     p.topic,
			Key:       sarama.StringEncoder(p.topic),
			Value:     sarama.StringEncoder(bytes),
			Partition: -1,
		},
	)
	return err
}

func NewProducer(brokers []string, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	p, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize Sarama producer")
	}
	return &producer{
		saramaProd: p,
		topic:      topic,
	}, nil
}
