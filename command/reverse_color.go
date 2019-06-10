package command

import (
	"strconv"
)

// ReverseColor changes text color
type ReverseColor struct {
	Enabled bool
}

// GetType returns reverse color command type
func (ReverseColor) GetType() Type {
	return ReverseColorCommandType
}

// ToString converts this command to a string
func (r ReverseColor) ToString() string {
	return `ReverseColor(enabled="` + strconv.FormatBool(r.Enabled) + `")`
}
