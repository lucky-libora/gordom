package node

import (
	"strings"
)

type attrComparator func(value string, exprValue string) bool

var attrsComparatorMap = map[string]attrComparator{
	"~=": func(value string, exprValue string) bool {
		temp := strings.Split(exprValue, " ")
		for _, t := range temp {
			if value == t {
				return true
			}
		}
		return false
	},
	"^=": func(value string, exprValue string) bool {
		return strings.HasPrefix(value, exprValue)
	},
	"$=": func(value string, exprValue string) bool {
		return strings.HasSuffix(value, exprValue)
	},
	"*=": func(value string, exprValue string) bool {
		return strings.Contains(value, exprValue)
	},
	"!=": func(value string, exprValue string) bool {
		return value != exprValue
	},
}

func attrCheck(expr string) NodeChecker {
	if !strings.Contains(expr, "=") {
		return func(node *Node) bool {
			_, has := node.Attrs[expr]
			return has
		}
	}

	temp := strings.Split(expr, "=")
	exprValue := strings.Trim(temp[1], "'")

	for splitter, comparator := range attrsComparatorMap {
		if strings.Contains(expr, splitter) {
			exprKey := strings.Split(expr, splitter)[0]
			return func(node *Node) bool {
				value, has := node.Attrs[exprKey]
				return has && comparator(value, exprValue)
			}
		}
	}

	return func(node *Node) bool {
		value, has := node.Attrs[temp[0]]
		return has && value == exprValue
	}
}
