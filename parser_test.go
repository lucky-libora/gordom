package gordom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Text struct {
	Text string `$:"#text"`
}

func TestParse(t *testing.T) {
	html := `
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
	text := &Text{}
	err := Parse(html, text)
	assert.Nil(t, err)
	assert.Equal(t, "SomeText", text.Text)
}

func TestParseNotPointer(t *testing.T) {
	html := `
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
	text := Text{}
	err := Parse(html, text)
	assert.NotNil(t, err)
}

func TestParseNotStruct(t *testing.T) {
	html := `
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
	text := ""
	err := Parse(html, &text)
	assert.NotNil(t, err)
}

type NotFound struct {
	Text string `$:"#test"`
}

func TestParseNotFound(t *testing.T) {
	html := `
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
	notFound := &NotFound{}
	err := Parse(html, notFound)
	assert.Nil(t, err)
	assert.Equal(t, "", notFound.Text)
}

type NoQuery struct {
	Text string
}

func TestParseNoQuery(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		Test
	</body>
 	`
	noQuery := &NoQuery{}
	err := Parse(html, noQuery)
	assert.Nil(t, err)
	assert.Equal(t, "Test", noQuery.Text)
}

type Float struct {
	Value float64 `$:"#text"`
}

func TestParseFloat(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">3.14333</div>
	</body>
	`
	float := &Float{}
	err := Parse(html, float)
	assert.Nil(t, err)
	assert.Equal(t, 3.14333, float.Value)
}

func TestParseFloatError(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">blabla</div>
	</body>
	`
	float := &Float{}
	err := Parse(html, float)
	assert.NotNil(t, err)
}

type Int struct {
	Value int32 `$:"#text"`
}

func TestParseInt(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">3</div>
	</body>
	`
	i := &Int{}
	err := Parse(html, i)
	assert.Nil(t, err)
	assert.Equal(t, int32(3), i.Value)
}

func TestParseIntError(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">3.14</div>
	</body>
	`
	i := &Int{}
	err := Parse(html, i)
	assert.NotNil(t, err)
}

type UInt struct {
	Value uint32 `$:"#text"`
}

func TestParseUInt(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">3</div>
	</body>
	`
	ui := &UInt{}
	err := Parse(html, ui)
	assert.Nil(t, err)
	assert.Equal(t, uint32(3), ui.Value)
}

func TestParseUIntError(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">-3</div>
	</body>
	`
	ui := &UInt{}
	err := Parse(html, ui)
	assert.NotNil(t, err)
}

type Struct struct {
	Text  Text
	Int   Int
	UInt  UInt
	Float Float
}

func TestParseStruct(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">3</div>
	</body>
	`
	s := &Struct{}
	err := Parse(html, s)
	assert.Nil(t, err)
	assert.Equal(t, "3", s.Text.Text)
	assert.Equal(t, float64(3), s.Float.Value)
	assert.Equal(t, int32(3), s.Int.Value)
	assert.Equal(t, uint32(3), s.UInt.Value)
}

func TestParseStructError(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img id="img" src="img.png"/>
		</div>
		<div id="text">blabla</div>
	</body>
	`
	s := &Struct{}
	err := Parse(html, s)
	assert.NotNil(t, err)
}

type Image struct {
	Src string `value:"[src]"`
}

type ImageCollection struct {
	Images []Image `$:"img"`
}

func TestParseArray(t *testing.T) {
	html := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		<div id="main" class="main test">
			<img src="img1.png"/>
			<img src="img2.png"/>
			<img src="img3.png"/>
		</div>
	</body>
	`
	ic := &ImageCollection{}
	err := Parse(html, ic)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(ic.Images))
	for _, img := range ic.Images {
		assert.NotNil(t, img)
	}
}

type GitHubProjectFile struct {
	Link string `value:"[href]"`
	Name string
}

type GitHubProject struct {
	Files []GitHubProjectFile `$:"a.js-navigation-open"`
}

func TestParseFromUrl(t *testing.T) {
	project := &GitHubProject{}
	err := ParseFromUrl("https://github.com/lucky-libora/go-parse-it", project)
	assert.Nil(t, err)
	assert.NotEmpty(t, project.Files)
}

func TestParseFromUrlError(t *testing.T) {
	project := &GitHubProject{}
	err := ParseFromUrl("https://github404/", project)
	assert.NotNil(t, err)
}
