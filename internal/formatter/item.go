package formatter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ValueToString(v any, indent string, indentLevel int) (result string, err error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {

	case reflect.String:
		result = fmt.Sprintf("\"%s\"", rv.String())

	case reflect.Bool:
		result = strconv.FormatBool(rv.Bool())

	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		result = strconv.FormatInt(rv.Int(), 10)

	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		result = strconv.FormatUint(rv.Uint(), 10)

	case reflect.Float32,
		reflect.Float64:
		result = strconv.FormatFloat(rv.Float(), 'f', -1, 64)

	case reflect.Slice:
		result, err = SliceToString(v.([]any), indent, indentLevel+1)

	default: // fallback
		return "", fmt.Errorf(
			"unsupported format type: %s",
			reflect.TypeOf(v).String(),
		)
	}

	return result, nil
}

func SliceToString(s []any, indent string, indentLevel int) (result string, err error) {
	result += "{\n"

	for _, v := range s {
		r, err := ValueToString(v, indent, indentLevel)
		if err != nil {
			return "", err
		}

		result += fmt.Sprintf("%s%s,\n", strings.Repeat(indent, indentLevel), r)
	}

	result += fmt.Sprintf("%s}", strings.Repeat(indent, indentLevel-1))

	return result, nil
}
