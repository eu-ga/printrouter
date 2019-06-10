package command

// Text prints text
type Text struct {
	Text string
}

// GetType returns text command type
func (Text) GetType() Type {
	return TextCommandType
}

// ToString converts this command to a string
func (t Text) ToString() string {
	return `Text(text="` + t.Text + `")`
}
