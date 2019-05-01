package protocal

import (
	"errors"
	"fmt"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
)

// get all report
func (gomd *Govasmd) GetReports(
	filter string,
	filterId string,
	noteDetails int,
	overrideDetails int) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_reports")
	if filter != "" {
		cmd.CreateAttr("report_filter", filter)
	}
	if filterId != "" {
		cmd.CreateAttr("report_filt_id", filterId)
	}
	fmt.Println(checkIntDefine(noteDetails), "note:   ", noteDetails)
	if checkIntDefine(noteDetails) {
		cmd.CreateAttr("note_details", string(noteDetails))
	}
	if checkIntDefine(overrideDetails) {
		cmd.CreateAttr("override_details", string(overrideDetails))
	}
	cmd.CreateAttr("ignore_pagination", "1")
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) GetReport(reportId, filter, filterID, deltaReportId, reportFormatId string) (string, error) {
	if reportId == "" {
		err := errors.New("GetReport requires a reportId argument")
		logrus.Error("In GetReport() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_reports")
	cmd.CreateAttr("report_id", reportId)
	addFilter(cmd, filter, filterID)
	if deltaReportId != "" {
		cmd.CreateAttr("delta_report_id", deltaReportId)
	}
	if reportFormatId != "" {
		cmd.CreateAttr("format_id", reportFormatId)
	}
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) DeleteReport(reportId string) (string, error) {
	if reportId == "" {
		err := errors.New("DeleteReport requires a reportId argument")
		logrus.Error("In DeleteReport() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("delete_report")
	cmd.CreateAttr("report_id", reportId)
	return gomd.sendXmlCommand(doc)
}
