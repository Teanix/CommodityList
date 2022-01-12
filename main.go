package main

import (
	"CommodityList/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	// cfg, err := tool.ParseConfig("./config/app.json")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = tool.OrmEngine(cfg)
	// if err != nil {
	// 	logger.Error(err.Error())
	// }
	// //初始化Redis
	// tool.InitRedisStore()
	//生成webapp
	app := gin.Default()
	// //设置全局跨域访问
	// app.Use(Cors())
	// //集成session功能
	// tool.InitSession(app)
	// //路由配置
	registerRouter(app)

	app.Run(":8081")
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.RoomController).Router(router)
	// new(controller.FoodCategoryController).Router(router)
	// new(controller.ShopController).Router(router)
}
