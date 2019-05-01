package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadGetResults struct {
	Filter          string `form:"filter"`
	FilterId        string `form:"filter_id"`
	TaskId          string `form:"task_id"`
	NoteDetails     int    `form:"note_details"`
	OverrideDetails int    `form:"override_details"`
	Details         int    `form:"details"`
}

// get all results
func GetResults(c *gin.Context) {
	logrus.Info("\n\n-------Get all results in scanner--------")
	var payload payloadGetResults
	payload.NoteDetails = -1
	payload.OverrideDetails = -1
	payload.Details = -1
	_ = c.ShouldBindQuery(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.GetResults(payload.Filter, payload.FilterId, payload.TaskId, payload.NoteDetails, payload.OverrideDetails, payload.Details)
	if err != nil {
		logrus.Error("Get result fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get results success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func GetResult(c *gin.Context) {
	logrus.Info("\n\n-------Get result in scanner--------")
	resultId := c.Param("result_id")
	govas := AuthConnection(c)
	respond, err := govas.GetResult(resultId)
	if err != nil {
		logrus.Error("Get result fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get result success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
