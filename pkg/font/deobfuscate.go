package font

import (
	"html"

	"golang.org/x/image/font/sfnt"
)

var glyphNameMap = map[string]string{
	"period": ".",
	"zero":   "0",
	"one":    "1",
	"two":    "2",
	"three":  "3",
	"four":   "4",
	"five":   "5",
	"six":    "6",
	"seven":  "7",
	"eight":  "8",
	"nine":   "9",
}

func glyphName(font *sfnt.Font, r rune) (ret string, err error) {
	b := &sfnt.Buffer{}
	index, err := font.GlyphIndex(b, r)
	if err != nil {
		return
	}
	if index == 0 {
		return
	}
	ret, err = font.GlyphName(b, index)
	return
}

// Deobfuscate font protected number.
// characters that not found in font remains unchanged.
func Deobfuscate(v string, font *sfnt.Font) (ret string, err error) {
	for _, i := range html.UnescapeString(v) {
		var name string
		name, err = glyphName(font, i)
		trueValue, ok := glyphNameMap[name]
		if ok {
			ret += string(trueValue)
		} else {
			ret += string(i)
		}
	}
	return
}
