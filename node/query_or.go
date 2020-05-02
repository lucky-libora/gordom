package node

import "strings"

func compileOrQuery(query string) Checker {
	temp := strings.Split(query, ",")
	var res Checker
	for _, q := range temp {
		checker := compileSingleQuery(q)
		res = composeCheckersOr(checker, res)
	}
	return res
}
