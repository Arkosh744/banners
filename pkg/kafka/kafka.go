package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Arkosh744/banners/internal/log"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"time"
)

type Producer struct {
	producer sarama.SyncProducer
	topic    string
}

type Handler func(id string)

func NewProducer(producer sarama.SyncProducer, topic string) Producer {
	return Producer{
		producer: producer,
		topic:    topic,
	}
}

const MsgKey = "stats"

type Message struct {
	BannerID  int       `json:"bannerId"`
	SlotID    int       `json:"slotId"`
	GroupID   int       `json:"groupId"`
	Type      string    `json:"-"`
	Timestamp time.Time `json:"-"`
}

func (p *Producer) SendMessage(message Message) error {
	messageData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(MsgKey),
		Value:     sarama.ByteEncoder(messageData),
		Partition: -1,
		Timestamp: time.Now(),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Info("notification sent",
		zap.Int("BannerID", message.BannerID),
		zap.Int("GroupID", message.GroupID),
		zap.Int("SlotID", message.SlotID),
		zap.String("Type", message.Type),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)

	return nil
}
