package models

type LocalNode struct {
	NodeId   int16
	Active   string
	NodeName string
	Ip       string
	Type     string
}

type Awareness struct {
	Masters []string
	Slaves  []string
	Witness []string
}
