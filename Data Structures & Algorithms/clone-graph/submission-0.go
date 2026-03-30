/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}

	cloned := make(map[*Node]*Node)

	queue := []*Node{node}
	cloned[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, neighbor := range curr.Neighbors {
			if _, found := cloned[neighbor]; !found {
				cloned[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			cloned[curr].Neighbors = append(cloned[curr].Neighbors, cloned[neighbor])
		}
	}

	return cloned[node]
}
