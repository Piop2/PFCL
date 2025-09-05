package formatter

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/piop2/pfcl/internal/errors"
	"github.com/piop2/pfcl/internal/model"
	"github.com/piop2/pfcl/internal/utils"
)

func Format(v map[string]any, writer io.Writer, indent string) errors.ErrPFCL {
	cursorQueue := model.Queue[[]string]{}

	// Enqueue Root Cursor before start format
	cursorQueue.Enqueue([]string{})

	for !cursorQueue.IsEmpty() {
		// get current cursor
		cursor, _ := cursorQueue.Dequeue()

		table, _ := utils.GetTableAtCursor(v, cursor)

		// write table only when it is not the root
		if len(cursor) != 0 {
			_, err := writer.Write(
				[]byte(
					fmt.Sprintf("\n[%s]\n", strings.Join(cursor, ".")),
				),
			)
			if err != nil {
				return errors.ToErrPFCL(err)
			}
		}

		// format items
		for key, value := range table {
			// TODO: 다른 타입도 받을 수 있게 추가
			//   - [v]: table
			//   - [v]: bool
			//   - [v]: number ( ~int, ~float )
			//   - [v]: string
			//   - [ ]: list

			// enqueue it if value type is map[string]any
			if _, ok := value.(map[string]any); ok {
				cursorQueue.Enqueue(append(cursor, key))
				continue
			}

			// format item
			var s string
			rv := reflect.ValueOf(value)
			switch rv.Kind() {

			case reflect.String:
				s = fmt.Sprintf("\"%s\"", rv.String())

			case reflect.Bool:
				s = strconv.FormatBool(rv.Bool())

			case reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64:
				s = strconv.FormatInt(rv.Int(), 10)

			case reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64:
				s = strconv.FormatUint(rv.Uint(), 10)

			case reflect.Float32,
				reflect.Float64:
				s = strconv.FormatFloat(rv.Float(), 'f', -1, 64)

			default: // fallback
				return errors.ToErrPFCL(
					fmt.Errorf(
						"unsupported format type: %s",
						reflect.TypeOf(value).String(),
					),
				)
			}

			// write converted value
			_, err := writer.Write(
				[]byte(
					fmt.Sprintf("%s = %s\n", key, s),
				),
			)
			if err != nil {
				return errors.ToErrPFCL(err)
			}
		}

		// newline
		//_, err := writer.Write([]byte("\n"))
		//if err != nil {
		//	return errors.ToErrPFCL(err)
		//}
	}

	return nil
}
