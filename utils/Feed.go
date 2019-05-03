package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func SynsNVT(c *gin.Context) {
	logrus.Info("\n\n-------Get Syns NTV in scanner--------")
	govas := AuthConnection(c)
	respond, err := govas.SynNVT()
	if err != nil {
		logrus.Error("Sync NTV fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/text; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Syns NVT success")
		c.Data(http.StatusOK, "application/text; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
func SynsCert(c *gin.Context) {
	logrus.Info("\n\n-------Get Syns cert in scanner--------")
	govas := AuthConnection(c)
	respond, err := govas.SynCert()
	if err != nil {
		logrus.Error("Sync cert fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/text; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Sync cert success")
		c.JSON(http.StatusOK, gin.H{"detail": respond})

	}
	_ = govas.Disconnect()
}

func SynsSCAP(c *gin.Context) {
	logrus.Info("\n\n-------Get Sync SCAP in scanner--------")
	govas := AuthConnection(c)
	respond, err := govas.SynSCAP()
	if err != nil {
		logrus.Error("Sync SCAP fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/text; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Sync SCAP success")
		c.JSON(http.StatusOK, gin.H{"detail": respond})
	}
	_ = govas.Disconnect()
}
func GetFeeds(c *gin.Context) {
	logrus.Info("\n\n-------Get feed in scanner--------")
	govas := AuthConnection(c)
	respond, err := govas.GetFeeds()
	if err != nil {
		logrus.Error("Get feeds fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/text; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get feed success")
		c.Data(http.StatusOK, "application/text; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
