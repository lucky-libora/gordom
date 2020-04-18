package go_parser_it

import (
	"github.com/lucky-libora/go-parse-it/node"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Text struct {
	Text string `$:"#text"`
}

func TestParse(t *testing.T) {
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
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, Text{}).(Text)
	assert.Equal(t, "SomeText", text.Text)
}

type NotFound struct {
	Text string `$:"#test"`
}

func TestParseNotFound(t *testing.T) {
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
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, NotFound{}).(NotFound)
	assert.Equal(t, "", text.Text)
}

type NoQuery struct {
	Text string
}

func TestParseNoQuery(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		Test
	</body>
 	`
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, NoQuery{}).(NoQuery)
	assert.Equal(t, "Test", text.Text)
}

type NoQueryCollection struct {
	NoQuery []NoQuery
}

func TestParseNoQueryArray(t *testing.T) {
	s := `
	<html>
	<head>
		<meta>
	</head>
	<body>
		Test
	</body>
 	`
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, NoQueryCollection{}).(NoQueryCollection)
	assert.Equal(t, "Test", text.NoQuery[0].Text)
}

type Float struct {
	Value float64 `$:"#text"`
}

func TestParseFloat(t *testing.T) {
	s := `
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
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, Float{}).(Float)
	assert.Equal(t, 3.14333, text.Value)
}

type Int struct {
	Value int32 `$:"#text"`
}

func TestParseInt(t *testing.T) {
	s := `
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
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, Int{}).(Int)
	assert.Equal(t, int32(3), text.Value)
}

type UInt struct {
	Value uint32 `$:"#text"`
}

func TestParseUInt(t *testing.T) {
	s := `
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
	doc := node.ParseHtml(strings.NewReader(s))
	text := Parse(doc, UInt{}).(UInt)
	assert.Equal(t, uint32(3), text.Value)
}

type Struct struct {
	Text  Text
	Int   Int
	UInt  UInt
	Float Float
}

func TestParseStruct(t *testing.T) {
	s := `
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
	doc := node.ParseHtml(strings.NewReader(s))
	str := Parse(doc, Struct{}).(Struct)
	assert.Equal(t, uint32(3), str.UInt.Value)
}

type Image struct {
	Src string `value:"[src]"`
}

type ImageCollection struct {
	Images []Image `$:"img"`
}

func TestParseArray(t *testing.T) {
	s := `
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
	doc := node.ParseHtml(strings.NewReader(s))
	ic := Parse(doc, ImageCollection{}).(ImageCollection)
	assert.Equal(t, 3, len(ic.Images))
}
