package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"strings"
)

// get all task
func (gomd *Govasmd) GetTasks(
	filter string,
	filterId string,
	trash bool,
	details bool,
	schedulesOnly bool) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_tasks")
	addFilter(cmd, filter, filterId)
	cmd.CreateAttr("details", toString(details))
	cmd.CreateAttr("trash", toString(trash))
	cmd.CreateAttr("schedules_only", toString(schedulesOnly))
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) GetTask(taskId string) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_tasks")
	cmd.CreateAttr("task_id", taskId)
	cmd.CreateAttr("details", "1")
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) StartTask(taskId string) (string, error) {
	if taskId == "" {
		err := errors.New("StartTask requires a taskId argument")
		logrus.Error("In StartTask()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("start_task")
	cmd.CreateAttr("task_id", taskId)
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) ResumeTask(taskId string) (string, error) {
	if taskId == "" {
		err := errors.New("resumeTask requires a taskId argument")
		logrus.Error("In resumeTask()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("resume_task")
	cmd.CreateAttr("task_id", taskId)
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) StopTask(taskId string) (string, error) {
	if taskId == "" {
		err := errors.New("stopTask requires a taskId argument")
		logrus.Error("In stopTask()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("stop_task")
	cmd.CreateAttr("task_id", taskId)
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) CloneTask(taskId string) (string, error) {
	if taskId == "" {
		err := errors.New("createTask requires a taskId argument")
		logrus.Error("In createTask()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_task")
	cmd.CreateElement("copy").CreateCharData(taskId)
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) DeleteTask(taskId string, ultimate bool) (string, error) {
	if taskId == "" {
		err := errors.New("DeleteTask requires a taskId argument")
		logrus.Error("In DeleteTask()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("delete_task")
	cmd.CreateAttr("task_id", taskId)
	cmd.CreateAttr("ultimat", toString(ultimate))
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) CreateTask(
	name string,
	configId string,
	targetId string,
	scannerId string,
	alterable int,
	hostsOrdering string,
	scheduleId string,
	alertIds []string,
	comment string,
	schedulePeriods int,
	observers []string) (string, error) {
	if name == "" {
		err := errors.New("CreateTask requires a name argument")
		logrus.Error("In CreateTask()", err)
		return "", err
	}
	if configId == "" {
		err := errors.New("CreateTask requires a configId argument")
		logrus.Error("In CreateTask()", err)
		return "", err
	}

	if targetId == "" {
		err := errors.New("CreateTask requires a targetId argument")
		logrus.Error("In CreateTask()", err)
		return "", err
	}

	if scannerId == "" {
		err := errors.New("CreateTask requires a scannerId argument")
		logrus.Error("In CreateTask()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_task")
	cmd.CreateElement("name").CreateCharData(name)
	cmd.CreateElement("config").CreateAttr("id", configId)
	cmd.CreateElement("target").CreateAttr("id", targetId)
	cmd.CreateElement("scanner").CreateAttr("id", scannerId)
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}
	if (alterable >= 0) && (alterable <= 1) {
		cmd.CreateElement("alterable").CreateCharData(string(alterable))
	}

	if hostsOrdering != "" {
		// not sure about the possible values for hosts_orderning
		// it seems gvmd does not check the param
		// gsa allows to select 'sequential', 'random' or 'reverse'
		cmd.CreateElement("hosts_ordering").CreateCharData(hostsOrdering)
	}
	if len(alertIds) > 0 {
		for _, alert := range alertIds {
			cmd.CreateElement("alert").CreateAttr("id", alert)
		}
	}
	if scheduleId != "" {
		cmd.CreateElement("schedule").CreateAttr("id", scheduleId)
		if schedulePeriods != -1 {
			cmd.CreateElement("schedule_periods").CreateCharData(string(schedulePeriods))
		}
	}
	if observers != nil {
		// gvmd splits by comma and space
		// gvmd tries to lookup each value as user name and afterwards as
		// user id. So both user name and user id are possible
		cmd.CreateElement("observers").CreateCharData(strings.Join(observers, ","))
	}
	return gomd.sendXmlCommand(doc)

}

func (gomd *Govasmd) UpdateTask(
	taskId string,
	name string,
	configId string,
	targetId string,
	scannerId string,
	alterable int,
	hostsOrdering string,
	scheduleId string,
	alertIds []string,
	comment string,
	schedulePeriods int,
	observers []string) (string, error) {
	if taskId == "" {
		err := errors.New("UpdateTask requires a taskId argument")
		logrus.Error("In UpdateTask() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("modify_task")
	cmd.CreateAttr("task_id", taskId)
	if name != "" {
		cmd.CreateElement("name").CreateCharData(name)
	}
	if configId != "" {
		cmd.CreateElement("config").CreateAttr("id", configId)
	}
	if targetId != "" {
		cmd.CreateElement("target").CreateAttr("id", targetId)
	}
	if scannerId != "" {
		cmd.CreateElement("scanner").CreateAttr("id", scannerId)
	}
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}
	if (alterable >= 0) && (alterable <= 1) {
		cmd.CreateElement("alterable").CreateCharData(string(alterable))
	}
	if hostsOrdering != "" {
		// not sure about the possible values for hosts_orderning
		// it seems gvmd does not check the param
		// gsa allows to select 'sequential', 'random' or 'reverse'
		cmd.CreateElement("hosts_ordering").CreateCharData(hostsOrdering)
	}
	if len(alertIds) > 0 {
		for _, alert := range alertIds {
			cmd.CreateElement("alert").CreateAttr("id", alert)
		}
	}
	if scheduleId != "" {
		cmd.CreateElement("schedule").CreateAttr("id", scheduleId)
		if schedulePeriods != -1 {
			cmd.CreateElement("schedule_periods").CreateCharData(string(schedulePeriods))
		}
	}
	if observers != nil {
		// gvmd splits by comma and space
		// gvmd tries to lookup each value as user name and afterwards as
		// user id. So both user name and user id are possible
		cmd.CreateElement("observers").CreateCharData(strings.Join(observers, ","))
	}
	return gomd.sendXmlCommand(doc)

}
