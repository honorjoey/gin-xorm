package middleware

import (
	"github.com/Shopify/sarama"
)

var consumer sarama.Consumer

func StartConsumer() {
	var err error
	consumer, err = sarama.NewConsumer([]string{host+":"+port}, nil)
	if err != nil {
		kLog.Error("fail to start consumer, err:", err)
	}
}
