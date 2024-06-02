package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web/utils/xsq_net"
)

func Test(c *gin.Context) {
	// 如果需要，可以手动解析表单数据
	if err := c.Request.ParseForm(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse form"})
		return
	}

	fmt.Println(c.Keys)

	// 获取表单数据
	form := c.Request.Form

	for key, values := range form {
		fmt.Printf("Key: %s, Values: %v\n", key, values)

		// 如果字段只应该有一个值，可以使用 values[0]
		// 但要注意检查切片是否为空，以防出现索引越界错误
		if len(values) > 0 {
			for k, v := range values {

				fmt.Printf("First Value for Key %v: %v\n", k, v)
			}
		}
	}
	xsq_net.SucJson(c, form)
}
