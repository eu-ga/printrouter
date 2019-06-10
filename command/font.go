package command

const (
	// FontA font type FontA
	FontA FontBank = "FontA"
	// FontB font type FontB
	FontB FontBank = "FontB"
	// FontC font type FontC
	FontC FontBank = "FontC"
	// FontD font type FontD
	FontD FontBank = "FontD"
	// FontE font type FontE
	FontE FontBank = "FontE"
	// SpecialFontA font type SpecialFontA
	SpecialFontA FontBank = "SpecialFontA"
	// SpecialFontB font type SpecialFontB
	SpecialFontB FontBank = "SpecialFontB"
)

// FontBank available font types
type FontBank string

// Font changes text font
type Font struct {
	Font FontBank
}

// GetType return font command type
func (Font) GetType() Type {
	return FontCommandType
}

// ToString converts this command to a string
func (f Font) ToString() string {
	return `Font(font="` + string(f.Font) + `")`
}

// GetDefaultFont returns the default font type: FontA
func (Font) GetDefaultFont() FontBank {
	return FontA
}
