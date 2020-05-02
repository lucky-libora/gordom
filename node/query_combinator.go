package node

import (
	"regexp"
)

func CompileQuery(query string) Checker {
	q := prepareQuery(query)
	token := ""
	tokenType := noneToken
	var checker Checker

	addChecker := func() {
		if len(token) == 0 {
			return
		}
		prevChecker := checker
		newChecker := compileOrQuery(token)

		checker = composeCheckersAnd(newChecker, prevChecker)

		switch tokenType {
		case parentToken:
			checker = parentTransformer(checker)
		case descendantToken:
			checker = descendantTransformer(checker)
		case precededToken:
			checker = precededTransformer(checker)
		case immediatelyPrecededToken:
			checker = immediatelyPrecededTransformer(checker)
		}

		tokenType = noneToken
		token = ""
	}

	isQuoteOpened := false
	isAttrBracketOpened := false
	for _, ch := range q {
		if isQuoteOpened {
			token += string(ch)
			if ch == '\'' {
				isQuoteOpened = false
			}
			continue
		}
		if isAttrBracketOpened {
			token += string(ch)
			if ch == ']' {
				isAttrBracketOpened = false
			}
			continue
		}
		var tt queryTokenType
		switch ch {
		case '\'':
			isQuoteOpened = true
		case '[':
			isAttrBracketOpened = true
		case '>':
			tt = parentToken
		case ' ':
			tt = descendantToken
		case '+':
			tt = immediatelyPrecededToken
		case '~':
			tt = precededToken
		}
		if tt != 0 {
			tokenType = tt
			addChecker()
		} else {
			token += string(ch)
		}
	}
	addChecker()
	return checker
}

func parentTransformer(checker Checker) Checker {
	return func(node *Node) bool {
		if node.Parent == nil {
			return false
		}
		return checker(node.Parent)
	}
}

func descendantTransformer(checker Checker) Checker {
	return func(node *Node) bool {
		for _, parent := range node.Parents() {
			if checker(parent) {
				return true
			}
		}
		return false
	}
}

func precededTransformer(checker Checker) Checker {
	return func(node *Node) bool {
		for _, brother := range node.PrevBrothers() {
			if brother != node && checker(brother) {
				return true
			}
		}
		return false
	}
}

func immediatelyPrecededTransformer(checker Checker) Checker {
	return func(node *Node) bool {
		prevBrother := node.PrevBrother()
		if prevBrother == nil {
			return false
		}
		return checker(prevBrother)
	}
}

type regexReplace struct {
	regex *regexp.Regexp
	s     string
}

var regexes = []regexReplace{
	{
		regex: regexp.MustCompile(`\s>`),
		s:     ">",
	},
	{
		regex: regexp.MustCompile(`>\s`),
		s:     ">",
	},
	{
		regex: regexp.MustCompile(`\s\+`),
		s:     "+",
	},
	{
		regex: regexp.MustCompile(`\+\s`),
		s:     "+",
	},
	{
		regex: regexp.MustCompile(`\s~`),
		s:     "~",
	},
	{
		regex: regexp.MustCompile(`~\s`),
		s:     "~",
	},
	{
		regex: regexp.MustCompile(`\s,`),
		s:     ",",
	},
	{
		regex: regexp.MustCompile(`,\s`),
		s:     ",",
	},
}

func prepareQuery(query string) string {
	q := removeDoubleSpaces(query)
	for _, r := range regexes {
		q = r.regex.ReplaceAllString(q, r.s)
	}
	return q
}
