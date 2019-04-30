package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadAddUser struct {
	Role_ids    string   `json:"role_ids"`
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	Password    string   `json:"password"`
	Host_allow  bool     `json:"hosts_allow"`
	Hosts       []string `json:"hosts"`
	Iface_allow bool     `json:"ifaces_allow"`
	Ifaces      []string `json:"ifaces"`
}
type payloadDeleteUser struct {
	User_id        string `json:"user_id"`
	Name           string `json:"name"`
	Inheritor_id   string `json:"inheritor_id"`
	Inheritor_name string `json:"inheritor_name"`
}

type payloadUpdateUser struct {
	UserID      string   `json:"user_id"`
	Name        string   `json:"name"`
	NewName     string   `json:"new_name"`
	Password    string   `json:"password"`
	Host_allow  bool     `json:"hosts_allow"`
	Hosts       []string `json:"hosts"`
	Iface_allow bool     `json:"ifaces_allow"`
	Ifaces      []string `json:"ifaces"`
	Role_ids    string   `json:"role_ids"`
}

// get all users
func GetUsers(c *gin.Context) {
	logrus.Info("\n\n-------Get all users in scanner--------")
	govas := AuthConnection(c)
	respond, err := govas.GetUsers("", "")
	if err != nil {
		logrus.Error("Get user fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get users success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

// get one user with uuid
func GetUser(c *gin.Context) {
	logrus.Info("\n\n-------Get one user in scanner--------")
	uuid := c.Param("uuid")
	logrus.Info("UUID user: ", uuid)
	govas := AuthConnection(c)
	respond, err := govas.GetUser(uuid)
	if err != nil {
		logrus.Error("Get user fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get user success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

// create one user
func CreateUser(c *gin.Context) {
	logrus.Info("\n\n------------Add user----------")
	var payload payloadAddUser
	_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	govas := AuthConnection(c)
	respond, err := govas.CreateUser(payload.Name, payload.Password, payload.Hosts, payload.Host_allow, payload.Ifaces, payload.Iface_allow, payload.Role_ids)
	if err != nil {
		logrus.Info("Create user fail")
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create user success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

// delete user
func DelUser(c *gin.Context) {
	logrus.Info("\n\n------------Delete user----------")
	var payload payloadDeleteUser
	_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	fmt.Println("hereeee", payload.User_id)
	govas := AuthConnection(c)
	respond, err := govas.DeleteUser(payload.User_id, payload.Name, payload.Inheritor_id, payload.Inheritor_name)
	if err != nil {
		logrus.Info("Delete user fail")
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Delete user success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

//update user
func UpdateUser(c *gin.Context) {
	logrus.Info("\n\n------------Delete user----------")
	var payload payloadUpdateUser
	_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	govas := AuthConnection(c)
	respond, err := govas.UpdateUser(payload.UserID, payload.Name, payload.NewName, payload.Password, payload.Hosts, payload.Host_allow, payload.Ifaces, payload.Iface_allow, payload.Role_ids)
	if err != nil {
		respond = CreatJsonErro(err.Error())
		logrus.Info("Update user fail")
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Update user success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
