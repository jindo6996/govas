package utils

type User struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Comment     string `json:"comment"`
	Role        string `json:"role"`
	Password    string `json:"password"`
	Host_allow  string `json:"host_allow"`
	Hosts       string `json:"hosts"`
	Iface_allow string `json:"iface_allow"`
	Ifaces      string `json:"ifaces"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}
