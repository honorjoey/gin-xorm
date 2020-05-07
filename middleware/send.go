package middleware

import "github.com/Shopify/sarama"

type MsgSender struct {}

func (MsgSender)SendMsg(topic, message string) error {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)

	// 发送消息
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		kLog.Error("send msg failed, err:", err)
		return err
	}
	kLog.Infof("pid:%v offset:%v\n", pid, offset)
	return nil
}
