package node

type Document struct {
	Body *Node
	Head *Node
}

func NewDocument(body, head *Node) *Document {
	return &Document{
		Body: body,
		Head: head,
	}
}

func (doc *Document) Select(query string) []*Node {
	return doc.Body.Select(query)
}

func (doc *Document) SelectOne(query string) *Node {
	return doc.Body.SelectOne(query)
}
