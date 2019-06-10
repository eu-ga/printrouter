package command

// NewLine line break command
type NewLine struct {
}

// GetType returns new line command type
func (NewLine) GetType() Type {
	return NewLineCommandType
}

// ToString converts this command to a string
func (NewLine) ToString() string {
	return "NewLine()"
}
