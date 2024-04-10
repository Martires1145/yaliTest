package mq

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"time"
)

var (
	reader *kafka.Reader
	topic1 = viper.GetString("kafka.topic1")
	topic2 = viper.GetString("kafka.topic2")
)

func WriteMsg(ctx context.Context, msg []byte) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP("124.222.126.172:9092"),
		Topic:                  topic1,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		Async:                  true,
		AllowAutoTopicCreation: true,
		BatchSize:              1,
	}
	defer writer.Close()
	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(
			ctx,
			kafka.Message{Key: []byte("1145"), Value: msg},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Printf("写入失败:%v\n", err)
			}
		} else {
			break
		}
	}
}

func ReadMsg(ctx context.Context, channel *chan []byte) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"124.222.126.172:9092"},
		Topic:          topic2,
		CommitInterval: 500 * time.Millisecond,
		GroupID:        "rec_team",
		StartOffset:    kafka.FirstOffset,
		Partition:      0,
	})
	defer reader.Close()

	for {
		if message, err := reader.ReadMessage(ctx); err != nil {
			fmt.Printf("读kafka失败:%v\n", err)
			break
		} else {
			fmt.Println(string(message.Value))
			*channel <- message.Value
		}
	}
}
