package protocal

import "github.com/beevik/etree"

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
