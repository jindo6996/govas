package protocal

import (
	"fmt"
	"github.com/beevik/etree"
	"os/exec"
)

func (gomd *Govasmd) SynNVT() (string, error) {
	//doc := etree.NewDocument()
	//doc.CreateElement("sync_feed")
	//return gomd.sendXmlCommand(doc)
	var cmdName = "greenbone-nvt-sync"
	var cmd = exec.Command(cmdName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Printf("%v\n", string(output))
	return string(output), nil
}
func (gomd *Govasmd) SynCert() (string, error) {
	//doc := etree.NewDocument()
	//doc.CreateElement("sync_cert")
	//return gomd.sendXmlCommand(doc)
	var cmdName = "greenbone-certdata-sync"
	var cmd = exec.Command(cmdName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Printf("%v\n", string(output))
	return string(output), nil
}
func (gomd *Govasmd) SynSCAP() (string, error) {
	//doc := etree.NewDocument()
	//doc.CreateElement("sync_scap")
	//return gomd.sendXmlCommand(doc)
	var cmdName = "greenbone-scapdata-sync"
	var cmd = exec.Command(cmdName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Printf("%v\n", string(output))
	return string(output), nil
}
func (gomd *Govasmd) GetFeeds() (string, error) {
	doc := etree.NewDocument()
	doc.CreateElement("get_feeds")
	return gomd.sendXmlCommand(doc)
}
