package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func DictRoute(g *gin.RouterGroup) {
	roleGroup := g.Group("/dict", middlewares.JWTAuth(), middlewares.IsSuperAdminAuth())
	{
		//字典类型列表
		roleGroup.GET("/dict_type_list", controller.DictTypeList)
		//新增字典类型
		roleGroup.POST("/create_dict_type", controller.CreateDictType)
		//修改字典类型
		roleGroup.POST("/change_dict_type", controller.ChangeDictType)
		//删除字典类型
		roleGroup.POST("/delete_dict_type", controller.DeleteDictType)
		//字典数据列表
		roleGroup.GET("/dict_list", controller.DictList)
		//新增字典数据
		roleGroup.POST("/create_dict", controller.CreateDict)
		//修改字典数据
		roleGroup.POST("/change_dict", controller.ChangeDict)
		//删除字典
		roleGroup.POST("/delete_dict", controller.DeleteDict)
	}
}
