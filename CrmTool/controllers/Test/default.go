package Test

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.TplName = "index.tpl"
}

func (c *MainController) TestConnectKafka() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "MES_CRM_QC_RESULT_V2"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"10.6.34.74:9092,10.6.34.75:9092,10.6.34.73:9092"}, config)
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

	c.Ctx.WriteString("连接Kafka")
}
