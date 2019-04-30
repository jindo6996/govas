package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type payloadCreateTarget struct {
	Name               string   `json:"name"`
	MakeUnique         bool     `json:"make_unique"`
	AssetHostsFilter   string   `json:"asset_hosts_filter"`
	Hosts              []string `json:"hosts"`
	Comment            string   `json:"comment"`
	ExcludeHosts       []string `json:"exclude_hosts"`
	SshCredentialId    string   `json:"ssh_credential_id"`
	SshCredentialPort  int      `json:"ssh_credential_port"`
	SmbCredentialId    string   `json:"smb_credential_id"`
	SnmpCredentialId   string   `json:"snmp_credential_id"`
	EsxiCredentialId   string   `json:"esxi_credential_id"`
	AliveTests         string   `json:"alive_tests"`
	ReverseLookupOnly  bool     `json:"reverse_lookup_only"`
	ReverseLookupUnify bool     `json:"reverse_lookup_unify "`
	PortRange          string   `json:"port_range "`
	PortListId         string   `json:"port_list_id "`
}

type payloadUpdateTarget struct {
	TargetId           string   `json:"target_id"`
	Name               string   `json:"name"`
	Hosts              []string `json:"hosts"`
	Comment            string   `json:"comment"`
	ExcludeHosts       []string `json:"exclude_hosts"`
	SshCredentialId    string   `json:"ssh_credential_id"`
	SshCredentialPort  int      `json:"ssh_credential_port"`
	SmbCredentialId    string   `json:"smb_credential_id"`
	SnmpCredentialId   string   `json:"snmp_credential_id"`
	EsxiCredentialId   string   `json:"esxi_credential_id"`
	AliveTests         string   `json:"alive_tests"`
	ReverseLookupOnly  bool     `json:"reverse_lookup_only"`
	ReverseLookupUnify bool     `json:"reverse_lookup_unify "`
	PortListId         string   `json:"port_list_id "`
}

func GetTargets(c *gin.Context) {
	logrus.Info("\n\n-------Get all target in scanner--------")
	trash := String2Bool(c.Param("trash"))
	tasks := String2Bool(c.Param("tasks"))
	govas := AuthConnection(c)
	fmt.Println("trash: ", trash, "task: ", tasks)

	respond, err := govas.GetTargets("", "", trash, tasks)
	if err != nil {
		logrus.Error("Get target fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get target success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}
func GetTarget(c *gin.Context) {
	logrus.Info("\n\n-------Get all target in scanner--------")
	id := c.Param("id")
	govas := AuthConnection(c)
	respond, err := govas.GetTarget(id)
	if err != nil {
		logrus.Error("Get target fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Get target success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func CreateTarget(c *gin.Context) {
	logrus.Info("\n\n------------Create target----------")
	var payload payloadCreateTarget
	_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	govas := AuthConnection(c)
	respond, err := govas.CreateTarget(payload.Name, payload.MakeUnique, payload.AssetHostsFilter, payload.Hosts, payload.Comment, payload.ExcludeHosts, payload.SshCredentialId, payload.SshCredentialPort,
		payload.SmbCredentialId, payload.SnmpCredentialId, payload.EsxiCredentialId, payload.AliveTests, payload.ReverseLookupOnly, payload.ReverseLookupUnify, payload.PortRange, payload.PortListId)
	if err != nil {
		logrus.Info("Create user fail")
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Create target success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()

}

func DeleteTarget(c *gin.Context) {
	logrus.Info("\n\n-------Delete target in scanner--------")
	id := c.Param("id")
	ultimate := String2Bool(c.Param("ultimate"))
	govas := AuthConnection(c)
	respond, err := govas.DeleteTarget(id, ultimate)
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
func CloneTarget(c *gin.Context) {
	logrus.Info("\n\n-------clone target in scanner--------")
	id := c.Param("id")
	govas := AuthConnection(c)
	respond, err := govas.CloneTarget(id)
	if err != nil {
		logrus.Error("clone target fail: ", err)
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("clone target success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()
}

func UpdateTarget(c *gin.Context) {
	logrus.Info("\n\n------------Update target----------")
	var payload payloadUpdateTarget
	payload.ExcludeHosts = []string{""}
	_ = json.NewDecoder(c.Request.Body).Decode(&payload)
	fmt.Println(payload)
	fmt.Println("id: ", payload.TargetId)
	govas := AuthConnection(c)
	respond, err := govas.UpdateTarget(payload.TargetId, payload.Name, payload.Hosts, payload.Comment, payload.ExcludeHosts, payload.SshCredentialId, payload.SshCredentialPort,
		payload.SmbCredentialId, payload.SnmpCredentialId, payload.EsxiCredentialId, payload.AliveTests, payload.ReverseLookupOnly, payload.ReverseLookupUnify, payload.PortListId)
	if err != nil {
		logrus.Info("Update user fail")
		respond = CreatJsonErro(err.Error())
		c.Data(http.StatusBadRequest, "application/json; charset=utf-8", []byte(respond))
	} else {
		logrus.Info("Update target success")
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(respond))
	}
	_ = govas.Disconnect()

}
