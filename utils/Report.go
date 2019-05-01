package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadGetReports struct {
	Filter          string `form:"filter"`
	FilterId        string `form:"filter_id"`
	NoteDetails     int    `form:"note_details"`
	OverrideDetails int    `form:"override_details"`
}
type payloadGetReport struct {
	Filter         string `form:"filter"`
	FilterId       string `form:"filter_id"`
	DeltaReportId  string `form:"delta_report_id"`
	ReportFormatId string `form:"report_format_id"`
}

// get all reports
func GetReports(c *gin.Context) {
	logrus.Info("\n\n-------Get all reports in scanner--------")
	var payload payloadGetReports
	payload.NoteDetails = -1
	payload.OverrideDetails = -1
	_ = c.ShouldBindQuery(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.GetReports(payload.Filter, payload.FilterId, payload.NoteDetails, payload.OverrideDetails)
	if err != nil {
		logrus.Error("Get report fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get reports success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func GetReport(c *gin.Context) {
	logrus.Info("\n\n-------Get report in scanner--------")
	reportId := c.Param("report_id")
	govas := AuthConnection(c)
	var payload payloadGetReport
	_ = c.ShouldBindQuery(&payload)
	respond, err := govas.GetReport(reportId, payload.Filter, payload.FilterId, payload.DeltaReportId, payload.ReportFormatId)
	if err != nil {
		logrus.Error("Get report fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get report success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func DeleteReport(c *gin.Context) {
	logrus.Info("\n\n-------Delete report in scanner--------")
	reportId := c.Param("report_id")
	govas := AuthConnection(c)
	respond, err := govas.DeleteReport(reportId)
	if err != nil {
		logrus.Error("Delete report fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Delete report success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
