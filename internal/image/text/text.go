package text

import (
	"unicode"
)

func CalculateTextBoxSize(str string, fs float64) (width, height float64) {
	lineHeight := fs * 0.75

	var lineCount, wPx, mwPx float64
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

	width = mwPx * lineHeight * 0.5
	height = lineHeight*(lineCount+1) + 5*lineCount
	return
}
