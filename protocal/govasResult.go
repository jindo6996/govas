package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
)

// get all result
func (gomd *Govasmd) GetResults(
	filter string,
	filterId string,
	taskId string,
	noteDetails int,
	overrideDetails int,
	details int) (string, error) {

	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_results")
	addFilter(cmd, filter, filterId)
	if taskId != "" {
		cmd.CreateAttr("task_id", taskId)
	}
	if checkIntDefine(details) {
		cmd.CreateAttr("note_details", string(details))
	}
	if checkIntDefine(noteDetails) {
		cmd.CreateAttr("note_details", string(noteDetails))
	}
	if checkIntDefine(overrideDetails) {
		cmd.CreateAttr("override_details", string(overrideDetails))
	}
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) GetResult(resultId string) (string, error) {
	if resultId == "" {
		err := errors.New("create_target requires a resultId argument")
		logrus.Error("In create_target()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_results")
	cmd.CreateAttr("result_id", resultId)
	cmd.CreateAttr("details", "1")
	return gomd.sendXmlCommand(doc)
}
