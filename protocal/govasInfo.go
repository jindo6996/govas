package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"strings"
)

var INFO_TYPES = map[string]int{
	"CERT_BUND_ADV": 1,
	"CPE":           1,
	"CVE":           1,
	"DFN_CERT_ADV":  1,
	"OVALDEF":       1,
	"NVT":           1,
	"ALLINFO":       1,
}

func (gomd *Govasmd) GetInfos(infoType, filter, filterId, name string, details bool) (string, error) {
	if infoType == "" {
		err := errors.New("create_target requires a infoType argument")
		logrus.Error("In create_target()", err)
		return "", err
	}
	infoType = strings.ToUpper(infoType)
	if INFO_TYPES[infoType] != 1 {
		err := errors.New("get_info_list info_type argument must be one of CERT_BUND_ADV, CPE, CVE, DFN_CERT_ADV, OVALDEF, NVT or ALLINFO")
		logrus.Error("In GetInfos()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_info")
	cmd.CreateAttr("type", infoType)
	addFilter(cmd, filter, filterId)
	if name != "" {
		cmd.CreateAttr("name", name)
	}
	cmd.CreateAttr("details", toString(details))
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) GetInfo(infoType, infoId string) (string, error) {
	if infoType == "" {
		err := errors.New("create_target requires a infoType argument")
		logrus.Error("In create_target()", err)
		return "", err
	}
	if infoId == "" {
		err := errors.New("create_target requires a infoType argument")
		logrus.Error("In create_target()", err)
		return "", err
	}
	infoType = strings.ToUpper(infoType)
	if INFO_TYPES[infoType] != 1 {
		err := errors.New("get_info_list info_type argument must be one of CERT_BUND_ADV, CPE, CVE, DFN_CERT_ADV, OVALDEF, NVT or ALLINFO")
		logrus.Error("In GetInfos()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_info")
	cmd.CreateAttr("info_id", infoId)
	cmd.CreateAttr("type", infoType)
	cmd.CreateAttr("details", "1")
	return gomd.sendXmlCommand(doc)

}
