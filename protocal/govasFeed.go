package protocal

import "github.com/beevik/etree"

func (gomd *Govasmd) SynNVT() (string, error) {
	doc := etree.NewDocument()
	doc.CreateElement("sync_feed")
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) SynCert() (string, error) {
	doc := etree.NewDocument()
	doc.CreateElement("sync_cert")
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) SynSCAP() (string, error) {
	doc := etree.NewDocument()
	doc.CreateElement("sync_scap")
	return gomd.sendXmlCommand(doc)
}
func (gomd *Govasmd) GetFeeds() (string, error) {
	doc := etree.NewDocument()
	doc.CreateElement("get_feeds")
	return gomd.sendXmlCommand(doc)
}
