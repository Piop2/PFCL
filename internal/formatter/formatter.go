package formatter

import (
	"fmt"
	"io"
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
			// enqueue it if value type is map[string]any
			if _, ok := value.(map[string]any); ok {
				cursorQueue.Enqueue(append(cursor, key))
				continue
			}

			s, err := ValueToString(value, indent, 0)
			if err != nil {
				return errors.ToErrPFCL(err)
			}

			// write converted value
			_, err = writer.Write(
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
