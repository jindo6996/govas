package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadGetSchedules struct {
	Filter   string `form:"filter"`
	FilterId string `form:"filter_id"`
	Trash    bool   `form:"trash"`
	Tasks    bool   `form:"tasks"`
}
type payloadCreateSchedule struct {
	Name                string `json:"name"`
	Comment             string `json:"comment"`
	FirstTimeMinute     int    `json:"first_time_minute"`
	FirstTimeHour       int    `json:"first_time_hour"`
	FirstTimeDayOfMonth int    `json:"first_time_day_of_month"`
	FirstTimeMonth      int    `json:"first_time_month"`
	FirstTimeYear       int    `json:"first_time_year"`
	Duration            int    `json:"duration"`
	DurationUnit        string `json:"duration_unit"`
	Period              int    `json:"period"`
	PeriodUnit          string `json:"period_unit"`
	Timezone            string `json:"timezone"`
}
type payloadUpdateSchedule struct {
	ScheduleId          string `json:"schedule_id"`
	Name                string `json:"name"`
	Comment             string `json:"comment"`
	FirstTimeMinute     int    `json:"first_time_minute"`
	FirstTimeHour       int    `json:"first_time_hour"`
	FirstTimeDayOfMonth int    `json:"first_time_day_of_month"`
	FirstTimeMonth      int    `json:"first_time_month"`
	FirstTimeYear       int    `json:"first_time_year"`
	Duration            int    `json:"duration"`
	DurationUnit        string `json:"duration_unit"`
	Period              int    `json:"period"`
	PeriodUnit          string `json:"period_unit"`
	Timezone            string `json:"timezone"`
}

// get all schedules
func GetSchedules(c *gin.Context) {
	logrus.Info("\n\n-------Get all schedules in scanner--------")
	var payload payloadGetSchedules
	payload.Trash = false
	payload.Tasks = false
	_ = c.ShouldBindQuery(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.GetSchedules(payload.Filter, payload.FilterId, payload.Trash, payload.Tasks)
	if err != nil {
		logrus.Error("Get schedule fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get schedules success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func GetSchedule(c *gin.Context) {
	logrus.Info("\n\n-------Get schedule in scanner--------")
	scheduleId := c.Param("schedule_id")
	govas := AuthConnection(c)
	respond, err := govas.GetSchedule(scheduleId)
	if err != nil {
		logrus.Error("Get schedule fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get schedule success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
func CloneSchedule(c *gin.Context) {
	logrus.Info("\n\n-------Get schedule in scanner--------")
	scheduleId := c.Param("schedule_id")
	govas := AuthConnection(c)
	respond, err := govas.CloneSchedule(scheduleId)
	if err != nil {
		logrus.Error("Get schedule fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get schedule success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func DeleteSchedule(c *gin.Context) {
	logrus.Info("\n\n-------Get schedule in scanner--------")
	scheduleId := c.Param("schedule_id")
	ultimate := String2Bool(c.Param("ultimate"))
	govas := AuthConnection(c)
	respond, err := govas.DeleteSchedule(scheduleId, ultimate)
	if err != nil {
		logrus.Error("Delete schedule fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Delete schedule success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func CreateSchedule(c *gin.Context) {
	logrus.Info("\n\n-------Create schedules in scanner--------")
	var payload payloadCreateSchedule
	payload.Duration = -1
	payload.FirstTimeDayOfMonth = -1
	payload.FirstTimeHour = -1
	payload.FirstTimeMinute = -1
	payload.FirstTimeMonth = -1
	payload.FirstTimeYear = -1
	payload.Period = -1
	_ = c.ShouldBindJSON(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.CreateSchedule(payload.Name, payload.Comment, payload.FirstTimeMinute, payload.FirstTimeHour, payload.FirstTimeDayOfMonth, payload.FirstTimeMonth,
		payload.FirstTimeYear, payload.Duration, payload.DurationUnit, payload.Period, payload.PeriodUnit, payload.Timezone)
	if err != nil {
		logrus.Error("Create schedule fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create schedules success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
func UpdateSchedule(c *gin.Context) {
	logrus.Info("\n\n-------Create schedules in scanner--------")
	var payload payloadUpdateSchedule
	payload.Duration = -1
	payload.FirstTimeDayOfMonth = -1
	payload.FirstTimeHour = -1
	payload.FirstTimeMinute = -1
	payload.FirstTimeMonth = -1
	payload.FirstTimeYear = -1
	payload.Period = -1
	_ = c.ShouldBindJSON(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.UpdateSchedule(payload.ScheduleId, payload.Name, payload.Comment, payload.FirstTimeMinute, payload.FirstTimeHour, payload.FirstTimeDayOfMonth, payload.FirstTimeMonth,
		payload.FirstTimeYear, payload.Duration, payload.DurationUnit, payload.Period, payload.PeriodUnit, payload.Timezone)
	if err != nil {
		logrus.Error("Create schedule fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create schedules success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
