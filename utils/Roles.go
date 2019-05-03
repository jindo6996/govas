package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadAddRole struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Users   []string `json:"users"`
}
type payloadUpdateRole struct {
	RoleId  string   `json:"role_id"`
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Users   []string `json:"users"`
}

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
			logrus.Error("Get roles fail: ", err)
		} else {
			logrus.Info("Get roles success")
		}
	}
	govas.Disconnect()
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
}

// get one role with uuid
func GetRole(c *gin.Context) {
	logrus.Info("\n\n-------Get one role in scanner--------")
	uuid := c.Param("role_id")
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

func DeleteRole(c *gin.Context) {
	logrus.Info("\n\n-------Get role in scanner--------")
	roleId := c.Param("role_id")
	ultimate := String2Bool(c.Param("ultimate"))
	govas := AuthConnection(c)
	respond, err := govas.DeleteRole(roleId, ultimate)
	if err != nil {
		logrus.Error("Delete role fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Delete role success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func CloneRole(c *gin.Context) {
	logrus.Info("\n\n-------Clone role in scanner--------")
	roleId := c.Param("role_id")
	govas := AuthConnection(c)
	respond, err := govas.CloneRole(roleId)
	if err != nil {
		logrus.Error("Clone role fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Clone role success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

// create one role
func CreateRole(c *gin.Context) {
	logrus.Info("\n\n------------Add role----------")
	var payload payloadAddRole
	_ = c.ShouldBindJSON(&payload)
	govas := AuthConnection(c)
	respond, err := govas.CreateRole(payload.Name, payload.Comment, payload.Users)
	if err != nil {
		logrus.Info("Create role fail")
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create role success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
} // create one role
func UpdateRole(c *gin.Context) {
	logrus.Info("\n\n------------Add role----------")
	var payload payloadUpdateRole
	_ = c.ShouldBindJSON(&payload)
	govas := AuthConnection(c)
	respond, err := govas.UpdateRole(payload.RoleId, payload.Name, payload.Comment, payload.Users)
	if err != nil {
		logrus.Info("Create role fail")
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create role success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
