package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"govas/govas"
	"govas/protocal"
	"net"
)

func main() {
	var soc net.Conn
	connection := &govas.UnixSocketConnection{"/var/run/openvasmd.sock", 120, soc}
	logrus.Info("Time to connect")
	gvmd := protocal.Govasmd{connection, false, false}
	auR, err := gvmd.Authenticate("admin", "admin123")
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("here auth", auR)
	auR, err = gvmd.Authenticate("jin", "123123")
	fmt.Println("heeeeeeeeee: ", gvmd.Authenticated)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println("here auth", auR)
	tast, _ := gvmd.GetTask()
	fmt.Println(tast)
}
