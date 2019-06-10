package command

// Cut prints space for printer to cut paper
type Cut struct {
}

// GetType returns cut command type
func (Cut) GetType() Type {
	return CutCommandType
}

// ToString converts this command to a string
func (Cut) ToString() string {
	return `Cut()`
}
