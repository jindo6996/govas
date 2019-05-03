package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"strings"
)

// Get Roles
func (gomd *Govasmd) GetRoles(filter string, filter_id string) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_roles")
	addFilter(cmd, filter, filter_id)
	return gomd.sendXmlCommand(doc)
}

// Get one user on scanner with uuid
func (gomd *Govasmd) GetRole(uuid string) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_roles")
	cmd.CreateAttr("user_id", uuid)
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) DeleteRole(roleID string, ultimate bool) (string, error) {
	if roleID == "" {
		err := errors.New("DeleteRole requires a roleID argument")
		logrus.Error("In DeleteRole() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("delete_role")
	cmd.CreateAttr("role_id", roleID)
	cmd.CreateAttr("ultimate", toString(ultimate))
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) CloneRole(roleId string) (string, error) {
	if roleId == "" {
		err := errors.New("CloneRole requires a roleId argument")
		logrus.Error("In CloneRole() ", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_role")
	cmd.CreateElement("copy").CreateCharData(roleId)
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) CreateRole(name string, comment string, users []string) (string, error) {
	if name == "" {
		err := errors.New("CreateRole Requires name")
		logrus.Error("In CreateRole()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_role")
	cmd.CreateElement("name").CreateCharData(name)
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}
	if len(users) > 0 {
		cmd.CreateElement("users").CreateCharData(strings.Join(users, ","))
	}
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) UpdateRole(roleId, name, comment string, users []string) (string, error) {
	if name == "" {
		err := errors.New("UpdateRole Requires roleId")
		logrus.Error("In UpdateRole()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("modify_role")
	cmd.CreateAttr("role_id", roleId)
	if name != "" {
		cmd.CreateElement("name").CreateCharData(name)
	}
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}
	if len(users) > 0 {
		cmd.CreateElement("users").CreateCharData(strings.Join(users, ","))
	}
	return gomd.sendXmlCommand(doc)
}
