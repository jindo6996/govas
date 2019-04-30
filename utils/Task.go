package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadGetTasks struct {
	Filter        string `json:"filter"`
	FilterId      string `json:"filter_id"`
	Trash         bool   `json:"trash"`
	Details       bool   `json:"details"`
	SchedulesOnly bool   `json:"schedules_only"`
}

// get all tasks
func GetTasks(c *gin.Context) {
	logrus.Info("\n\n-------Get all tasks in scanner--------")
	var payload payloadGetTasks
	_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	govas := AuthConnection(c)
	respond, err := govas.GetTasks(payload.Filter, payload.FilterId, payload.Trash, payload.Details, payload.SchedulesOnly)
	if err != nil {
		logrus.Error("Get task fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get tasks success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
