package prototype

import (
	"reflect"
	"testing"
)

func TestMaster_Clone(t *testing.T) {
	m1 := &Master{
		name: "m1",
		nodes: []Node{
			{name: "node1"},
			{name: "node2"},
		},
	}

	m2 := m1

	m2.name = "m2"
	m2.nodes = []Node{{name: "node3"}, {name: "node4"}}

	if m1.name != m2.name {
		t.Errorf("m1 and m2 name %v !=  %v", m1.name, m2.name)
	}

	if !reflect.DeepEqual(m1.nodes, m2.nodes) {
		t.Errorf("m1 and m2 nodes %v !=  %v", m1.nodes, m2.nodes)
	}

	m1.name = "m1"
	m1Clone := m1.Clone()

	if m1.name == m1Clone.name {
		t.Errorf("m1 and m1Clone name %v == %v", m1.name, m1Clone.name)
	}

	if reflect.DeepEqual(m1.nodes, m1Clone.nodes) {
		t.Errorf("m1 and m1Clone nodes %v !=  %v", m1.nodes, m1Clone.nodes)
	}

}
