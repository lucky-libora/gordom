package node

import (
	"strings"
)

func pseudoClassCheck(expr string) Checker {
	switch expr {
	case "empty":
		return emptyCheck
	case "first-child":
		return firstChildCheck
	case "last-child":
		return lastChildCheck
	case "only-child":
		return onlyChildCheck
	}
	if strings.Contains(expr, "(") {
		return parameterizedPseudoCheck(expr)
	}
	return falseCheck
}

func parameterizedPseudoCheck(expr string) Checker {
	temp := strings.Split(expr, "(")[1]
	value := strings.Trim(temp, "()'")
	if strings.HasPrefix(expr, "contains") {
		return func(node *Node) bool {
			return strings.Contains(node.InnerText(), value)
		}
	}
	if strings.HasPrefix(expr, "has") {
		checker := CompileQuery(value)
		return func(node *Node) bool {
			for _, child := range node.Children {
				if child.Find(checker) != nil {
					return true
				}
			}
			return false
		}
	}
	return falseCheck
}

func falseCheck(n *Node) bool {
	return false
}

func emptyCheck(node *Node) bool {
	return !node.HasChildren()
}

func firstChildCheck(node *Node) bool {
	return node.Parent.FirstChild() == node
}

func lastChildCheck(node *Node) bool {
	return node.Parent.LastChild() == node
}

func onlyChildCheck(node *Node) bool {
	return len(node.Parent.Children) == 1
}
