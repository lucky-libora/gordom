package node

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDocument_SelectOneById(t *testing.T) {
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
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#main")
	assert.Equal(t, "main", got.Id)
}

func TestDocument_SelectByTag(t *testing.T) {
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
	doc := ParseHtml(strings.NewReader(s))
	got := doc.Select("div")[0]
	assert.Equal(t, "main", got.Id)
}

func TestDocument_SelectByAttr(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[src]")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByAttrValue(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[src='img.png']")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByAttrValueContains(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[src*='png']")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByAttrValueNotEqual(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[src!='png']")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByAttrValueByPrefix(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[src^='img']")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByAttrValueBySuffix(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[src$='png']")
	assert.Equal(t, "img", got.Id, "img")
}

func TestDocument_SelectByAttrValueByMultipleValues(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[id~='main test']")
	assert.Equal(t, "main", got.Id)
}

func TestDocument_SelectByAttrValueByMultipleValuesNil(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("[id~='main1 test']")
	assert.Nil(t, got)
}

func TestDocument_SelectByParent(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#main > #img")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByDescendant(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("body #img")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectByDescendantNil(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#test #img")
	assert.Nil(t, got)
}

func TestDocument_SelectByPreceded(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
			<img id="img1" src="img1.png"/>
			<img id="img2" src="img2.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#img ~ #img2")
	assert.Equal(t, "img2", got.Id)
}

func TestDocument_SelectByPrecededNil(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
			<img id="img1" src="img1.png"/>
			<img id="img2" src="img2.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#img3 ~ #img1")
	assert.Nil(t, got)
}

func TestDocument_SelectByImmediatelyPreceded(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
			<img id="img1" src="img1.png"/>
			<img id="img2" src="img2.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#img1 + #img2")
	assert.Equal(t, "img2", got.Id)
}

func TestDocument_SelectByImmediatelyPrecededNil(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
			<img id="img1" src="img1.png"/>
			<img id="img2" src="img2.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("* + #img")
	assert.Nil(t, got)
}

func TestDocument_SelectEmpty(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="empty"></div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("div:empty")
	assert.Equal(t, "empty", got.Id)
}

func TestDocument_SelectFirstChild(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#main:first-child")
	assert.Equal(t, "main", got.Id)
}

func TestDocument_SelectLastChild(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="empty"></div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("div:last-child")
	assert.Equal(t, "empty", got.Id)
}

func TestDocument_SelectOnlyChild(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="empty"></div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("#img:only-child")
	assert.Equal(t, "img", got.Id)
}

func TestDocument_SelectHas(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="empty"></div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("div:has(#img)")
	assert.Equal(t, "main", got.Id)
}

func TestDocument_SelectContains(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">SomeText</div>
	</body>
 	`
	doc := ParseHtml(strings.NewReader(s))
	got := doc.SelectOne("div:contains(Some)")
	assert.Equal(t, "text", got.Id)
}
