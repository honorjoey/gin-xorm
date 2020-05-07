package middleware

import  (
	"github.com/Shopify/sarama"
)

var producer sarama.SyncProducer

func StartProducer() {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	var err error
	producer, err = sarama.NewSyncProducer([]string{host+":"+port}, config)
	if err != nil {
		kLog.Error("producer closed, err:", err)
	}
}
