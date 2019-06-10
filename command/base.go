package command

const (
	// NewLineCommandType new line command type
	NewLineCommandType Type = "NewLine"
	// TextCommandType text command type
	TextCommandType Type = "Text"
	// ReverseColorCommandType reverse color command type
	ReverseColorCommandType Type = "ReverseColor"
	// FontCommandType font command type
	FontCommandType Type = "Font"
	// DashedLineCommandType dashed line command type
	DashedLineCommandType Type = "DashedLine"
	// BoldCommandType bold command type
	BoldCommandType Type = "Bold"
	// BitmapCommandType Bitmap command type
	BitmapCommandType Type = "Bitmap"
	// CutCommandType Cut command type
	CutCommandType Type = "Cut"
)

// Type printer command type
type Type string

// PrinterCommand printable command
type PrinterCommand interface {
	GetType() Type
	ToString() string
}
