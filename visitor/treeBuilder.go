package visitor

import (
	"fmt"
	"strings"
)

func (v *TreeShapeListener) AddNode(label string) int {
	id := v.NodeId
	v.Nodes = append(v.Nodes, fmt.Sprintf("  node%d [label=\"%s\"];", id, label))
	v.NodeId++
	return id
}

func (v *TreeShapeListener) AddEdge(parent, child int) {
	v.Edges = append(v.Edges, fmt.Sprintf("  node%d -> node%d;", parent, child))
}

func (v *TreeShapeListener) ToDot() string {
	var sb strings.Builder
	sb.WriteString("digraph G {\n")
	for _, node := range v.Nodes {
		sb.WriteString(node + "\n")
	}
	for _, edge := range v.Edges {
		sb.WriteString(edge + "\n")
	}
	sb.WriteString("}\n")
	return sb.String()
}
