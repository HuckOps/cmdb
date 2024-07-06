package controllers

import (
	"cmdb/pkg/restful"
	"cmdb/utils/iplist"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetMachineList(ctx *gin.Context)  {

}


func CreateMachineByIPs(ctx *gin.Context)  {
	// 创建逻辑
	// type: 1 物理机
	// 传入格式：SN:主ip:其他ip
	// type: 2 虚拟机
	// type: 3 公有云
	// 获取body文本
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		restful.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	result, err := iplist.ParseListFromString(string(body))
	if err != nil {
		restful.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(result)
}