package gordom

import (
	"errors"
	"github.com/lucky-libora/gordom/node"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func getValue(n *node.Node, typeField reflect.StructField) string {
	if n == nil {
		return ""
	}
	valueTag := typeField.Tag.Get("value")
	if strings.HasPrefix(valueTag, "[") && strings.HasSuffix(valueTag, "]") {
		attrKey := strings.Trim(valueTag, "[]")
		return n.Attrs[attrKey]
	}
	return n.InnerText()
}

func findMany(n *node.Node, typeField reflect.StructField) []*node.Node {
	query := typeField.Tag.Get("$")
	if len(query) > 0 {
		return n.Select(query)

	}
	return []*node.Node{n}
}

func findOne(node *node.Node, typeField reflect.StructField) *node.Node {
	query := typeField.Tag.Get("$")
	if len(query) > 0 {
		return node.SelectOne(query)

	}
	return node
}

func parseFloat(value string, valueField reflect.Value) error {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	valueField.SetFloat(f)
	return nil
}

func parseInt(value string, valueField reflect.Value) error {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	valueField.SetInt(i)
	return nil
}

func parseSlice(node *node.Node, typeField reflect.StructField, valueField reflect.Value) error {
	slice := reflect.MakeSlice(typeField.Type, 1, 1)
	sliceType := slice.Index(0).Type()
	slice = reflect.MakeSlice(typeField.Type, 0, 0)
	found := findMany(node, typeField)
	for _, n := range found {
		value, err := createNewStruct(n, sliceType)
		if err != nil {
			return err
		}
		slice = reflect.Append(slice, value)
	}
	valueField.Set(slice)
	return nil
}

func parseStruct(node *node.Node, typeField reflect.StructField, valueField reflect.Value) error {
	str, err := createNewStruct(node, typeField.Type)
	if err != nil {
		return err
	}
	valueField.Set(str)
	return nil
}

func parseUint(value string, valueField reflect.Value) error {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return err
	}
	valueField.SetUint(i)
	return nil
}

func parseValue(node *node.Node, typeField reflect.StructField, valueField reflect.Value) error {
	kind := typeField.Type.Kind()
	switch kind {
	case reflect.Slice:
		return parseSlice(node, typeField, valueField)
	case reflect.Struct:
		return parseStruct(node, typeField, valueField)
	default:
		value := getValue(findOne(node, typeField), typeField)
		switch kind {
		case reflect.Float32, reflect.Float64:
			return parseFloat(value, valueField)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return parseInt(value, valueField)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return parseUint(value, valueField)
		case reflect.String:
			valueField.SetString(value)
		}
	}
	return nil
}

func setFields(node *node.Node, value reflect.Value, t reflect.Type) (reflect.Value, error) {
	for i := 0; i < value.NumField(); i++ {
		typeField := t.Field(i)
		valueField := value.Field(i)
		err := parseValue(node, typeField, valueField)
		if err != nil {
			return value, err
		}
	}
	return value, nil
}

func createNewStruct(node *node.Node, t reflect.Type) (reflect.Value, error) {
	value := reflect.New(t).Elem()
	return setFields(node, value, t)
}

func parseByType(node *node.Node, ptr interface{}) error {
	if reflect.ValueOf(ptr).Kind() != reflect.Ptr {
		return errors.New("pointer to struct should be passed")
	}
	value := reflect.Indirect(reflect.ValueOf(ptr))
	t := value.Type()
	if value.Kind() != reflect.Struct {
		return errors.New("pointer to struct should be passed")
	}
	_, err := setFields(node, value, t)
	return err
}

func parseDocument(doc *node.Document, ptr interface{}) error {
	return parseByType(doc.Body, ptr)
}

func Parse(html string, ptr interface{}) error {
	return ParseReader(strings.NewReader(html), ptr)
}

func ParseReader(reader io.Reader, ptr interface{}) error {
	doc := node.ParseHtml(reader)
	return parseDocument(doc, ptr)
}

func ParseFromUrl(url string, ptr interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	err = ParseReader(resp.Body, ptr)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
