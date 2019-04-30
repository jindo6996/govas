package protocal

import "github.com/beevik/etree"

// get all task
func (gomd *Govasmd) GetTasks(
	filter string,
	filterId string,
	trash bool,
	details bool,
	schedulesOnly bool) (string, error) {
	doc := etree.NewDocument()
	cmd := doc.CreateElement("get_tasks")
	addFilter(cmd, filter, filterId)
	cmd.CreateAttr("details", toString(details))
	cmd.CreateAttr("schedules_only", toString(schedulesOnly))
	return gomd.sendXmlCommand(doc)
}
