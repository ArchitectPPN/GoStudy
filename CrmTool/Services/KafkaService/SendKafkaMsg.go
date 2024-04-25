package KafkaService

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type KafkaService struct {
	Client []string
	Topic  string
}

func (kafka *KafkaService) SetConfig(url []string, topic string) *KafkaService {
	kafka.Client = url
	kafka.Topic = topic

	return kafka
}

func (kafka *KafkaService) Send(customMsg string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = kafka.Topic
	msg.Value = sarama.StringEncoder(customMsg)
	// 连接kafka
	client, err := sarama.NewSyncProducer(kafka.Client, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	defer func(client sarama.SyncProducer) {
		err := client.Close()
		if err != nil {
			fmt.Println("关闭失败!", err)
		}
	}(client)

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
