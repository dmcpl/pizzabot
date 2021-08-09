package algorithms

import (
	"example.com/slice/pizzabot/grid"
)

func NodeTreeBruteForce(path *grid.Path) []*grid.Path {
	//todo unimplemented
	//origin := &grid.Point{0, 0}
	return nil
}

func NodeTreeClosestPoint(path *grid.Path) []*grid.Path {
	origin := &grid.Point{0, 0}
	root := buildTree(origin, path)

	acc := make([]*grid.Path, 0)
	traverseTree(root, nil, &acc)

	return acc
}

type node struct {
	Point    *grid.Point
	Rest     *grid.Path
	SubNodes []*node
}

func newNode(point *grid.Point, rest *grid.Path) *node {
	return &node{Point: point,
		Rest:     rest,
		SubNodes: make([]*node, 0, 3)}
}

type NodeList struct {
	Nodes []*node
}

func buildTree(point *grid.Point, rest *grid.Path) *node {
	node := newNode(point, rest)

	if len(rest.Points) == 0 {
		return node
	}

	c := findClosestPoints(point, rest)
	for _, p := range c.Points {
		node.SubNodes = append(node.SubNodes, buildTree(p, rest.Remove(p)))
	}

	return node
}

func traverseTree(root *node, nodePath *NodeList, acc *[]*grid.Path) {
	if nodePath == nil {
		nodePath = &NodeList{Nodes: make([]*node, 0)}
	}
	nodePath.Nodes = append(nodePath.Nodes, root)

	if len(root.SubNodes) == 0 {
		p := grid.NewPath()
		for _, n := range nodePath.Nodes {
			p.Points = append(p.Points, n.Point)
		}

		*acc = append(*acc, p)
		nodePath.Nodes = nodePath.Nodes[:len(nodePath.Nodes)-1]
		return
	}

	for _, n := range root.SubNodes {
		traverseTree(n, nodePath, acc)
	}

	nodePath.Nodes = nodePath.Nodes[:len(nodePath.Nodes)-1]
}
