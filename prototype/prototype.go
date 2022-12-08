package prototype

type Master struct {
	nodes []Node
	name  string
}

type Node struct {
	name string
}

func (node *Node) clone() Node {
	return Node{name: node.name + "_clone"}
}

func (m *Master) Clone() Master {
	cloneMaster := &Master{
		name: m.name + "_clone",
	}
	var tmpNodes []Node
	for _, node := range m.nodes {
		copy := node.clone()
		tmpNodes = append(tmpNodes, copy)
	}
	cloneMaster.nodes = tmpNodes
	return *cloneMaster
}
