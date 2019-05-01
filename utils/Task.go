package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type payloadGetTasks struct {
	Filter        string `form:"filter"`
	FilterId      string `form:"filter_id"`
	Trash         bool   `form:"trash"`
	Details       bool   `form:"details"`
	SchedulesOnly bool   `form:"schedules_only"`
}
type payloadCreateTask struct {
	Name            string   `json:"name"`
	ConfigId        string   `json:"config_id"`
	TargetId        string   `json:"target_id"`
	ScannerId       string   `json:"scanner_id"`
	Alterable       int      `json:"alterable"`
	HostsOrdering   string   `json:"hosts_ordering"`
	ScheduleId      string   `json:"schedule_id"`
	AlertIds        []string `json:"alert_ids"`
	Comment         string   `json:"comment"`
	SchedulePeriods int      `json:"schedule_periods"`
	Observers       []string `json:"observers"`
}
type payloadUpdateTask struct {
	TaskId          string   `json:"task_id"`
	Name            string   `json:"name"`
	ConfigId        string   `json:"config_id"`
	TargetId        string   `json:"target_id"`
	ScannerId       string   `json:"scanner_id"`
	Alterable       int      `json:"alterable"`
	HostsOrdering   string   `json:"hosts_ordering"`
	ScheduleId      string   `json:"schedule_id"`
	AlertIds        []string `json:"alert_ids"`
	Comment         string   `json:"comment"`
	SchedulePeriods int      `json:"schedule_periods"`
	Observers       []string `json:"observers"`
}

// get all tasks
func GetTasks(c *gin.Context) {
	logrus.Info("\n\n-------Get all tasks in scanner--------")
	var payload payloadGetTasks
	//_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	_ = c.ShouldBindQuery(&payload)
	fmt.Println(payload)
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
func GetTask(c *gin.Context) {
	logrus.Info("\n\n-------Get task in scanner--------")
	taskId := c.Param("task_id")
	govas := AuthConnection(c)
	respond, err := govas.GetTask(taskId)
	if err != nil {
		logrus.Error("Get task fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get task success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

// start, clone, resume, stop task
func ActionTask(c *gin.Context) {
	logrus.Info("\n\n-------Action task in scanner--------")
	taskId := c.Param("task_id")
	govas := AuthConnection(c)
	var respond string
	var err error
	switch strings.ToUpper(c.Param("action")) {
	case "START":
		respond, err = govas.StartTask(taskId)
	case "STOP":
		respond, err = govas.StopTask(taskId)
	case "RESUME":
		respond, err = govas.ResumeTask(taskId)
	case "CLONE":
		respond, err = govas.CloneTask(taskId)
	default:
		err = errors.New("Unknown action")
	}
	if err != nil {
		logrus.Error("Action with task fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get task success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func DeleteTask(c *gin.Context) {
	logrus.Info("\n\n-------Delete target in scanner--------")
	id := c.Param("task_id")
	ultimate := String2Bool(c.Param("ultimate"))
	govas := AuthConnection(c)
	respond, err := govas.DeleteTask(id, ultimate)
	if err != nil {
		logrus.Error("Delete target fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Delete target success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func CreateTask(c *gin.Context) {
	logrus.Info("\n\n-------Create tasks in scanner--------")
	var payload payloadCreateTask
	payload.SchedulePeriods = -1
	payload.Alterable = 0

	//_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	_ = c.ShouldBindJSON(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.CreateTask(payload.Name, payload.ConfigId, payload.TargetId, payload.ScannerId, payload.Alterable, payload.HostsOrdering, payload.ScheduleId, payload.AlertIds,
		payload.Comment, payload.SchedulePeriods, payload.Observers)
	if err != nil {
		logrus.Error("Create task fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create tasks success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func UpdateTask(c *gin.Context) {
	logrus.Info("\n\n-------Update tasks in scanner--------")
	var payload payloadUpdateTask
	payload.SchedulePeriods = -1
	payload.Alterable = -1
	//_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	_ = c.ShouldBindJSON(&payload)
	fmt.Println(payload)
	govas := AuthConnection(c)
	respond, err := govas.UpdateTask(payload.TaskId, payload.Name, payload.ConfigId, payload.TargetId, payload.ScannerId, payload.Alterable, payload.HostsOrdering, payload.ScheduleId, payload.AlertIds,
		payload.Comment, payload.SchedulePeriods, payload.Observers)
	if err != nil {
		logrus.Error("Update task fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Update tasks success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
