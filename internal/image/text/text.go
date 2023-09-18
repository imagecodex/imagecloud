package text

import (
	"unicode"
)

func CalculateTextBoxSize(str string, fs int) (width, height int) {
	var wPx, mwPx int
	var lineCount = 1
	for _, v := range str {
		switch {
		case v == '\n':
			lineCount += 1
			if wPx > mwPx {
				mwPx = wPx
			}
			wPx = 0
		case v == '\r':
		case unicode.Is(unicode.Han, v):
			wPx += 2
		default:
			wPx += 1
		}
	}

	if wPx > mwPx {
		mwPx = wPx
	}

	width = mwPx * fs * 58 / 100
	height = fs * 125 / 100 * lineCount
	return
}
