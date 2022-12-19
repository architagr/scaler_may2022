package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

var visited map[int]*Node

func CloneGraph(node *Node) *Node {
	visited = make(map[int]*Node)
	return clone(node)
}
func clone(node *Node) *Node {
	if node == nil {
		return node
	}
	newNode, ok := visited[node.Val]
	if ok {
		return newNode
	}
	newNode = new(Node)
	newNode.Val = node.Val
	visited[node.Val] = newNode

	if len(node.Neighbors) > 0 {
		newNode.Neighbors = make([]*Node, len(node.Neighbors))
		for i, nNode := range node.Neighbors {
			newNode.Neighbors[i] = clone(nNode)
		}
	}

	return newNode
}
