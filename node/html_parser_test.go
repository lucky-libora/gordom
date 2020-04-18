package node

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img />
		</div>
	</body>
 	`
	got := ParseHtml(strings.NewReader(s))
	assert.False(t, got.Body == nil)
	assert.False(t, got.Head == nil)
	assert.Equal(t, "main", got.Body.FirstChild().Id)
	assert.Equal(t, "img", got.Body.FirstChild().FirstChild().Tag)
}
