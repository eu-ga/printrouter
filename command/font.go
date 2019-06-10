package command

const (
	// FontA font type FontA
	FontA Font = "FontA"
	// FontB font type FontB
	FontB Font = "FontB"
	// FontC font type FontC
	FontC Font = "FontC"
	// FontD font type FontD
	FontD Font = "FontD"
	// FontE font type FontE
	FontE Font = "FontE"
	// SpecialFontA font type SpecialFontA
	SpecialFontA Font = "SpecialFontA"
	// SpecialFontB font type SpecialFontB
	SpecialFontB Font = "SpecialFontB"
)

// Font changes text font
type Font string

// GetType return font command type
func (Font) GetType() Type {
	return FontCommandType
}

// ToString converts this command to a string
func (f Font) ToString() string {
	return `Font(font="` + string(f) + `")`
}

// GetDefaultFont returns the default font type: FontA
func (Font) GetDefaultFont() Font {
	return FontA
}
