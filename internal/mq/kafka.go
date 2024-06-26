package mq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"time"
)

var (
	reader  *kafka.Reader
	topic1  = viper.GetString("kafka.topic1")
	topic2  = viper.GetString("kafka.topic2")
	host    = viper.GetString("kafka.host")
	TimeOut = viper.GetString("kafka.TimeOut")
)

func WriteTopicID(ctx context.Context, topicR, topicW string) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(host),
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
			kafka.Message{Key: []byte("1145"), Value: []byte(topicW)},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Printf("写入失败:%v\n", err)
			}
		} else {
			fmt.Printf("写入完成")
		}
	}

	if err := writer.WriteMessages(
		ctx,
		kafka.Message{Key: []byte("1145"), Value: []byte(topicR)},
	); err != nil {
		fmt.Printf("写入失败:%v\n", err)
	} else {
		fmt.Printf("写入完成")
	}
}

func ReadTopicAck(ctx context.Context, topic string) bool {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:          []string{host},
		Topic:            topic,
		CommitInterval:   500 * time.Millisecond,
		GroupID:          "rec_team",
		StartOffset:      kafka.LastOffset,
		ReadBatchTimeout: 10 * time.Millisecond,
		Partition:        0,
	})
	defer reader.Close()

	if message, err := reader.ReadMessage(ctx); err != nil {
		fmt.Printf("读kafka失败:%v\n", err)
		return false
	} else {
		if string(message.Value) == "ACK" {
			return true
		}
	}
	return false
}

func WriteTopicAck(ctx context.Context, topic string) {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(host),
		Topic:                  topic,
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
			kafka.Message{Key: []byte("1145"), Value: []byte("ACK")},
		); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				fmt.Printf("写入失败:%v\n", err)
			}
		} else {
			fmt.Printf("写入完成")
		}
	}
}

func WriteParams(ctx context.Context, params []byte) error {
	return WriteMsg(ctx, topic2, params)
}

func ReadParamsResult(ctx context.Context) (bool, error) {
	dataChan := make(chan []byte)
	ReadMsg(ctx, topic2, &dataChan)

	data := <-dataChan
	result := make(map[string]any)

	err := json.Unmarshal(data, &result)
	if err != nil {
		return false, err
	}

	if useKafka, ok := result["is_stream"]; ok {
		return useKafka.(bool), nil
	}

	return false, errors.New(result["err"].(string))
}

func WriteMsg(ctx context.Context, topic string, msg []byte) error {
	writer := &kafka.Writer{
		Addr:                   kafka.TCP(host),
		Topic:                  topic,
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
				return err
			}
		} else {
			fmt.Printf("写入完成")
		}
	}
	return nil
}

func ReadMsg(ctx context.Context, topic string, channel *chan []byte) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:          []string{host},
		Topic:            topic,
		CommitInterval:   500 * time.Millisecond,
		GroupID:          "rec_team",
		StartOffset:      kafka.LastOffset,
		ReadBatchTimeout: 10 * time.Millisecond,
		Partition:        0,
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
