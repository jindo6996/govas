package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadGetInfos struct {
	TypeInfo string `form:"type_info"`
	Filter   string `form:"filter"`
	FilterId string `form:"filter_id"`
	Name     string `form:"name"`
	Details  bool   `form:"details"`
}

func GetInfos(c *gin.Context) {
	logrus.Info("\n\n-------Get all infos in scanner--------")
	var payload payloadGetInfos
	payload.Details = false
	_ = c.ShouldBindQuery(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.GetInfos(payload.TypeInfo, payload.Filter, payload.FilterId, payload.Name, payload.Details)
	if err != nil {
		logrus.Error("Get info fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get infos success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func GetInfo(c *gin.Context) {
	logrus.Info("\n\n-------Get info in scanner--------")
	infoId := c.Param("info_id")
	infoType := c.Param("info_type")
	govas := AuthConnection(c)
	respond, err := govas.GetInfo(infoType, infoId)
	if err != nil {
		logrus.Error("Get info fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get info success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
