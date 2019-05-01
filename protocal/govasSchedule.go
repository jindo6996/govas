package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"strconv"
)

var TimeUnits = map[string]int{
	"second": 1,
	"minute": 1,
	"hour":   1,
	"day":    1,
	"week":   1,
	"month":  1,
	"year":   1,
	"decade": 1,
}

// get all schedule
func (gomd *Govasmd) GetSchedules(filter, filterId string, trash, tasks bool) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_schedules")
	addFilter(cmd, filter, filterId)
	cmd.CreateAttr("trash", toString(trash))
	if !trash {
		cmd.CreateAttr("tasks", toString(trash))
	}
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) GetSchedule(scheduleId string) (string, error) {
	if scheduleId == "" {
		err := errors.New("create_target requires a scheduleId argument")
		logrus.Error("In create_target()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_schedules")
	cmd.CreateAttr("schedule_id", scheduleId)
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) CloneSchedule(scheduleId string) (string, error) {
	if scheduleId == "" {
		err := errors.New("CloneSchedule requires a scheduleId argument")
		logrus.Error("In CloneSchedule() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_schedule")
	cmd.CreateElement("copy").CreateCharData(scheduleId)
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) DeleteSchedule(scheduleId string, ultimate bool) (string, error) {
	if scheduleId == "" {
		err := errors.New("DeleteSchedule requires a scheduleId argument")
		logrus.Error("In DeleteSchedule() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("delete_schedule")
	cmd.CreateAttr("schedule_id", scheduleId)
	cmd.CreateAttr("ultimate", toString(ultimate))
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) CreateSchedule(
	name, comment string, firstTimeMinute, firstTimeHour, firstTimeDayOfMonth,
	firstTimeMonth, firstTimeYear, duration int, durationUnit string, period int,
	periodUnit, timezone string) (string, error) {
	if name == "" {
		err := errors.New("create_schedule requires a name argument")
		logrus.Error("In create_schedule()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_schedule")
	cmd.CreateElement("name").CreateCharData(name)
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}
	if (checkIntDefine2(firstTimeMinute)) || (checkIntDefine2(firstTimeHour)) ||
		(checkIntDefine2(firstTimeDayOfMonth)) || (checkIntDefine2(firstTimeMonth)) ||
		(checkIntDefine2(firstTimeYear)) {
		if !checkIntDefine2(firstTimeMinute) {
			err := errors.New("Setting first_time requires firstTimeMinute argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if !checkIntDefine2(firstTimeHour) {
			err := errors.New("Setting first_time requires firstTimeHour argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}

		if !checkIntDefine2(firstTimeDayOfMonth) {
			err := errors.New("Setting first_time requires firstTimeDayOfMonth argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if !checkIntDefine2(firstTimeMonth) {
			err := errors.New("Setting first_time requires firstTimeMonth argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if !checkIntDefine2(firstTimeYear) {
			err := errors.New("Setting first_time requires firstTimeYear argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		timeXml := cmd.CreateElement("first_time")
		timeXml.CreateElement("minute").CreateCharData(strconv.Itoa(firstTimeMinute))
		timeXml.CreateElement("hour").CreateCharData(strconv.Itoa(firstTimeHour))
		timeXml.CreateElement("day_of_month").CreateCharData(strconv.Itoa(firstTimeDayOfMonth))
		timeXml.CreateElement("month").CreateCharData(strconv.Itoa(firstTimeMonth))
		timeXml.CreateElement("year").CreateCharData(strconv.Itoa(firstTimeYear))
	}
	if checkIntDefine2(duration) {
		if durationUnit == "" {
			err := errors.New("Setting duration requires duration_unit argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if TimeUnits[durationUnit] != 1 {
			err := errors.New("Invalid durationUnit: " + durationUnit)
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		durationXml := cmd.CreateElement("duration")
		durationXml.CreateCharData(strconv.Itoa(duration))
		durationXml.CreateElement("unit").CreateCharData(durationUnit)
	}
	if checkIntDefine2(period) {
		if periodUnit == "" {
			err := errors.New("Setting period requires period_unit argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if TimeUnits[periodUnit] != 1 {
			err := errors.New("Invalid periodUnit: " + periodUnit)
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		periodXml := cmd.CreateElement("period")
		periodXml.CreateCharData(strconv.Itoa(period))
		periodXml.CreateElement("unit").CreateCharData(periodUnit)
	}
	if timezone != "" {
		cmd.CreateElement("timezone").CreateCharData(timezone)
	}
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) UpdateSchedule(
	scheduleId, name, comment string, firstTimeMinute, firstTimeHour, firstTimeDayOfMonth,
	firstTimeMonth, firstTimeYear, duration int, durationUnit string, period int,
	periodUnit, timezone string) (string, error) {
	if scheduleId == "" {
		err := errors.New("create_schedule requires a scheduleId argument")
		logrus.Error("In create_schedule()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("modify_schedule")
	cmd.CreateAttr("schedule_id", scheduleId)
	if name != "" {
		cmd.CreateElement("name").CreateCharData(name)
	}
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}
	if (checkIntDefine2(firstTimeMinute)) || (checkIntDefine2(firstTimeHour)) ||
		(checkIntDefine2(firstTimeDayOfMonth)) || (checkIntDefine2(firstTimeMonth)) ||
		(checkIntDefine2(firstTimeYear)) {
		if !checkIntDefine2(firstTimeMinute) {
			err := errors.New("Setting first_time requires firstTimeMinute argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if !checkIntDefine2(firstTimeHour) {
			err := errors.New("Setting first_time requires firstTimeHour argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}

		if !checkIntDefine2(firstTimeDayOfMonth) {
			err := errors.New("Setting first_time requires firstTimeDayOfMonth argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if !checkIntDefine2(firstTimeMonth) {
			err := errors.New("Setting first_time requires firstTimeMonth argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if !checkIntDefine2(firstTimeYear) {
			err := errors.New("Setting first_time requires firstTimeYear argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		timeXml := cmd.CreateElement("first_time")
		timeXml.CreateElement("minute").CreateCharData(strconv.Itoa(firstTimeMinute))
		timeXml.CreateElement("hour").CreateCharData(strconv.Itoa(firstTimeHour))
		timeXml.CreateElement("day_of_month").CreateCharData(strconv.Itoa(firstTimeDayOfMonth))
		timeXml.CreateElement("month").CreateCharData(strconv.Itoa(firstTimeMonth))
		timeXml.CreateElement("year").CreateCharData(strconv.Itoa(firstTimeYear))
	}
	if checkIntDefine2(duration) {
		if durationUnit == "" {
			err := errors.New("Setting duration requires duration_unit argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if TimeUnits[durationUnit] != 1 {
			err := errors.New("Invalid durationUnit: " + durationUnit)
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		durationXml := cmd.CreateElement("duration")
		durationXml.CreateCharData(strconv.Itoa(duration))
		durationXml.CreateElement("unit").CreateCharData(durationUnit)
	}
	if checkIntDefine2(period) {
		if periodUnit == "" {
			err := errors.New("Setting period requires period_unit argument")
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		if TimeUnits[periodUnit] != 1 {
			err := errors.New("Invalid periodUnit: " + periodUnit)
			logrus.Error("In CreateSchedule()", err)
			return "", err
		}
		periodXml := cmd.CreateElement("period")
		periodXml.CreateCharData(strconv.Itoa(period))
		periodXml.CreateElement("unit").CreateCharData(periodUnit)
	}
	if timezone != "" {
		cmd.CreateElement("timezone").CreateCharData(timezone)
	}
	return gomd.sendXmlCommand(doc)
}
