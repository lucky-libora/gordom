package node

func (node *Node) DoUntil(f func(n *Node) bool) *Node {
	node.doUntil(f)
	return node
}

func (node *Node) ForEach(f func(n *Node)) *Node {
	f(node)
	return node.ForEachChild(func(child *Node) {
		child.ForEach(f)
	})
}

func (node *Node) ForEachChild(f func(n *Node)) *Node {
	for _, child := range node.Children {
		f(child)
	}
	return node
}

func (node *Node) Filter(pred func(n *Node) bool) []*Node {
	var filtered []*Node
	node.ForEach(func(n *Node) {
		if pred(n) {
			filtered = append(filtered, n)
		}
	})
	return filtered
}

func (node *Node) FilterByTag(tag string) []*Node {
	return node.Filter(func(n *Node) bool {
		return n.Tag == tag
	})
}

func (node *Node) Find(pred func(n *Node) bool) *Node {
	var res *Node = nil
	node.DoUntil(func(n *Node) bool {
		if pred(n) {
			res = n
			return true
		}
		return false
	})
	return res
}

func (node *Node) doUntil(f func(n *Node) bool) bool {
	if f(node) {
		return true
	}
	for _, child := range node.Children {
		res := child.doUntil(f)
		if res {
			return true
		}
	}
	return false
}
