package govas

import (
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"time"
)

type UnixSocketConnection struct {
	Path    string
	TimeOut int
	Socket  net.Conn
}

func (unixSoc *UnixSocketConnection) Connect() {
	unixSo, err := net.Dial("unix", unixSoc.Path)
	if err != nil {
		logrus.Error(err)
		os.Exit(100)
	}
	unixSoc.Socket = unixSo
	errTimeOut := unixSoc.Socket.SetReadDeadline(time.Now().Add(time.Duration(unixSoc.TimeOut) * time.Second))
	if errTimeOut != nil {
		logrus.Error("Set timeout error", errTimeOut)
	}
}

// disconnect unixsocket
func (unixSoc *UnixSocketConnection) Disconnect() error {
	if unixSoc.Socket != nil {
		err := unixSoc.Socket.Close()
		if err != nil {
			logrus.Error("Close connection error", err)
			return err
		}
	}
	return nil
}
