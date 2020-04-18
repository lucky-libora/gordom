package node

type queryTokenType = byte
type NodeChecker = func(node *Node) bool

const (
	noneToken                queryTokenType = 1
	tagToken                 queryTokenType = 2
	classToken               queryTokenType = 3
	idToken                  queryTokenType = 4
	anyToken                 queryTokenType = 5
	attrToken                queryTokenType = 6
	pseudoClassToken         queryTokenType = 7
	parentToken              queryTokenType = 8
	descendantToken          queryTokenType = 9
	immediatelyPrecededToken queryTokenType = 10
	precededToken            queryTokenType = 11
)

func composeCheckers(checker1 NodeChecker, checker2 NodeChecker) NodeChecker {
	return func(node *Node) bool {
		if !checker1(node) {
			return false
		}
		return checker2 == nil || checker2(node)
	}
}
