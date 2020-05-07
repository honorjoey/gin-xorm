package middleware

import (
	"github.com/Shopify/sarama"
)

type MsgReceiver struct {}

func (MsgReceiver) GetMsg(topic string) {
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		kLog.Errorf("fail to get list of partition:err%v\n", err)
		return
	}
	kLog.Info("partitionList", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			kLog.Errorf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				kLog.Infof("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
}
