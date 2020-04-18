package node

func (node *Node) Select(query string) []*Node {
	queryCheck := CompileQuery(query)
	return node.Filter(queryCheck)

}

func (node *Node) SelectOne(query string) *Node {
	queryCheck := CompileQuery(query)
	return node.Find(queryCheck)
}
