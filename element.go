package hyperscript

import "syscall/js"

type (
	Element struct {
		NodeName   string
		NodeType   int
		Children   VNodes
		Attributes Object
		Base       js.Value
	}
)

func (vn *Element) GetNodeName() string {
	return vn.NodeName
}

func (vn *Element) GetNodeType() int {
	return vn.NodeType
}

func (vn *Element) GetChildren() VNodes {
	return vn.Children
}
