package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"govas/govas"
	"strings"
)

type Govasmd struct {
	Connection    *govas.UnixSocketConnection
	Authenticated bool
	Connected     bool
}

// check connected
func (govas *Govasmd) IsConnected() bool {
	return govas.Connected
}

// connect
func (govas *Govasmd) Connect() {
	if !govas.IsConnected() {
		govas.Connection.Connect()
		govas.Connected = true
	}
}

// Disconnect
func (govas *Govasmd) Disconnect() error {
	if govas.IsConnected() {
		err := govas.Connection.Disconnect()
		if err != nil {
			return err
		}
		govas.Connected = false
		return nil
	}
	return errors.New("You must be AUTHENTICATE")
}

// Get the Greenbone Manager Protocol version used by the remote gvmd
func (gomd *Govasmd) GetVersion() (string, error) {
	doc := etree.NewDocument()
	doc.CreateElement("get_version")
	return gomd.sendXmlCommand(doc)
}

//Checks if the user is authenticated
func (gomd *Govasmd) IsAuthenticated() bool {
	return gomd.Authenticated
}

//Authenticate to gvmd.
func (gomd *Govasmd) Authenticate(username string, password string) (string, error) {
	doc := etree.NewDocument()
	authenticate := doc.CreateElement("authenticate")
	credentials := authenticate.CreateElement("credentials")
	if username == "" {
		err := errors.New("Authenticate requires username")
		logrus.Error("In Authenticate()", err)
		return "", err
	}
	if password == "" {
		err := errors.New("Authenticate requires password")
		logrus.Error("In Authenticate()", err)
		return "", err
	}

	user := credentials.CreateElement("username")
	user.CreateCharData(username)
	passwd := credentials.CreateElement("password")
	passwd.CreateCharData(password)
	resp, _ := gomd.sendXmlCommand(doc)
	value := gjson.Get(resp, "authenticate_response.-status").Int()
	if value == 200 {
		gomd.Authenticated = true
	}
	return resp, nil
}

// Get all user on scanner
func (gomd *Govasmd) GetUsers(filter string, filter_id string) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_users")
	addFilter(cmd, filter, filter_id)
	return gomd.sendXmlCommand(doc)
}

// Get one user on scanner with uuid
func (gomd *Govasmd) GetUser(uuid string) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_users")
	cmd.CreateAttr("user_id", uuid)
	return gomd.sendXmlCommand(doc)
}

// Create user
func (gomd *Govasmd) CreateUser(name string, password string, hosts []string, hostAllow bool, ifaces []string, ifaceallow bool, roleId string) (string, error) {
	if name == "" {
		err := errors.New("Requires username")
		logrus.Error("In CreateUser()", err)
		return "", err
	}
	if password == "" {
		err := errors.New("Requires password")
		logrus.Error("In CreateUser()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_user")
	nameXml := cmd.CreateElement("name")
	nameXml.CreateCharData(name)
	passXml := cmd.CreateElement("password")
	passXml.CreateCharData(password)
	if hosts != nil {
		hostsData := strings.Join(hosts, ",")
		hostsXML := cmd.CreateElement("hosts")
		hostsXML.CreateCharData(hostsData)
		hostsXML.CreateAttr("allow", toString(hostAllow))
	}
	if ifaces != nil {
		ifaceData := strings.Join(ifaces, ",")
		ifaceXML := cmd.CreateElement("ifaces")
		ifaceXML.CreateCharData(ifaceData)
		ifaceXML.CreateAttr("allow", toString(ifaceallow))
	}
	if roleId != "" {
		roleXML := cmd.CreateElement("role")
		roleXML.CreateAttr("id", roleId)
	}
	return gomd.sendXmlCommand(doc)
}

// Delete user
func (gomd *Govasmd) DeleteUser(userID string, name string, inheritorID string, inheritorName string) (string, error) {
	if (userID == "") && (name == "") {
		err := errors.New("delete_user requires a user_id or name argument")
		logrus.Error("In DeleteUser()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("delete_user")
	if userID != "" {
		cmd.CreateAttr("user_id", userID)
	}

	if name != "" {
		cmd.CreateAttr("name", name)
	}
	if inheritorID != "" {
		cmd.CreateAttr("inheritor_id", inheritorID)
	}
	if inheritorName != "" {
		cmd.CreateAttr("inheritor_name", inheritorName)
	}
	return gomd.sendXmlCommand(doc)
}

// update user
func (gomd *Govasmd) UpdateUser(userID string, name string, newName string, password string, hosts []string, hostAllow bool, ifaces []string, ifaceallow bool, roleId string) (string, error) {
	if (userID == "") && (newName == "") {
		err := errors.New("modify_user requires a user_id or name argument")
		logrus.Error("In UpdateUser()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("modify_user")
	if userID != "" {
		cmd.CreateAttr("user_id", userID)
	} else {
		cmd.CreateElement("name").CreateCharData(name)
	}
	if newName != "" {
		cmd.CreateElement("new_name").CreateCharData(newName)
	}
	if password != "" {
		cmd.CreateElement("password").CreateCharData(password)
	}
	if hosts != nil {
		hostsData := strings.Join(hosts, ",")
		hostsXML := cmd.CreateElement("hosts")
		hostsXML.CreateCharData(hostsData)
		hostsXML.CreateAttr("allow", toString(hostAllow))
	}
	if ifaces != nil {
		ifaceData := strings.Join(ifaces, ",")
		ifaceXML := cmd.CreateElement("ifaces")
		ifaceXML.CreateCharData(ifaceData)
		ifaceXML.CreateAttr("allow", toString(ifaceallow))
	}
	if roleId != "" {
		cmd.CreateElement("role").CreateAttr("id", roleId)
	}
	return gomd.sendXmlCommand(doc)
}
