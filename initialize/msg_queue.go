package initialize

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"web/controller"
	"web/global"
)

func InitMsgQueue(rocketMQ string) (c rocketmq.PushConsumer, mqErr error) {
	rlog.SetLogLevel("error")

	c, mqErr = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{rocketMQ}),
		consumer.WithGroupName("purchase"),
	)

	if mqErr != nil {
		return
	}

	if err := c.Subscribe("topicName", consumer.MessageSelector{}, controller.SubscribeMsg); err != nil {
		global.Logger["err"].Infof("消费topic：purchase_order失败:%s", err.Error())
	}

	return
}
