package node

func compileSingleQuery(query string) NodeChecker {
	token := ""
	tokenType := noneToken
	var checker NodeChecker

	appendCheck := func() {
		if len(token) == 0 {
			return
		}
		var newChecker NodeChecker
		switch tokenType {
		case tagToken:
			newChecker = tagCheck(token)
		case classToken:
			newChecker = classCheck(token)
		case idToken:
			newChecker = idCheck(token)
		case anyToken:
			newChecker = anyChecker
		case attrToken:
			newChecker = attrCheck(token)
		case pseudoClassToken:
			newChecker = pseudoClassCheck(token)
		}
		if newChecker != nil {
			checker = composeCheckers(newChecker, checker)
		}
		tokenType = noneToken
		token = ""
	}

	isAttrOpened := false
	isBracketOpened := false

	for _, ch := range query {
		if isAttrOpened {
			if ch == ']' {
				appendCheck()
				isAttrOpened = false
			} else {
				token += string(ch)
			}
			continue
		}
		if isBracketOpened {
			token += string(ch)
			if ch == ')' {
				appendCheck()
				isBracketOpened = false
			}
			continue
		}

		switch ch {
		case '*':
			tokenType = anyToken
			appendCheck()
		case '.':
			appendCheck()
			tokenType = classToken
		case '#':
			appendCheck()
			tokenType = idToken
		case '[':
			appendCheck()
			tokenType = attrToken
			isAttrOpened = true
		case '(':
			token += string(ch)
			isBracketOpened = true
		case ':':
			appendCheck()
			tokenType = pseudoClassToken
		default:
			if tokenType == noneToken {
				tokenType = tagToken
			}
			token += string(ch)
		}
	}
	appendCheck()

	return checker
}

func tagCheck(token string) NodeChecker {
	return func(node *Node) bool {
		return node.Tag == token
	}
}

func classCheck(token string) NodeChecker {
	return func(node *Node) bool {
		return node.HasClass(token)
	}
}

func idCheck(token string) NodeChecker {
	return func(node *Node) bool {
		return node.Id == token
	}
}

func anyChecker(node *Node) bool {
	return true
}
