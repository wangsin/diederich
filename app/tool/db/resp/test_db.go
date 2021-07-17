package dbresp

type EchoDBConf struct {
	Echo interface{}            `json:"echo"`
	Conf map[string]interface{} `json:"conf"`
}
