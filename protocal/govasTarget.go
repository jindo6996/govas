package protocal

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/sirupsen/logrus"
	"strings"
)

var ALIVE_TESTS = map[string]int{
	"Consider Alive":                   1,
	"ICMP, TCP-ACK Service & ARP Ping": 1,
	"TCP-ACK Service & ARP Ping":       1,
	"ICMP & ARP Ping":                  1,
	"ICMP & TCP-ACK Service Ping":      1,
	"ARP Ping":                         1,
	"TCP-ACK Service Ping":             1,
	"TCP-SYN Service Ping":             1,
	"ICMP Ping":                        1,
	"Scan Config Default":              1,
}

func (gomd *Govasmd) GetTargets(filter string, filter_id string, trash bool, tasks bool) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_targets")
	addFilter(cmd, filter, filter_id)
	if trash {
		cmd.CreateAttr("trash", toString(trash))
	}
	if tasks {
		cmd.CreateAttr("tasks", toString(tasks))
	}
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) GetTarget(id string) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_targets")
	cmd.CreateAttr("target_id", id)
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) CreateTarget(
	name string,
	makeUnique bool,
	assetHostsFilter string,
	hosts []string,
	comment string,
	excludeHosts []string,
	sshCredentialId string,
	sshCredentialPort int,
	smbCredentialId string,
	snmpCredentialId string,
	esxiCredentialId string,
	aliveTests string,
	reverseLookupOnly bool,
	reverseLookupUnify bool,
	portRange string,
	portListId string) (string, error) {
	if name == "" {
		err := errors.New("create_target requires a name argument")
		logrus.Error("In create_target()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_target")
	nameXml := cmd.CreateElement("name")
	nameXml.CreateCharData(name)
	nameXml.CreateElement("make_unique").CreateCharData(toString(makeUnique))

	if assetHostsFilter != "" {
		cmd.CreateElement("asset_hosts").CreateAttr("filter", assetHostsFilter)
	} else if len(hosts) > 0 {
		cmd.CreateElement("hosts").CreateCharData(strings.Join(hosts, ","))
	} else {
		err := errors.New("create_target requires either a hosts or an asset_hosts_filter argument")
		logrus.Error("In CreateTarget()", err)
		return "", err
	}

	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}

	if len(excludeHosts) > 0 {
		cmd.CreateElement("exclude_hosts").CreateCharData(strings.Join(excludeHosts, ","))
	}

	if sshCredentialId != "" {
		sshXml := cmd.CreateElement("ssh_credential")
		sshXml.CreateAttr("id", sshCredentialId)
		if sshCredentialPort != 0 {
			sshXml.CreateElement("port").CreateCharData(string(sshCredentialPort))
		}
	}

	if smbCredentialId != "" {
		cmd.CreateElement("smb_credential").CreateAttr("id", smbCredentialId)
	}

	if esxiCredentialId != "" {
		cmd.CreateElement("esxi_credential").CreateAttr("id", esxiCredentialId)
	}

	if snmpCredentialId != "" {
		cmd.CreateElement("snmp_credential").CreateAttr("id", snmpCredentialId)
	}

	if aliveTests != "" {
		if ALIVE_TESTS[aliveTests] != 1 {
			err := errors.New("alive_tests Invalid")
			logrus.Error("In CreateTarget()", err)
			return "", err
		}
		cmd.CreateElement("alive_tests").CreateCharData(aliveTests)
	}

	cmd.CreateElement("reverse_lookup_only").CreateCharData(toString(reverseLookupOnly))
	cmd.CreateElement("reverse_lookup_unify").CreateCharData(toString(reverseLookupUnify))

	if portRange != "" {
		cmd.CreateElement("port_range").CreateCharData(portRange)
	}
	if portListId != "" {
		cmd.CreateElement("port_list").CreateAttr("id", portListId)
	}

	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) DeleteTarget(targetId string, ultimate bool) (string, error) {
	if targetId == "" {
		err := errors.New("DeleteTarget requires a targetId argument")
		logrus.Error("In DeleteTarget()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("delete_target")
	cmd.CreateAttr("target_id", targetId)
	cmd.CreateAttr("ultimate", toString(ultimate))
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) CloneTarget(targetId string) (string, error) {
	if targetId == "" {
		err := errors.New("CloneTarget requires a targetId argument")
		logrus.Error("In CloneTarget()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("create_target")
	cmd.CreateElement("copy").CreateCharData(targetId)
	return gomd.sendXmlCommand(doc)
}

func (gomd *Govasmd) UpdateTarget(
	targetId string,
	name string,
	hosts []string,
	comment string,
	excludeHosts []string,
	sshCredentialId string,
	sshCredentialPort int,
	smbCredentialId string,
	snmpCredentialId string,
	esxiCredentialId string,
	aliveTests string,
	reverseLookupOnly bool,
	reverseLookupUnify bool,
	portListId string) (string, error) {
	if targetId == "" {
		err := errors.New("UpdateTarget requires a targetId argument")
		logrus.Error("In UpdateTarget()", err)
		return "", err
	}
	doc := etree.NewDocument()
	cmd := doc.CreateElement("modify_target")
	cmd.CreateAttr("target_id", targetId)
	nameXml := cmd.CreateElement("name")
	nameXml.CreateCharData(name)
	if len(hosts) > 0 {
		cmd.CreateElement("hosts").CreateCharData(strings.Join(hosts, ","))
	}
	if comment != "" {
		cmd.CreateElement("comment").CreateCharData(comment)
	}

	if len(excludeHosts) > 0 {
		cmd.CreateElement("exclude_hosts").CreateCharData(strings.Join(excludeHosts, ","))
	}

	if sshCredentialId != "" {
		sshXml := cmd.CreateElement("ssh_credential")
		sshXml.CreateAttr("id", sshCredentialId)
		if sshCredentialPort != 0 {
			sshXml.CreateElement("port").CreateCharData(string(sshCredentialPort))
		}
	}

	if smbCredentialId != "" {
		cmd.CreateElement("smb_credential").CreateAttr("id", smbCredentialId)
	}

	if esxiCredentialId != "" {
		cmd.CreateElement("esxi_credential").CreateAttr("id", esxiCredentialId)
	}

	if snmpCredentialId != "" {
		cmd.CreateElement("snmp_credential").CreateAttr("id", snmpCredentialId)
	}

	if aliveTests != "" {
		if ALIVE_TESTS[aliveTests] != 1 {
			err := errors.New("alive_tests Invalid")
			logrus.Error("In CreateTarget()", err)
			return "", err
		}
		cmd.CreateElement("alive_tests").CreateCharData(aliveTests)
	}

	cmd.CreateElement("reverse_lookup_only").CreateCharData(toString(reverseLookupOnly))
	cmd.CreateElement("reverse_lookup_unify").CreateCharData(toString(reverseLookupUnify))

	if portListId != "" {
		cmd.CreateElement("port_list").CreateAttr("id", portListId)
	}

	return gomd.sendXmlCommand(doc)
}
