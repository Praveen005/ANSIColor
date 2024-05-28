package main

import (
	"fmt"
	"strconv"
)

const colorFmt = "\x1b[%vm%s\x1b[0m"

// Paint describes a terminal color.
type Paint int

// Defines basic ANSI colors.
// `iota` is reset to 0 whenever the const keyword appears.
// In a const block, `iota` increments by 1 for each line.
const (
	Black     Paint = iota + 30 // 30
	Red                         // 31
	Green                       // 32
	Yellow                      // 33
	Blue                        // 34
	Magenta                     // 35
	Cyan                        // 36
	LightGray                   // 37
	DarkGray  = 90

	Bold = 1
)

// Colorize returns an ASCII colored string based on given color.
func Colorize(s string, c Paint) string {

	if c == 0 {
		return s
	}
	return fmt.Sprintf(colorFmt, c, s)
}

// Colorize a particular part of the string
func ColorizeAt(s string, idx int, color Paint, underscore bool) string {
	if idx < 0 || len(s) <= idx {
		return s
	}

	chrs := []rune(s)
	left := string(chrs[:idx])

	mid := fmt.Sprintf(colorFmt, color, string(chrs[idx]))
	if underscore {
		ansiAttributes := fmt.Sprintf("4;%d", color)
		mid = fmt.Sprintf(colorFmt, ansiAttributes, string(chrs[idx]))
	}
	right := string(chrs[idx+1:])
	return fmt.Sprintf("%s%s%s", left, mid, right)
}

// ANSIColorize colors a string. `ESC[38;5;{ID}m` -- Sets foreground color. ID ranges from 1 to 255
func ANSIColorize(text string, color int) string {
	return "\033[38;5;" + strconv.Itoa(color) + "m" + text + "\033[0m"
}

// Highlight colorize bytes at given indices.
func Highlight(bb []byte, ii []int, c int) []byte {
	b := make([]byte, 0, len(bb))
	for i, j := 0, 0; i < len(bb); i++ {
		if j < len(ii) && ii[j] == i {
			b = append(b, colorizeByte(bb[i], 209)...)
			j++
		} else {
			b = append(b, bb[i])
		}
	}

	return b
}

func colorizeByte(b byte, color int) []byte {
	return []byte(ANSIColorize(string(b), color))
}

func main(){

	coloredText := Colorize("Prateek", 36)
	fmt.Println(coloredText)

	colorAtIndex := ColorizeAt("Praveen Kumar", 8, 33, true)
	fmt.Println(colorAtIndex)

	foregroundColored := ANSIColorize("Praveen Ji", 158)
	fmt.Println(foregroundColored)

	text := "Mr. Modi looks young!"

	highlightedByte := Highlight([]byte(text), []int{1, 5, 7, 9, 11, 13, 15, 17, 19}, 35)
	fmt.Println(string(highlightedByte))
}