package utils

import "errors"

// GetTableAtCursor get nested map along the cursor path
func GetTableAtCursor(v map[string]any, cursor []string) (table map[string]any, err error) {
	table = v
	for _, key := range cursor {
		if nested, ok := table[key].(map[string]any); ok {
			table = nested
		} else {
			err = errors.New("table name error")
			return nil, err
		}
	}
	return table, nil
}
