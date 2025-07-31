package events

import (
	"XM_assignment/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
	"log/slog"
)

type eventReciever chan domain.Event

func NewEventReciever() eventReciever {
	return make(eventReciever)
}

func (er eventReciever) ProduceEvent(_ context.Context, message domain.Event) error {
	er <- message
	return nil
}

type eventProducer struct {
	msgChan  <-chan domain.Event
	stopChan chan struct{}
	topic    string
	brokers  []string
}

func NewProducer(messsageChannel <-chan domain.Event, topic string, brokers []string) *eventProducer {
	return &eventProducer{
		msgChan: messsageChannel,
		topic:   topic,
		brokers: brokers,
	}
}

func (p *eventProducer) Start() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(p.brokers, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	p.stopChan = make(chan struct{})
	slog.Info("kafka is ready to produce events")
workingCycle:
	for {
		select {
		case <-p.stopChan:
			break workingCycle
		case event := <-p.msgChan:
			eventKey, err := json.Marshal(event.Key)
			if err != nil {
				log.Err(err)
				continue
			}
			eventMessage, err := json.Marshal(event.Message)
			msg := &sarama.ProducerMessage{
				Topic: p.topic,
				Key:   sarama.StringEncoder(eventKey),
				Value: sarama.StringEncoder(eventMessage),
				Headers: []sarama.RecordHeader{
					{
						Key:   []byte("operation"),
						Value: []byte(event.Operation),
					},
				},
			}

			partition, offset, err := producer.SendMessage(msg)
			if err != nil {
				log.Err(err)
			} else {
				fmt.Printf("Sent message with key=%s, value=%s, partition=%d, offset=%d\n",
					eventKey, eventMessage, partition, offset)
			}
		}
	}
	slog.Info("kafka producer successfully shutdown")
	return nil
}
func (p *eventProducer) Stop() error {
	p.stopChan <- struct{}{}
	return nil
}
func (p *eventProducer) Info() string {
	return fmt.Sprintf("kafka eventProducer - brokers:%v. topic: %v.", p.brokers, p.topic)
}
