package node

import (
	"strings"
)

type Node struct {
	Attrs    map[string]string `json:"attrs"`
	Children []*Node           `json:"children"`
	Classes  []string          `json:"classes"`
	Id       string            `json:"id"`
	Parent   *Node             `json:"-"`
	Tag      string            `json:"tag"`
	Text     string            `json:"text"`
}

func NewNode(tag string, parent *Node) *Node {
	return &Node{
		Attrs:    make(map[string]string),
		Children: []*Node{},
		Classes:  []string{},
		Tag:      tag,
		Parent:   parent,
	}
}

func (node *Node) Brothers() []*Node {
	if node.Parent == nil {
		return []*Node{node}
	}
	return node.Parent.Children
}

func (node *Node) CreateChild(tag string) *Node {
	child := NewNode(tag, node)
	node.Children = append(node.Children, child)
	return child
}

func (node *Node) FirstChild() *Node {
	if !node.HasChildren() {
		return nil
	}
	return node.Children[0]
}

func (node *Node) HasChildren() bool {
	return len(node.Children) > 0
}

func (node *Node) HasClass(class string) bool {
	for _, cls := range node.Classes {
		if cls == class {
			return true
		}
	}
	return false
}

func (node *Node) InnerText() string {
	res := node.Text
	node.ForEachChild(func(child *Node) {
		res += child.InnerText()
	})
	return res
}

func (node *Node) LastChild() *Node {
	if !node.HasChildren() {
		return nil
	}
	index := len(node.Children) - 1
	return node.Children[index]
}

func (node *Node) Parents() []*Node {
	if node.Parent == nil {
		return []*Node{}
	}
	return append(node.Parent.Parents(), node.Parent)
}

func (node *Node) PrevBrother() *Node {
	brothers := node.Brothers()
	for i, brother := range brothers {
		if brother == node && i != 0 {
			return brothers[i-1]
		}
	}
	return nil
}

func (node *Node) PrevBrothers() []*Node {
	brothers := node.Brothers()
	prevBrothers := []*Node{}
	for _, brother := range brothers {
		if brother == node {
			return prevBrothers
		} else {
			prevBrothers = append(prevBrothers, brother)
		}
	}
	return prevBrothers
}

func (node *Node) String() string {
	res := node.Tag
	if len(node.Classes) > 0 {
		res += "." + strings.Join(node.Classes, ".")
	}
	if len(node.Attrs) > 0 {
		for key, value := range node.Attrs {
			res += "[" + key + "=" + value + "]"
		}
	}
	return res
}
