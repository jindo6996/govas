package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
)

func (gomd *Govasmd) Restore(entityID string) (string, error) {
	if entityID == "" {
		err := errors.New("Restore requires a reportId argument")
		logrus.Error("In Restore() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("restore")
	cmd.CreateAttr("id", entityID)
	return gomd.sendXmlCommand(doc)
}
