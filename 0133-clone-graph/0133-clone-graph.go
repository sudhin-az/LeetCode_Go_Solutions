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
   m := make(map[*Node]*Node)
   var clone func(*Node)*Node
   clone = func(n *Node)*Node {
    if copiedNode, found := m[n]; found {
        return copiedNode
    }
    copyNode := &Node{Val: n.Val}
    m[n] = copyNode

    for _, neighbor := range n.Neighbors {
        copyNode.Neighbors = append(copyNode.Neighbors, clone(neighbor))
    }
    return copyNode
   }
   return clone(node)
}
