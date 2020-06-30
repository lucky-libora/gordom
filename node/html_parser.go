package node

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

const BodyTag = "body"

var selfClosingTags = []string{
	"area", "br", "col", "command", "embed", "hr", "img", "input", "keygen", "link", "meta", "param", "source",
	"track", "wbr",
}

func createChild(tagName string, tokenizer *html.Tokenizer, node *Node) *Node {
	newNode := node.CreateChild(tagName)
	parseAttrs(tokenizer, newNode)
	parseClasses(newNode)
	parseId(newNode)
	return newNode
}

func endTagHandler(node *Node) *Node {
	if node.Tag == BodyTag {
		return node
	}
	return node.Parent
}

func isTagSelfClosing(tag string) bool {
	for _, v := range selfClosingTags {
		if v == tag {
			return true
		}
	}
	return false
}

func parseAttrs(tokenizer *html.Tokenizer, node *Node) *Node {
	for {
		keyBytes, valBytes, isMore := tokenizer.TagAttr()
		key := string(keyBytes)
		val := string(valBytes)
		if len(key) == 0 {
			return node
		}
		node.Attrs[key] = val
		if !isMore {
			return node
		}
	}
}

func parseClasses(node *Node) *Node {
	classes, has := node.Attrs["class"]
	if !has {
		return node
	}
	for _, class := range strings.Split(classes, " ") {
		node.Classes = append(node.Classes, class)
	}
	return node
}

func parseId(node *Node) *Node {
	id, has := node.Attrs["id"]
	if has {
		node.Id = id
	}
	return node
}

func selfClosingTagHandler(tokenizer *html.Tokenizer, node *Node) *Node {
	tag, _ := tokenizer.TagName()
	createChild(string(tag), tokenizer, node)
	return node
}

func startTagHandler(tokenizer *html.Tokenizer, node *Node) *Node {
	tag, _ := tokenizer.TagName()
	tagName := string(tag)
	if isTagSelfClosing(tagName) {
		selfClosingTagHandler(tokenizer, node)
		return node
	}
	return createChild(tagName, tokenizer, node)
}

func textTokenHandler(tokenizer *html.Tokenizer, node *Node) *Node {

	text := cleanText(string(tokenizer.Text()))
	if len(text) != 0 {
		child := node.CreateChild("")
		child.Text = text
	}
	return node
}

func ParseHtml(reader io.Reader) *Document {
	tokenizer := html.NewTokenizer(reader)
	root := NewNode("", nil)
	currentNode := root
	for {
		token := tokenizer.Next()
		if token == 0 {
			break
		}
		switch token {
		case html.ErrorToken:
			continue
		case html.StartTagToken:
			currentNode = startTagHandler(tokenizer, currentNode)
		case html.EndTagToken:
			currentNode = endTagHandler(currentNode)
		case html.SelfClosingTagToken:
			currentNode = selfClosingTagHandler(tokenizer, currentNode)
		case html.TextToken:
			currentNode = textTokenHandler(tokenizer, currentNode)
		}
	}
	htmlNode := root.FirstChild()
	htmlNode.Parent = nil

	head := htmlNode.SelectOne("head")
	body := htmlNode.SelectOne("body")
	return NewDocument(body, head)
}
