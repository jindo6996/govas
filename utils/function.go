package utils

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"govas/govas"
	"govas/protocal"
	"net"
	"strings"
)

const DefaultUnixSocketPath = "/var/run/gvmd.sock"

func CreatJsonErro(mess string) string {
	return "{\"Error\":\"" + mess + "\"}"
}

func InitConnection() protocal.Govasmd {
	var soc net.Conn
	connection := &govas.UnixSocketConnection{DefaultUnixSocketPath, 60, soc}
	logrus.Info("Time to connect")
	return protocal.Govasmd{connection, false, false}
}

func AuthConnection(c *gin.Context) protocal.Govasmd {
	govas := InitConnection()
	session := sessions.Default(c)
	user := session.Get("username").(string)
	pass := session.Get("password").(string)
	_, _ = govas.Authenticate(user, pass)
	return govas
}

func String2Bool(x string) bool {
	x = strings.ToUpper(x)
	if x == "TRUE" {
		return true
	} else {
		return false
	}
}
