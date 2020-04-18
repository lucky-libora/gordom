package node

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Brothers(t *testing.T) {
	parentNode := NewNode("a", nil)
	n1 := parentNode.CreateChild("b")
	n2 := parentNode.CreateChild("c")
	assert.Equal(t, []*Node{n1, n2}, n1.Brothers())
}

func TestNode_BrothersNil(t *testing.T) {
	parentNode := NewNode("a", nil)
	assert.Equal(t, []*Node{parentNode}, parentNode.Brothers())
}

func TestNode_CreateChild(t *testing.T) {
	parentNode := NewNode("a", nil)
	node := parentNode.CreateChild("b")
	assert.Equal(t, "b", node.Tag)
	assert.Equal(t, []*Node{node}, parentNode.Children)
}

func TestNode_FirstChild(t *testing.T) {
	parentNode := NewNode("a", nil)
	n1 := parentNode.CreateChild("b")
	parentNode.CreateChild("c")
	assert.Equal(t, n1, parentNode.FirstChild())
}

func TestNode_FirstChildNil(t *testing.T) {
	parentNode := NewNode("a", nil)
	assert.Nil(t, parentNode.FirstChild())
}

func TestNode_HasChildren(t *testing.T) {
	parentNode := NewNode("a", nil)
	parentNode.CreateChild("b")
	assert.True(t, parentNode.HasChildren())
}

func TestNode_HasChildrenFalse(t *testing.T) {
	parentNode := NewNode("a", nil)
	assert.False(t, parentNode.HasChildren())
}

func TestNode_HasClass(t *testing.T) {
	node := NewNode("a", nil)
	node.Classes = append(node.Classes, "a")
	assert.True(t, node.HasClass("a"))
}

func TestNode_HasClassFalse(t *testing.T) {
	node := NewNode("a", nil)
	assert.False(t, node.HasClass("a"))
}

func TestNode_LastChild(t *testing.T) {
	parentNode := NewNode("a", nil)
	parentNode.CreateChild("b")
	n2 := parentNode.CreateChild("c")
	assert.Equal(t, n2, parentNode.LastChild())
}

func TestNode_LastChildNil(t *testing.T) {
	parentNode := NewNode("a", nil)
	assert.Nil(t, parentNode.LastChild())
}

func TestNode_Parents(t *testing.T) {
	n1 := NewNode("a", nil)
	n2 := n1.CreateChild("b")
	n3 := n2.CreateChild("c")
	assert.Equal(t, []*Node{n1, n2}, n3.Parents())
}

func TestNode_ParentsEmpty(t *testing.T) {
	n1 := NewNode("a", nil)
	assert.Empty(t, n1.Parents())
}

func TestNode_PrevBrother(t *testing.T) {
	parentNode := NewNode("a", nil)
	n1 := parentNode.CreateChild("b")
	n2 := parentNode.CreateChild("c")
	assert.Equal(t, n1, n2.PrevBrother())
}

func TestNode_PrevBrotherNil(t *testing.T) {
	parentNode := NewNode("a", nil)
	n1 := parentNode.CreateChild("b")
	assert.Nil(t, n1.PrevBrother())
}

func TestNode_PrevBrotherNil2(t *testing.T) {
	parentNode := NewNode("a", nil)
	assert.Nil(t, parentNode.PrevBrother())
}

func TestNode_PrevBrothers(t *testing.T) {
	parentNode := NewNode("a", nil)
	n1 := parentNode.CreateChild("b")
	n2 := parentNode.CreateChild("c")
	n3 := parentNode.CreateChild("d")
	assert.Equal(t, []*Node{n1, n2}, n3.PrevBrothers())
}

func TestNode_PrevBrothersEmpty(t *testing.T) {
	parentNode := NewNode("a", nil)
	n1 := parentNode.CreateChild("b")
	assert.Empty(t, n1.PrevBrothers())
}

func TestNode_PrevBrothersEmpty2(t *testing.T) {
	parentNode := NewNode("a", nil)
	assert.Empty(t, parentNode.PrevBrothers())
}

func TestNode_String(t *testing.T) {
	parentNode := NewNode("a", nil)
	node := NewNode("b", parentNode)
	node.Attrs = map[string]string{"a": "a"}
	node.Classes = []string{"a", "b"}
	node.Text = "test"
	assert.Equal(t, node.String(), "b.a.b[a=a]")
}
