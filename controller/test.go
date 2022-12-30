package controller

import (
	"context"
	"errors"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gin-gonic/gin"
	"web/utils/xsq_net"
)

func FormParams(c *gin.Context) {

	xsq_net.SucJson(c, c.ClientIP())
}

func SubscribeMsg(ctx context.Context, messages ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	return consumer.ConsumeRetryLater, errors.New("异常")
}
