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
	BannerID int64 `json:"bannerID"`
	SlotID   int64 `json:"slotID"`
	GroupID  int64 `json:"groupID"`
}

func (p *Producer) SendMessage(BannerID, SlotID, GroupID int64) error {
	statMessage := Message{
		BannerID: BannerID,
		SlotID:   SlotID,
		GroupID:  GroupID,
	}

	statMessageBytes, err := json.Marshal(statMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(MsgKey),
		Value:     sarama.ByteEncoder(statMessageBytes),
		Partition: -1,
		Timestamp: time.Now(),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Info("notification sent",
		zap.Int64("BannerID", BannerID),
		zap.Int64("GroupID", GroupID),
		zap.Int64("SlotID", SlotID),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)

	return nil
}
