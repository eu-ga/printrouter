package command

import (
	"strconv"
)

// Bold printer following sequence as bold
type Bold struct {
	Enabled bool
}

// GetType return bold command type
func (Bold) GetType() Type {
	return BoldCommandType
}

// ToString converts this command to a string
func (b Bold) ToString() string {
	return `Bold(enabled="` + strconv.FormatBool(b.Enabled) + `")`
}
