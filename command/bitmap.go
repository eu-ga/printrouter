package command

import (
	"strconv"
)

// Bitmap printer a bitmap
type Bitmap struct {
	Rawb64 string
	Width  int
	Height int
}

// GetType return bitmap command type
func (Bitmap) GetType() Type {
	return BoldCommandType
}

// ToString converts this command to a string
func (b Bitmap) ToString() string {
	return `Bitmap(rawb64="` + b.Rawb64 + `", width="` + strconv.Itoa(b.Width) + `", height="` + strconv.Itoa(b.Height) + `")`
}
