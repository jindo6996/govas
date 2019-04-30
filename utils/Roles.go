package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//get all roles
func GetRoles(c *gin.Context) {
	logrus.Info("\n\n------------Get Roles----------")
	govas := AuthConnection(c)
	respond, err := govas.GetRoles("", "")
	if err != nil {
		respond = CreatJsonErro("Error in server")
	} else {
		//_, err = fmt.Fprint(w, respond)
		if err != nil {
			logrus.Error("Get users fail: ", err)
		} else {
			logrus.Info("Get users success")
		}
	}
	govas.Disconnect()
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
}

// get one role with uuid
func GetRole(c *gin.Context) {
	logrus.Info("\n\n-------Get one role in scanner--------")
	uuid := c.Param("uuid")
	logrus.Info("Role's UUID: ", uuid)
	govas := AuthConnection(c)
	respond, err := govas.GetRole(uuid)
	if err != nil {
		respond = CreatJsonErro("Error in server")
	} else {
		//_, err = fmt.Fprint(w, respond)
		if err != nil {
			logrus.Error("Get role fail: ", err)
		} else {
			logrus.Info("Get role success")
		}
	}
	govas.Disconnect()
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
}
