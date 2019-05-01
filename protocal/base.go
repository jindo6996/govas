package protocal

import (
	"errors"
	xj "github.com/basgys/goxml2json"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"strings"
)

const DefaultByteReceive = 1024

func (gomd *Govasmd) sendXmlCommand(xmlCmd *etree.Document) (string, error) {
	cmd, _ := xmlCmd.WriteToString()
	println("\r\n")
	logrus.Info("--------In function sendXmlCommand -------")
	logrus.Info("Resquest to server: ", cmd)
	xmlC, err := gomd.sendCommand(cmd)
	if err != nil {
		logrus.Error("Error when send cmd to connection")
		return "", err
	}
	respondJson, err := xj.Convert(strings.NewReader(xmlC))
	if err != nil {
		logrus.Error("Error when convert XML to JSON: ", err)
		return "", err
	}
	end := 100
	if end > len(respondJson.String()) {
		end = len(respondJson.String())
	}
	logrus.Info("JSON respond: ", respondJson.String()[0:end])
	respondJsonString := respondJson.String()
	status := gjson.Get(respondJsonString, "*.-status").String()
	if status[0] != '2' {
		logrus.Error("Error in respond's sendXmlCommnand: ", respondJsonString)
		err := errors.New(gjson.Get(respondJsonString, "*.-status_text").String())
		return "", err
	}
	return respondJsonString, nil
}

func (gomd *Govasmd) sendCommand(cmd string) (string, error) {
	err := gomd.send(cmd)
	if err != nil {
		return "", err
	}
	respond, err := gomd.read()
	if err != nil {
		return "", err
	}
	return respond, nil
}

func (gomd *Govasmd) send(cmd string) error {
	gomd.Connect()
	_, err := gomd.Connection.Socket.Write([]byte(cmd))
	if err != nil {
		logrus.Error("Error when send request to connection (in send): ", err)
		return err
	}
	return nil
}

func (gomd *Govasmd) read() (string, error) {
	respond := ""
	for {
		buf := make([]byte, DefaultByteReceive)
		nr, err := gomd.Connection.Socket.Read(buf)
		data := buf[0:nr]
		if err != nil {
			logrus.Error("Error when read to connection (in read): ", err)
			return "", err
		}
		respond = respond + string(data)
		if nr < DefaultByteReceive {
			break
		}
	}
	end := 100
	if end > len(respond) {
		end = len(respond)
	}
	logrus.Info("Respond XML: ", respond[0:end])
	return respond, nil

}

// add filter
func addFilter(cmd *etree.Element, filter string, filter_id string) {
	if filter != "" {
		cmd.CreateAttr("filter", filter)
	}
	if filter_id != "" {
		cmd.CreateAttr("filt_id", filter_id)
	}
}

//to string from bool
func toString(condition bool) string {
	if condition {
		return "1"
	} else {
		return "0"
	}
}

func checkIntDefine(x int) bool {
	if (x == 0) || (x == 1) {
		return true
	}
	return false
}
func checkIntDefine2(x int) bool {
	if x >= 0 {
		return true
	}
	return false
}
