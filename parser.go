package go_parser_it

import (
	"go-parser-it/node"
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
	return n.Text
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

func parseValue(node *node.Node, typeField reflect.StructField, valueField reflect.Value) {
	kind := typeField.Type.Kind()
	switch kind {
	case reflect.Slice:
		slice := reflect.MakeSlice(typeField.Type, 1, 1)
		sliceType := slice.Index(0).Type()
		slice = reflect.MakeSlice(typeField.Type, 0, 0)
		found := findMany(node, typeField)
		for _, n := range found {
			slice = reflect.Append(slice, parseByType(n, sliceType))
		}
		valueField.Set(slice)
	case reflect.Struct:
		valueField.Set(parseByType(findOne(node, typeField), typeField.Type))
	default:
		value := getValue(findOne(node, typeField), typeField)
		switch kind {
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(value, 64)
			check(err)
			valueField.SetFloat(f)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(value, 10, 64)
			check(err)
			valueField.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			i, err := strconv.ParseUint(value, 10, 64)
			check(err)
			valueField.SetUint(i)
		case reflect.String:
			valueField.SetString(value)
		}
	}
}

func parseByType(node *node.Node, t reflect.Type) reflect.Value {
	value := reflect.New(t).Elem()
	for i := 0; i < value.NumField(); i++ {
		typeField := t.Field(i)
		valueField := value.Field(i)
		parseValue(node, typeField, valueField)
	}
	return value
}

func Parse(node *node.Document, value interface{}) interface{} {
	t := reflect.TypeOf(value)
	return parseByType(node.Body, t).Interface()
}
