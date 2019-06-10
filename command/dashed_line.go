package command

// DashedLine prints a dashed line. Line length depends on the font chosen.
type DashedLine struct {
	Font FontBank
}

// GetType return dashed line command type
func (DashedLine) GetType() Type {
	return DashedLineCommandType
}

// ToString converts this command to a string
func (d DashedLine) ToString() string {
	return `DashedLine(font="` + string(d.Font) + `")`
}
