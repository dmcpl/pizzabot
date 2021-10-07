package algorithms

import (
	"github.com/dmcpl/pizzabot/grid"
)

type PathFinderFunc func(point grid.Point, rest grid.Path) grid.Path

// NodeTreeClosestPoint find the closest point pathfinding with branching at
// points with equal distance.
func NodeTreeClosestPoint(path grid.Path) []grid.Path {
	return nodeTree(path, findClosestPoints)
}

// NodeTreeBruteForce extract every single combination of routes, each point branches at
// all other points available. Can get nuts quickly!
// number of paths = (number of points - 1)! (factorial)
func NodeTreeBruteForce(path grid.Path) []grid.Path {
	return nodeTree(path, findAllPoints)
}

// nodeTree build and traverse node tree and return all viable paths extracted
func nodeTree(path grid.Path, pathfinderFn PathFinderFunc) []grid.Path {
	origin := grid.Point{0, 0}
	root := buildTree(origin, path, pathfinderFn)

	acc := make([]grid.Path, 0)
	traverseTree(root, nil, &acc)
	return acc
}

// Node represents a single point in a possibility of paths
type Node struct {
	Point    grid.Point
	Rest     grid.Path
	SubNodes []*Node
}

func newNode(point grid.Point, rest grid.Path) *Node {
	return &Node{Point: point,
		Rest:     rest,
		SubNodes: make([]*Node, 0, 3)}
}

// NodeList used for route tracing
type NodeList struct {
	Nodes []*Node
}

// buildTree build a node tree which contains a point the path left at this point
// and have recursively stored sub-nodes with the same data. The sub-nodes are decided
// by the pathfinder function used pathfinderFn and represent the next viable points in
// the current path.
func buildTree(point grid.Point, rest grid.Path, pathfinderFn PathFinderFunc) *Node {
	node := newNode(point, rest)

	if len(rest.Points) == 0 {
		return node
	}

	c := pathfinderFn(point, rest)
	for _, p := range c.Points {
		newPath := grid.Remove(rest, p)
		node.SubNodes = append(node.SubNodes, buildTree(p, newPath, pathfinderFn))
	}

	return node
}

// traverseTree moves through the tree using nodePath to record the current branch
// once a leaf is encountered nodePath will write its contents into the accumulator acc
// which will hold all paths found.
func traverseTree(root *Node, nodePath *NodeList, acc *[]grid.Path) {
	if nodePath == nil {
		nodePath = &NodeList{Nodes: make([]*Node, 0)}
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
