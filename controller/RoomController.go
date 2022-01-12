package controller

import (
	"CommodityList/param"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomController struct { //方法
}

func (RC *RoomController) Router(engine *gin.Engine) {
	engine.POST("/roomset", RC.Roomset)
}

func (RC *RoomController) Roomset(context *gin.Context) {
	room_info := param.Room{}
	err := context.BindJSON(&room_info) //获取JSON

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//测试输入
	context.JSON(http.StatusOK, gin.H{
		"room_id":     room_info.RoomID,
		"room_passwd": room_info.RoomPasswd,
		"room_owener": room_info.RoomOwener,
	})

	
}
