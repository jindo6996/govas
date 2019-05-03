package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Restore(c *gin.Context) {
	logrus.Info("\n\n-------Restore in scanner--------")
	entityId := c.Param("entity_id")
	govas := AuthConnection(c)
	respond, err := govas.Restore(entityId)
	if err != nil {
		logrus.Error("Restore fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Restore success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
