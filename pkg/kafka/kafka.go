package kafka

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Arkosh744/banners/internal/log"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
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
	BannerID int64  `json:"bannerId"`
	SlotID   int64  `json:"slotId"`
	GroupID  int64  `json:"groupId"`
	MsgType  string `json:"msgType"`
}

func (p *Producer) SendMessage(bannerID, slotID, groupID int64, msgType string) error {
	statMessage := Message{
		BannerID: bannerID,
		SlotID:   slotID,
		GroupID:  groupID,
		MsgType:  msgType,
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
		zap.Int64("bannerID", bannerID),
		zap.Int64("groupID", groupID),
		zap.Int64("slotID", slotID),
		zap.Int32("partition", partition),
		zap.Int64("offset", offset),
	)

	return nil
}
